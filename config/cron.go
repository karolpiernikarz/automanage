package config

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/karolpiernikarz/automanage/jobs"
)

func InitCron() {
	cron := gocron.NewScheduler(time.UTC)
	cron.Every("1m").Do(func() {
		err := jobs.CreateOrderBoxCache()
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	cron.Every("5m").Do(func() {
		jobs.CreateRestaurantsCache()
	})
	cron.StartAsync()
}
