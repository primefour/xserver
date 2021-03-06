package einterfaces

import (
	"github.com/primefour/xserver/model"
)

type MfaInterface interface {
	GenerateSecret(user *model.User) (string, []byte, *model.AppError)
	Activate(user *model.User, token string) *model.AppError
	Deactivate(userId string) *model.AppError
	ValidateToken(secret, token string) (bool, *model.AppError)
}

var theMfaInterface MfaInterface

func RegisterMfaInterface(newInterface MfaInterface) {
	theMfaInterface = newInterface
}

func GetMfaInterface() MfaInterface {
	return theMfaInterface
}
