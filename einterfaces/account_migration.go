package einterfaces

import "github.com/primefour/xserver/model"

type AccountMigrationInterface interface {
	MigrateToLdap(fromAuthService string, forignUserFieldNameToMatch string, force bool) *model.AppError
}
