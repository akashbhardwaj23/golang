package cron

import (
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron            *cron.Cron
	predefinedUsers []string
}
