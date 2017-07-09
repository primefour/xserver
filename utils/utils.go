package utils

import (
	"github.com/primefour/xserver/model"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	LOWERCASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS           = "0123456789"
	SYMBOLS           = " !\"\\#$%&'()*+,-./:;<=>?@[]^_`|~"
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
	address := r.Header.Get(model.HEADER_FORWARDED)

	if len(address) == 0 {
		address = r.Header.Get(model.HEADER_REAL_IP)
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
	return rand.Intn((r.End-r.Begin)+1) + r.Begin
}
