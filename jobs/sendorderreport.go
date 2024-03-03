package jobs

import (
	"time"

	"github.com/karolpiernikarz/automanage/email"
	"github.com/karolpiernikarz/automanage/helpers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CreateOrderReportAndSend() {
	orders := helpers.GetOrdersFromAllRestaurants()
	filepath, err := helpers.CreateXlsxFileFromAllOrders(orders)
	if err != nil {
		log.WithField("error", err)
	}
	err = email.SendWithAttachment(viper.GetString("app.notifyemail"), "order reports", "order reports", filepath, filepath, false)
	if err != nil {
		log.WithField("error", err)
	}
}

func CreateWeeklyOrderReportAndSend() {
	currentTime := time.Now()
	weekStart := currentTime.AddDate(0, 0, int(time.Monday-currentTime.Weekday()-7))
	weekEnd := weekStart.AddDate(0, 0, 6)
	orders := helpers.GetOrdersFromAllRestaurantsWithTime(weekStart, weekEnd)
	filepath, err := helpers.CreateXlsxFileFromAllOrders(orders)
	if err != nil {
		log.WithField("error", err)
	}
	err = email.SendWithAttachment(viper.GetString("app.notifyemail"), "order reports", "order reports", filepath, filepath, false)
	if err != nil {
		log.WithField("error", err)
	}
}
