package cronJobs

import (
	"schedule-events/packages/checkEvents"
	"schedule-events/packages/common/config"

	"github.com/robfig/cron/v3"
)

func getCheckJobDuration() string {
	return config.GoDotEnvVariable("CRON_TIME")
}

func startJobCheckProcess(cron *cron.Cron) {
	cron.AddFunc(getCheckJobDuration(), func() {
		checkEvents.CheckEvents()
	})
}

func StartCron() {
	c := cron.New()
	startJobCheckProcess(c)
	c.Start()
}
