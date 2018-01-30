package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"github.com/pborman/uuid"
	"io"
	"io/ioutil"
	mrand "math/rand"
	"net"
	"net/http"
	"net/mail"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

var (
	IMAGE_EXTENSIONS = [5]string{".jpg", ".jpeg", ".gif", ".bmp", ".png"}
	IMAGE_MIME_TYPES = map[string]string{".jpg": "image/jpeg", ".jpeg": "image/jpeg", ".gif": "image/gif", ".bmp": "image/bmp", ".png": "image/png", ".tiff": "image/tiff"}
)

type StringInterface map[string]interface{}
type StringMap map[string]string
type StringArray []string
type EncryptStringMap map[string]string

const (
	LOWERCASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS           = "0123456789"
	SYMBOLS           = " !\"\\#$%&'()*+,-./:;<=>?@[]^_`|~"
)

const (
	HEADER_REQUEST_ID         = "X-Request-ID"
	HEADER_VERSION_ID         = "X-Version-ID"
	HEADER_CLUSTER_ID         = "X-Cluster-ID"
	HEADER_ETAG_SERVER        = "ETag"
	HEADER_ETAG_CLIENT        = "If-None-Match"
	HEADER_FORWARDED          = "X-Forwarded-For"
	HEADER_REAL_IP            = "X-Real-IP"
	HEADER_FORWARDED_PROTO    = "X-Forwarded-Proto"
	HEADER_TOKEN              = "token"
	HEADER_BEARER             = "BEARER"
	HEADER_AUTH               = "Authorization"
	HEADER_REQUESTED_WITH     = "X-Requested-With"
	HEADER_REQUESTED_WITH_XML = "XMLHttpRequest"
)

func StringArrayIntersection(arr1, arr2 []string) []string {
	arrMap := map[string]bool{}
	result := []string{}

	for _, value := range arr1 {
		arrMap[value] = true
	}

	for _, value := range arr2 {
		if arrMap[value] {
			result = append(result, value)
		}
	}

	return result
}

func RemoveDuplicatesFromStringArray(arr []string) []string {
	result := make([]string, 0, len(arr))
	seen := make(map[string]bool)

	for _, item := range arr {
		if !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}

	return result
}

func GetIpAddress(r *http.Request) string {
	address := r.Header.Get(HEADER_FORWARDED)

	if len(address) == 0 {
		address = r.Header.Get(HEADER_REAL_IP)
	}

	if len(address) == 0 {
		address, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	return address
}

func FindDir(dir string) string {
	fileName := "."
	if _, err := os.Stat("./" + dir + "/"); err == nil {
		fileName, _ = filepath.Abs("./" + dir + "/")
	} else if _, err := os.Stat("../" + dir + "/"); err == nil {
		fileName, _ = filepath.Abs("../" + dir + "/")
	} else if _, err := os.Stat("../../" + dir + "/"); err == nil {
		fileName, _ = filepath.Abs("../../" + dir + "/")
	}

	return fileName + "/"
}

func UrlEncode(str string) string {
	strs := strings.Split(str, " ")

	for i, s := range strs {
		strs[i] = url.QueryEscape(s)
	}

	return strings.Join(strs, "%20")
}

func MillisFromTime(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

type Range struct {
	Begin int
	End   int
}

func RandIntFromRange(r Range) int {
	if r.End-r.Begin <= 0 {
		return r.Begin
	}
	return mrand.Intn((r.End-r.Begin)+1) + r.Begin
}

func HashSha256(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func CloseBody(r *http.Response) {
	if r.Body != nil {
		ioutil.ReadAll(r.Body)
		r.Body.Close()
	}
}

// NewId is a globally unique identifier.  It is a [A-Z0-9] string 26
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// with the padding stripped off.
func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String()
}

func NewRandomString(length int) string {
	var b bytes.Buffer
	str := make([]byte, length+8)
	rand.Read(str)
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(str)
	encoder.Close()
	b.Truncate(length) // removes the '==' padding
	return b.String()
}

// GetMillis is a convience method to get milliseconds since epoch.
func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// MapToJson converts a map to a json string
func MapToJson(objmap map[string]string) string {
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

// MapToJson converts a map to a json string
func MapBoolToJson(objmap map[string]bool) string {
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

// MapFromJson will decode the key/value pair map
func MapFromJson(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)

	var objmap map[string]string
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]string)
	} else {
		return objmap
	}
}

// MapFromJson will decode the key/value pair map
func MapBoolFromJson(data io.Reader) map[string]bool {
	decoder := json.NewDecoder(data)

	var objmap map[string]bool
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]bool)
	} else {
		return objmap
	}
}

func ArrayToJson(objmap []string) string {
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ArrayFromJson(data io.Reader) []string {
	decoder := json.NewDecoder(data)

	var objmap []string
	if err := decoder.Decode(&objmap); err != nil {
		return make([]string, 0)
	} else {
		return objmap
	}
}

func ArrayFromInterface(data interface{}) []string {
	stringArray := []string{}

	dataArray, ok := data.([]interface{})
	if !ok {
		return stringArray
	}

	for _, v := range dataArray {
		if str, ok := v.(string); ok {
			stringArray = append(stringArray, str)
		}
	}

	return stringArray
}

func StringInterfaceToJson(objmap map[string]interface{}) string {
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func StringInterfaceFromJson(data io.Reader) map[string]interface{} {
	decoder := json.NewDecoder(data)

	var objmap map[string]interface{}
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]interface{})
	} else {
		return objmap
	}
}

func StringToJson(s string) string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func StringFromJson(data io.Reader) string {
	decoder := json.NewDecoder(data)

	var s string
	if err := decoder.Decode(&s); err != nil {
		return ""
	} else {
		return s
	}
}

func IsLower(s string) bool {
	if strings.ToLower(s) == s {
		return true
	}

	return false
}

func IsValidEmail(email string) bool {

	if !IsLower(email) {
		return false
	}

	if _, err := mail.ParseAddress(email); err == nil {
		return true
	}

	return false
}

func IsValidChannelIdentifier(s string) bool {

	if !IsValidAlphaNumHyphenUnderscore(s, true) {
		return false
	}
	return true
}

func IsValidAlphaNum(s string) bool {
	validAlphaNum := regexp.MustCompile(`^[a-z0-9]+([a-z\-0-9]+|(__)?)[a-z0-9]+$`)

	return validAlphaNum.MatchString(s)
}

func IsValidAlphaNumHyphenUnderscore(s string, withFormat bool) bool {
	if withFormat {
		validAlphaNumHyphenUnderscore := regexp.MustCompile(`^[a-z0-9]+([a-z\-\_0-9]+|(__)?)[a-z0-9]+$`)
		return validAlphaNumHyphenUnderscore.MatchString(s)
	}

	validSimpleAlphaNumHyphenUnderscore := regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)
	return validSimpleAlphaNumHyphenUnderscore.MatchString(s)
}

func Etag(parts ...interface{}) string {
	etag := CurrentVersion
	for _, part := range parts {
		etag += fmt.Sprintf(".%v", part)
	}

	return etag
}

var validHashtag = regexp.MustCompile(`^(#\pL[\pL\d\-_.]*[\pL\d])$`)
var puncStart = regexp.MustCompile(`^[^\pL\d\s#]+`)
var hashtagStart = regexp.MustCompile(`^#{2,}`)
var puncEnd = regexp.MustCompile(`[^\pL\d\s]+$`)

func ParseHashtags(text string) (string, string) {
	words := strings.Fields(text)

	hashtagString := ""
	plainString := ""
	for _, word := range words {
		// trim off surrounding punctuation
		word = puncStart.ReplaceAllString(word, "")
		word = puncEnd.ReplaceAllString(word, "")

		// and remove extra pound #s
		word = hashtagStart.ReplaceAllString(word, "#")

		if validHashtag.MatchString(word) {
			hashtagString += " " + word
		} else {
			plainString += " " + word
		}
	}

	if len(hashtagString) > 1000 {
		hashtagString = hashtagString[:999]
		lastSpace := strings.LastIndex(hashtagString, " ")
		if lastSpace > -1 {
			hashtagString = hashtagString[:lastSpace]
		} else {
			hashtagString = ""
		}
	}

	return strings.TrimSpace(hashtagString), strings.TrimSpace(plainString)
}

func IsFileExtImage(ext string) bool {
	ext = strings.ToLower(ext)
	for _, imgExt := range IMAGE_EXTENSIONS {
		if ext == imgExt {
			return true
		}
	}
	return false
}

func GetImageMimeType(ext string) string {
	ext = strings.ToLower(ext)
	if len(IMAGE_MIME_TYPES[ext]) == 0 {
		return "image"
	} else {
		return IMAGE_MIME_TYPES[ext]
	}
}

func ClearMentionTags(post string) string {
	post = strings.Replace(post, "<mention>", "", -1)
	post = strings.Replace(post, "</mention>", "", -1)
	return post
}

var UrlRegex = regexp.MustCompile(`^((?:[a-z]+:\/\/)?(?:(?:[a-z0-9\-]+\.)+(?:[a-z]{2}|aero|arpa|biz|com|coop|edu|gov|info|int|jobs|mil|museum|name|nato|net|org|pro|travel|local|internal))(:[0-9]{1,5})?(?:\/[a-z0-9_\-\.~]+)*(\/([a-z0-9_\-\.]*)(?:\?[a-z0-9+_~\-\.%=&amp;]*)?)?(?:#[a-zA-Z0-9!$&'()*+.=-_~:@/?]*)?)(?:\s+|$)$`)
var PartialUrlRegex = regexp.MustCompile(`/([A-Za-z0-9]{26})/([A-Za-z0-9]{26})/((?:[A-Za-z0-9]{26})?.+(?:\.[A-Za-z0-9]{3,})?)`)

func IsValidHttpUrl(rawUrl string) bool {
	if strings.Index(rawUrl, "http://") != 0 && strings.Index(rawUrl, "https://") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return false
	}

	return true
}

func IsValidHttpsUrl(rawUrl string) bool {
	if strings.Index(rawUrl, "https://") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return false
	}

	return true
}

func IsValidTurnOrStunServer(rawUri string) bool {
	if strings.Index(rawUri, "turn:") != 0 && strings.Index(rawUri, "stun:") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawUri); err != nil {
		return false
	}

	return true
}

func IsSafeLink(link *string) bool {
	if link != nil {
		if IsValidHttpUrl(*link) {
			return true
		} else if strings.HasPrefix(*link, "/") {
			return true
		} else {
			return false
		}
	}

	return true
}

func IsValidWebsocketUrl(rawUrl string) bool {
	if strings.Index(rawUrl, "ws://") != 0 && strings.Index(rawUrl, "wss://") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return false
	}

	return true
}

func AbsPath(filePath string) string {
	fullPath, _ := filepath.Abs(filePath)
	return fullPath
}

func CheckFile(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	} else {
		return true
	}
}
