package einterfaces

import (
	"time"

	"github.com/primefour/xserver/model"
)

type ElasticsearchInterface interface {
	Start() *model.AppError
	IndexPost(post *model.Post, teamId string) *model.AppError
	SearchPosts(channels *model.ChannelList, searchParams []*model.SearchParams) ([]string, *model.AppError)
	DeletePost(post *model.Post) *model.AppError
	//TestConfig(cfg *model.Config) *model.AppError
	PurgeIndexes() *model.AppError
	DataRetentionDeleteIndexes(cutoff time.Time) *model.AppError
}
