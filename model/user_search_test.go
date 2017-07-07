package model

import (
	"strings"
	"testing"
)

func TestUserSearchJson(t *testing.T) {
	userSearch := UserSearch{Term: NewId(), TeamId: NewId()}
	json := userSearch.ToJson()
	ruserSearch := UserSearchFromJson(strings.NewReader(json))

	if userSearch.Term != ruserSearch.Term {
		t.Fatal("Terms do not match")
	}
}
