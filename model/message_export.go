
package model

type MessageExport struct {
	ChannelId          *string
	ChannelDisplayName *string

	UserId    *string
	UserEmail *string

	PostId       *string
	PostCreateAt *int64
	PostMessage  *string
	PostType     *string
	PostFileIds  StringArray
}
