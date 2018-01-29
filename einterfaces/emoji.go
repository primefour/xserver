package einterfaces

import (
	"github.com/primefour/xserver/model"
)

type EmojiInterface interface {
	CanUserCreateEmoji(string, []*model.TeamMember) bool
}
