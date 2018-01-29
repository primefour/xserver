package jobs

import (
	"github.com/primefour/xserver/model"
)

type LdapSyncInterface interface {
	MakeWorker() model.Worker
	MakeScheduler() model.Scheduler
}
