package einterfaces

import (
	"github.com/primefour/xserver/model"
	"mime/multipart"
)

type BrandInterface interface {
	SaveBrandImage(*multipart.FileHeader) *model.AppError
	GetBrandImage() ([]byte, *model.AppError)
}
