package einterfaces

import (
	"github.com/primefour/xserver/model"
)

type SamlInterface interface {
	ConfigureSP() *model.AppError
	BuildRequest(relayState string) (*model.SamlAuthRequest, *model.AppError)
	DoLogin(encodedXML string, relayState map[string]string) (*model.User, *model.AppError)
	GetMetadata() (string, *model.AppError)
}
