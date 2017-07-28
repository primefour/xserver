package utils

import (
	"testing"
	"time"
)

func TestStringArrayIntersection(t *testing.T) {
	a := []string{
		"abc",
		"def",
		"ghi",
	}
	b := []string{
		"jkl",
	}
	c := []string{
		"def",
	}

	if len(StringArrayIntersection(a, b)) != 0 {
		t.Fatal("should be 0")
	}

	if len(StringArrayIntersection(a, c)) != 1 {
		t.Fatal("should be 1")
	}
}

func TestRemoveDuplicatesFromStringArray(t *testing.T) {
	a := []string{
		"a",
		"b",
		"a",
		"a",
		"b",
		"c",
		"a",
	}

	if len(RemoveDuplicatesFromStringArray(a)) != 3 {
		t.Fatal("should be 3")
	}
}

func TestUrlEncode(t *testing.T) {

	toEncode := "testing 1 2 3"
	encoded := UrlEncode(toEncode)

	if encoded != "testing%201%202%203" {
		t.Log(encoded)
		t.Fatal("should be equal")
	}

	toEncode = "testing123"
	encoded = UrlEncode(toEncode)

	if encoded != "testing123" {
		t.Log(encoded)
		t.Fatal("should be equal")
	}

	toEncode = "testing$#~123"
	encoded = UrlEncode(toEncode)

	if encoded != "testing%24%23~123" {
		t.Log(encoded)
		t.Fatal("should be equal")
	}
}

var format = "2006-01-02 15:04:05.000000000"

func TestMillisFromTime(t *testing.T) {
	input, _ := time.Parse(format, "2015-01-01 12:34:00.000000000")
	actual := MillisFromTime(input)
	expected := int64(1420115640000)

	if actual != expected {
		t.Fatalf("TestMillisFromTime failed, %v=%v", expected, actual)
	}
}

func TestYesterday(t *testing.T) {
	actual := Yesterday()
	expected := time.Now().AddDate(0, 0, -1)

	if actual.Year() != expected.Year() || actual.Day() != expected.Day() || actual.Month() != expected.Month() {
		t.Fatalf("TestYesterday failed, %v=%v", expected, actual)
	}
}

func TestStartOfDay(t *testing.T) {
	input, _ := time.Parse(format, "2015-01-01 12:34:00.000000000")
	actual := StartOfDay(input)
	expected, _ := time.Parse(format, "2015-01-01 00:00:00.000000000")

	if actual != expected {
		t.Fatalf("TestStartOfDay failed, %v=%v", expected, actual)
	}
}

func TestEndOfDay(t *testing.T) {
	input, _ := time.Parse(format, "2015-01-01 12:34:00.000000000")
	actual := EndOfDay(input)
	expected, _ := time.Parse(format, "2015-01-01 23:59:59.999999999")

	if actual != expected {
		t.Fatalf("TestEndOfDay failed, %v=%v", expected, actual)
	}
}

func TestNewId(t *testing.T) {
	for i := 0; i < 1000; i++ {
		id := NewId()
		if len(id) > 26 {
			t.Fatal("ids shouldn't be longer than 26 chars")
		}
	}
}

func TestRandomString(t *testing.T) {
	for i := 0; i < 1000; i++ {
		r := NewRandomString(32)
		if len(r) != 32 {
			t.Fatal("should be 32 chars")
		}
	}
}

func TestAppError(t *testing.T) {
	err := NewLocAppError("TestAppError", "message", nil, "")
	json := err.ToJson()
	rerr := AppErrorFromJson(strings.NewReader(json))
	if err.Message != rerr.Message {
		t.Fatal()
	}

	t.Log(err.Error())
}

func TestAppErrorJunk(t *testing.T) {
	rerr := AppErrorFromJson(strings.NewReader("<html><body>This is a broken test</body></html>"))
	if "body: <html><body>This is a broken test</body></html>" != rerr.DetailedError {
		t.Fatal()
	}
}

func TestMapJson(t *testing.T) {

	m := make(map[string]string)
	m["id"] = "test_id"
	json := MapToJson(m)

	rm := MapFromJson(strings.NewReader(json))

	if rm["id"] != "test_id" {
		t.Fatal("map should be valid")
	}

	rm2 := MapFromJson(strings.NewReader(""))
	if len(rm2) > 0 {
		t.Fatal("make should be ivalid")
	}
}

func TestValidEmail(t *testing.T) {
	if !IsValidEmail("corey+test@hulen.com") {
		t.Error("email should be valid")
	}

	if IsValidEmail("@corey+test@hulen.com") {
		t.Error("should be invalid")
	}
}

func TestValidLower(t *testing.T) {
	if !IsLower("corey+test@hulen.com") {
		t.Error("should be valid")
	}

	if IsLower("Corey+test@hulen.com") {
		t.Error("should be invalid")
	}
}

func TestEtag(t *testing.T) {
	etag := Etag("hello", 24)
	if len(etag) <= 0 {
		t.Fatal()
	}
}

var hashtags = map[string]string{
	"#test":           "#test",
	"test":            "",
	"#test123":        "#test123",
	"#123test123":     "",
	"#test-test":      "#test-test",
	"#test?":          "#test",
	"hi #there":       "#there",
	"#bug #idea":      "#bug #idea",
	"#bug or #gif!":   "#bug #gif",
	"#hüllo":          "#hüllo",
	"#?test":          "",
	"#-test":          "",
	"#yo_yo":          "#yo_yo",
	"(#brakets)":      "#brakets",
	")#stekarb(":      "#stekarb",
	"<#less_than<":    "#less_than",
	">#greater_than>": "#greater_than",
	"-#minus-":        "#minus",
	"_#under_":        "#under",
	"+#plus+":         "#plus",
	"=#equals=":       "#equals",
	"%#pct%":          "#pct",
	"&#and&":          "#and",
	"^#hat^":          "#hat",
	"##brown#":        "#brown",
	"*#star*":         "#star",
	"|#pipe|":         "#pipe",
	":#colon:":        "#colon",
	";#semi;":         "#semi",
	"#Mötley;":        "#Mötley",
	".#period.":       "#period",
	"¿#upside¿":       "#upside",
	"\"#quote\"":      "#quote",
	"/#slash/":        "#slash",
	"\\#backslash\\":  "#backslash",
	"#a":              "",
	"#1":              "",
	"foo#bar":         "",
}

func TestParseHashtags(t *testing.T) {
	for input, output := range hashtags {
		if o, _ := ParseHashtags(input); o != output {
			t.Fatal("failed to parse hashtags from input=" + input + " expected=" + output + " actual=" + o)
		}
	}
}

func TestIsValidAlphaNum(t *testing.T) {
	cases := []struct {
		Input  string
		Result bool
	}{
		{
			Input:  "test",
			Result: true,
		},
		{
			Input:  "test-name",
			Result: true,
		},
		{
			Input:  "test--name",
			Result: true,
		},
		{
			Input:  "test__name",
			Result: true,
		},
		{
			Input:  "-",
			Result: false,
		},
		{
			Input:  "__",
			Result: false,
		},
		{
			Input:  "test-",
			Result: false,
		},
		{
			Input:  "test--",
			Result: false,
		},
		{
			Input:  "test__",
			Result: false,
		},
		{
			Input:  "test:name",
			Result: false,
		},
	}

	for _, tc := range cases {
		actual := IsValidAlphaNum(tc.Input)
		if actual != tc.Result {
			t.Fatalf("case: %v\tshould returned: %#v", tc, tc.Result)
		}
	}
}

func TestIsValidAlphaNumHyphenUnderscore(t *testing.T) {
	casesWithFormat := []struct {
		Input  string
		Result bool
	}{
		{
			Input:  "test",
			Result: true,
		},
		{
			Input:  "test-name",
			Result: true,
		},
		{
			Input:  "test--name",
			Result: true,
		},
		{
			Input:  "test__name",
			Result: true,
		},
		{
			Input:  "test_name",
			Result: true,
		},
		{
			Input:  "test_-name",
			Result: true,
		},
		{
			Input:  "-",
			Result: false,
		},
		{
			Input:  "__",
			Result: false,
		},
		{
			Input:  "test-",
			Result: false,
		},
		{
			Input:  "test--",
			Result: false,
		},
		{
			Input:  "test__",
			Result: false,
		},
		{
			Input:  "test:name",
			Result: false,
		},
	}

	for _, tc := range casesWithFormat {
		actual := IsValidAlphaNumHyphenUnderscore(tc.Input, true)
		if actual != tc.Result {
			t.Fatalf("case: %v\tshould returned: %#v", tc, tc.Result)
		}
	}

	casesWithoutFormat := []struct {
		Input  string
		Result bool
	}{
		{
			Input:  "test",
			Result: true,
		},
		{
			Input:  "test-name",
			Result: true,
		},
		{
			Input:  "test--name",
			Result: true,
		},
		{
			Input:  "test__name",
			Result: true,
		},
		{
			Input:  "test_name",
			Result: true,
		},
		{
			Input:  "test_-name",
			Result: true,
		},
		{
			Input:  "-",
			Result: true,
		},
		{
			Input:  "_",
			Result: true,
		},
		{
			Input:  "test-",
			Result: true,
		},
		{
			Input:  "test--",
			Result: true,
		},
		{
			Input:  "test__",
			Result: true,
		},
		{
			Input:  ".",
			Result: false,
		},

		{
			Input:  "test,",
			Result: false,
		},
		{
			Input:  "test:name",
			Result: false,
		},
	}

	for _, tc := range casesWithoutFormat {
		actual := IsValidAlphaNumHyphenUnderscore(tc.Input, false)
		if actual != tc.Result {
			t.Fatalf("case: '%v'\tshould returned: %#v", tc.Input, tc.Result)
		}
	}
}
