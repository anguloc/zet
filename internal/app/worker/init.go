package worker

import "github.com/anguloc/zet/pkg/cron"

func Init() {
	_ = cron.NewCron().RegisterJob("3 * * * *", newDmhyRss())
	_ = cron.NewCron().RegisterJob("5 * * * *", newDmhyParse())
}
