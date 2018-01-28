package model

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/primefour/xserver/utils/markdown"
)

const (
	POST_SYSTEM_MESSAGE_PREFIX  = "system_"
	POST_DEFAULT                = ""
	POST_SLACK_ATTACHMENT       = "slack_attachment"
	POST_SYSTEM_GENERIC         = "system_generic"
	POST_JOIN_LEAVE             = "system_join_leave" // Deprecated, use POST_JOIN_CHANNEL or POST_LEAVE_CHANNEL instead
	POST_JOIN_CHANNEL           = "system_join_channel"
	POST_LEAVE_CHANNEL          = "system_leave_channel"
	POST_JOIN_TEAM              = "system_join_team"
	POST_LEAVE_TEAM             = "system_leave_team"
	POST_ADD_REMOVE             = "system_add_remove" // Deprecated, use POST_ADD_TO_CHANNEL or POST_REMOVE_FROM_CHANNEL instead
	POST_ADD_TO_CHANNEL         = "system_add_to_channel"
	POST_REMOVE_FROM_CHANNEL    = "system_remove_from_channel"
	POST_ADD_TO_TEAM            = "system_add_to_team"
	POST_REMOVE_FROM_TEAM       = "system_remove_from_team"
	POST_HEADER_CHANGE          = "system_header_change"
	POST_DISPLAYNAME_CHANGE     = "system_displayname_change"
	POST_PURPOSE_CHANGE         = "system_purpose_change"
	POST_CHANNEL_DELETED        = "system_channel_deleted"
	POST_EPHEMERAL              = "system_ephemeral"
	POST_CHANGE_CHANNEL_PRIVACY = "system_change_chan_privacy"
	POST_FILEIDS_MAX_RUNES      = 150
	POST_FILENAMES_MAX_RUNES    = 4000
	POST_HASHTAGS_MAX_RUNES     = 1000
	POST_MESSAGE_MAX_RUNES      = 4000
	POST_PROPS_MAX_RUNES        = 8000
	POST_PROPS_MAX_USER_RUNES   = POST_PROPS_MAX_RUNES - 400 // Leave some room for system / pre-save modifications
	POST_CUSTOM_TYPE_PREFIX     = "custom_"
	PROPS_ADD_CHANNEL_MEMBER    = "add_channel_member"
)

type Post struct {
	Id         string `json:"id"`
	CreateAt   int64  `json:"create_at"`
	UpdateAt   int64  `json:"update_at"`
	EditAt     int64  `json:"edit_at"`
	DeleteAt   int64  `json:"delete_at"`
	IsPinned   bool   `json:"is_pinned"`
	UserId     string `json:"user_id"`
	ChannelId  string `json:"channel_id"`
	RootId     string `json:"root_id"`
	ParentId   string `json:"parent_id"`
	OriginalId string `json:"original_id"`

	Message string `json:"message"`

	// MessageSource will contain the message as submitted by the user if Message has been modified
	// by Mattermost for presentation (e.g if an image proxy is being used). It should be used to
	// populate edit boxes if present.
	MessageSource string `json:"message_source,omitempty" db:"-"`

	Type          string          `json:"type"`
	Props         StringInterface `json:"props"`
	Hashtags      string          `json:"hashtags"`
	Filenames     StringArray     `json:"filenames,omitempty"` // Deprecated, do not use this field any more
	FileIds       StringArray     `json:"file_ids,omitempty"`
	PendingPostId string          `json:"pending_post_id" db:"-"`
	HasReactions  bool            `json:"has_reactions,omitempty"`
}

type PostPatch struct {
	IsPinned     *bool            `json:"is_pinned"`
	Message      *string          `json:"message"`
	Props        *StringInterface `json:"props"`
	FileIds      *StringArray     `json:"file_ids"`
	HasReactions *bool            `json:"has_reactions"`
}

func (o *PostPatch) WithRewrittenImageURLs(f func(string) string) *PostPatch {
	copy := *o
	if copy.Message != nil {
		*copy.Message = RewriteImageURLs(*o.Message, f)
	}
	return &copy
}

type PostForIndexing struct {
	Post
	TeamId         string `json:"team_id"`
	ParentCreateAt *int64 `json:"parent_create_at"`
}

type PostAction struct {
	Id          string                 `json:"id"`
	Name        string                 `json:"name"`
	Integration *PostActionIntegration `json:"integration,omitempty"`
}

type PostActionIntegration struct {
	URL     string          `json:"url,omitempty"`
	Context StringInterface `json:"context,omitempty"`
}

type PostActionIntegrationRequest struct {
	UserId  string          `json:"user_id"`
	Context StringInterface `json:"context,omitempty"`
}

type PostActionIntegrationResponse struct {
	Update        *Post  `json:"update"`
	EphemeralText string `json:"ephemeral_text"`
}

func (o *Post) ToJson() string {
	copy := *o
	copy.StripActionIntegrations()
	b, err := json.Marshal(&copy)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func PostFromJson(data io.Reader) *Post {
	decoder := json.NewDecoder(data)
	var o Post
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *Post) Etag() string {
	return Etag(o.Id, o.UpdateAt)
}

func (o *Post) IsValid() *AppError {

	if len(o.Id) != 26 {
		return NewAppError("Post.IsValid", "model.post.is_valid.id.app_error", nil, "", http.StatusBadRequest)
	}

	if o.CreateAt == 0 {
		return NewAppError("Post.IsValid", "model.post.is_valid.create_at.app_error", nil, "id="+o.Id, http.StatusBadRequest)
	}

	if o.UpdateAt == 0 {
		return NewAppError("Post.IsValid", "model.post.is_valid.update_at.app_error", nil, "id="+o.Id, http.StatusBadRequest)
	}

	if len(o.UserId) != 26 {
		return NewAppError("Post.IsValid", "model.post.is_valid.user_id.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.ChannelId) != 26 {
		return NewAppError("Post.IsValid", "model.post.is_valid.channel_id.app_error", nil, "", http.StatusBadRequest)
	}

	if !(len(o.RootId) == 26 || len(o.RootId) == 0) {
		return NewAppError("Post.IsValid", "model.post.is_valid.root_id.app_error", nil, "", http.StatusBadRequest)
	}

	if !(len(o.ParentId) == 26 || len(o.ParentId) == 0) {
		return NewAppError("Post.IsValid", "model.post.is_valid.parent_id.app_error", nil, "", http.StatusBadRequest)
	}

	if len(o.ParentId) == 26 && len(o.RootId) == 0 {
		return NewAppError("Post.IsValid", "model.post.is_valid.root_parent.app_error", nil, "", http.StatusBadRequest)
	}

	if !(len(o.OriginalId) == 26 || len(o.OriginalId) == 0) {
		return NewAppError("Post.IsValid", "model.post.is_valid.original_id.app_error", nil, "", http.StatusBadRequest)
	}

	if utf8.RuneCountInString(o.Message) > POST_MESSAGE_MAX_RUNES {
		return NewAppError("Post.IsValid", "model.post.is_valid.msg.app_error", nil, "id="+o.Id, http.StatusBadRequest)
	}

	if utf8.RuneCountInString(o.Hashtags) > POST_HASHTAGS_MAX_RUNES {
		return NewAppError("Post.IsValid", "model.post.is_valid.hashtags.app_error", nil, "id="+o.Id, http.StatusBadRequest)
	}

	switch o.Type {
	case
		POST_DEFAULT,
		POST_JOIN_LEAVE,
		POST_ADD_REMOVE,
		POST_JOIN_CHANNEL,
		POST_LEAVE_CHANNEL,
		POST_JOIN_TEAM,
		POST_LEAVE_TEAM,
		POST_ADD_TO_CHANNEL,
		POST_REMOVE_FROM_CHANNEL,
		POST_ADD_TO_TEAM,
		POST_REMOVE_FROM_TEAM,
		POST_SLACK_ATTACHMENT,
		POST_HEADER_CHANGE,
		POST_PURPOSE_CHANGE,
		POST_DISPLAYNAME_CHANGE,
		POST_CHANNEL_DELETED,
		POST_CHANGE_CHANNEL_PRIVACY:
	default:
		if !strings.HasPrefix(o.Type, POST_CUSTOM_TYPE_PREFIX) {
			return NewAppError("Post.IsValid", "model.post.is_valid.type.app_error", nil, "id="+o.Type, http.StatusBadRequest)
		}
	}

	if utf8.RuneCountInString(ArrayToJson(o.Filenames)) > POST_FILENAMES_MAX_RUNES {
		return NewAppError("Post.IsValid", "model.post.is_valid.filenames.app_error", nil, "id="+o.Id, http.StatusBadRequest)
	}

	if utf8.RuneCountInString(ArrayToJson(o.FileIds)) > POST_FILEIDS_MAX_RUNES {
		return NewAppError("Post.IsValid", "model.post.is_valid.file_ids.app_error", nil, "id="+o.Id, http.StatusBadRequest)
	}

	if utf8.RuneCountInString(StringInterfaceToJson(o.Props)) > POST_PROPS_MAX_RUNES {
		return NewAppError("Post.IsValid", "model.post.is_valid.props.app_error", nil, "id="+o.Id, http.StatusBadRequest)
	}

	return nil
}

func (o *Post) SanitizeProps() {
	membersToSanitize := []string{
		PROPS_ADD_CHANNEL_MEMBER,
	}

	for _, member := range membersToSanitize {
		if _, ok := o.Props[member]; ok {
			delete(o.Props, member)
		}
	}
}

func (o *Post) PreSave() {
	if o.Id == "" {
		o.Id = NewId()
	}

	o.OriginalId = ""

	if o.CreateAt == 0 {
		o.CreateAt = GetMillis()
	}

	o.UpdateAt = o.CreateAt
	o.PreCommit()
}

func (o *Post) PreCommit() {
	if o.Props == nil {
		o.Props = make(map[string]interface{})
	}

	if o.Filenames == nil {
		o.Filenames = []string{}
	}

	if o.FileIds == nil {
		o.FileIds = []string{}
	}

	o.GenerateActionIds()
}

func (o *Post) MakeNonNil() {
	if o.Props == nil {
		o.Props = make(map[string]interface{})
	}
}

func (o *Post) AddProp(key string, value interface{}) {

	o.MakeNonNil()

	o.Props[key] = value
}

func (o *Post) IsSystemMessage() bool {
	return len(o.Type) >= len(POST_SYSTEM_MESSAGE_PREFIX) && o.Type[:len(POST_SYSTEM_MESSAGE_PREFIX)] == POST_SYSTEM_MESSAGE_PREFIX
}

func (p *Post) Patch(patch *PostPatch) {
	if patch.IsPinned != nil {
		p.IsPinned = *patch.IsPinned
	}

	if patch.Message != nil {
		p.Message = *patch.Message
	}

	if patch.Props != nil {
		p.Props = *patch.Props
	}

	if patch.FileIds != nil {
		p.FileIds = *patch.FileIds
	}

	if patch.HasReactions != nil {
		p.HasReactions = *patch.HasReactions
	}
}

func (o *PostPatch) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	}

	return string(b)
}

func PostPatchFromJson(data io.Reader) *PostPatch {
	decoder := json.NewDecoder(data)
	var post PostPatch
	err := decoder.Decode(&post)
	if err != nil {
		return nil
	}

	return &post
}

var channelMentionRegexp = regexp.MustCompile(`\B~[a-zA-Z0-9\-_]+`)

func (o *Post) ChannelMentions() (names []string) {
	if strings.Contains(o.Message, "~") {
		alreadyMentioned := make(map[string]bool)
		for _, match := range channelMentionRegexp.FindAllString(o.Message, -1) {
			name := match[1:]
			if !alreadyMentioned[name] {
				names = append(names, name)
				alreadyMentioned[name] = true
			}
		}
	}
	return
}

func (r *PostActionIntegrationRequest) ToJson() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (o *Post) Attachments() []*SlackAttachment {
	if attachments, ok := o.Props["attachments"].([]*SlackAttachment); ok {
		return attachments
	}
	var ret []*SlackAttachment
	if attachments, ok := o.Props["attachments"].([]interface{}); ok {
		for _, attachment := range attachments {
			if enc, err := json.Marshal(attachment); err == nil {
				var decoded SlackAttachment
				if json.Unmarshal(enc, &decoded) == nil {
					ret = append(ret, &decoded)
				}
			}
		}
	}
	return ret
}

func (o *Post) StripActionIntegrations() {
	attachments := o.Attachments()
	if o.Props["attachments"] != nil {
		o.Props["attachments"] = attachments
	}
	for _, attachment := range attachments {
		for _, action := range attachment.Actions {
			action.Integration = nil
		}
	}
}

func (o *Post) GetAction(id string) *PostAction {
	for _, attachment := range o.Attachments() {
		for _, action := range attachment.Actions {
			if action.Id == id {
				return action
			}
		}
	}
	return nil
}

func (o *Post) GenerateActionIds() {
	if o.Props["attachments"] != nil {
		o.Props["attachments"] = o.Attachments()
	}
	if attachments, ok := o.Props["attachments"].([]*SlackAttachment); ok {
		for _, attachment := range attachments {
			for _, action := range attachment.Actions {
				if action.Id == "" {
					action.Id = NewId()
				}
			}
		}
	}
}

var markdownDestinationEscaper = strings.NewReplacer(
	`\`, `\\`,
	`<`, `\<`,
	`>`, `\>`,
	`(`, `\(`,
	`)`, `\)`,
)

// WithRewrittenImageURLs returns a new shallow copy of the post where the message has been
// rewritten via RewriteImageURLs.
func (o *Post) WithRewrittenImageURLs(f func(string) string) *Post {
	copy := *o
	copy.Message = RewriteImageURLs(o.Message, f)
	if copy.MessageSource == "" && copy.Message != o.Message {
		copy.MessageSource = o.Message
	}
	return &copy
}

// RewriteImageURLs takes a message and returns a copy that has all of the image URLs replaced
// according to the function f. For each image URL, f will be invoked, and the resulting markdown
// will contain the URL returned by that invocation instead.
//
// Image URLs are destination URLs used in inline images or reference definitions that are used
// anywhere in the input markdown as an image.
func RewriteImageURLs(message string, f func(string) string) string {
	if !strings.Contains(message, "![") {
		return message
	}

	var ranges []markdown.Range

	markdown.Inspect(message, func(blockOrInline interface{}) bool {
		switch v := blockOrInline.(type) {
		case *markdown.ReferenceImage:
			ranges = append(ranges, v.ReferenceDefinition.RawDestination)
		case *markdown.InlineImage:
			ranges = append(ranges, v.RawDestination)
		default:
			return true
		}
		return true
	})

	if ranges == nil {
		return message
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Position < ranges[j].Position
	})

	copyRanges := make([]markdown.Range, 0, len(ranges))
	urls := make([]string, 0, len(ranges))
	resultLength := len(message)

	start := 0
	for i, r := range ranges {
		switch {
		case i == 0:
		case r.Position != ranges[i-1].Position:
			start = ranges[i-1].End
		default:
			continue
		}
		original := message[r.Position:r.End]
		replacement := markdownDestinationEscaper.Replace(f(markdown.Unescape(original)))
		resultLength += len(replacement) - len(original)
		copyRanges = append(copyRanges, markdown.Range{Position: start, End: r.Position})
		urls = append(urls, replacement)
	}

	result := make([]byte, resultLength)

	offset := 0
	for i, r := range copyRanges {
		offset += copy(result[offset:], message[r.Position:r.End])
		offset += copy(result[offset:], urls[i])
	}
	copy(result[offset:], message[ranges[len(ranges)-1].End:])

	return string(result)
}
