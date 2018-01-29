package einterfaces

import (
	"github.com/primefour/xserver/model"
)

type DataRetentionInterface interface {
	GetPolicy() (*model.DataRetentionPolicy, *model.AppError)
}
