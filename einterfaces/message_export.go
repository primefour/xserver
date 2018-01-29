package einterfaces

import (
	"context"

	"github.com/primefour/xserver/model"
)

type MessageExportInterface interface {
	StartSynchronizeJob(ctx context.Context, exportFromTimestamp int64) (*model.Job, *model.AppError)
}
