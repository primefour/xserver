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
