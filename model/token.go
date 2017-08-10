package model

import (
	"github.com/primefour/xserver/utils"
	"net/http"
)

const (
	TOKEN_SIZE            = 64
	MAX_TOKEN_EXIPRY_TIME = 1000 * 60 * 60 * 24 // 24 hour
)

type Token struct {
	Token    string
	CreateAt int64
	Type     string
	Extra    string
}

func NewToken(tokentype, extra string) *Token {
	return &Token{
		Token:    utils.NewRandomString(TOKEN_SIZE),
		CreateAt: utils.GetMillis(),
		Type:     tokentype,
		Extra:    extra,
	}
}

func (t *Token) IsValid() *utils.AppError {
	if len(t.Token) != TOKEN_SIZE {
		return utils.NewAppError("Token.IsValid", "model.token.is_valid.size", nil, "", http.StatusInternalServerError)
	}

	if t.CreateAt == 0 {
		return utils.NewAppError("Token.IsValid", "model.token.is_valid.expiry", nil, "", http.StatusInternalServerError)
	}
	return nil
}
