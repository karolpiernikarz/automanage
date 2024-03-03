package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/cache/filecache"
	"github.com/karolpiernikarz/automanage/helpers"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/karolpiernikarz/automanage/services/redis"
	"github.com/karolpiernikarz/automanage/utils"
	"github.com/spf13/viper"
)

func AwsSesWebhook(c *gin.Context) {

	var feedback models.AmazonSesFeedback
	err := c.BindJSON(&feedback)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if feedback.Mail.MessageId == "" {
		c.JSON(400, gin.H{
			"message": "message id is required",
		})
		return
	}

	// if amazon_ses_+feedback.Mail.Destination[0] exist in cache, return 200
	if redis.IsExist("amazon_ses_" + feedback.Mail.Destination[0]) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
		fmt.Println("amazon_ses_" + feedback.Mail.Destination[0] + " exist in cache")
		return
	} else {
		// if amazon_ses_+feedback.Mail.MessageId does not exist in cache, set it to cache for 1 day
		err = redis.Set("amazon_ses_"+feedback.Mail.Destination[0], "1", 24*time.Hour)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("amazon_ses_" + feedback.Mail.Destination[0] + " does not exist in cache")
	}

	// if amazon_ses_+feedback.Mail.MessageId exist in cache, return 200
	if redis.IsExist("amazon_ses_" + feedback.Mail.MessageId) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
		fmt.Println("amazon_ses_" + feedback.Mail.MessageId + " exist in cache")
		return
	} else {
		// if amazon_ses_+feedback.Mail.MessageId does not exist in cache, set it to cache for 30 day
		err = redis.Set("amazon_ses_"+feedback.Mail.MessageId, "1", 24*time.Hour*30)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("amazon_ses_" + feedback.Mail.MessageId + " does not exist in cache")
	}

	switch feedback.EventType {
	case "Bounce":
		dcFiles, err := filecache.DockerComposeFiles()
		if err != nil {
			fmt.Println(err)
		}
		restaurantInfo := models.DockerCompose{}
		restaurantId := ""
		for _, dcFile := range dcFiles {
			if dcFile.Services.App.Environment.MAIL_FROM_ADDRESS == feedback.Mail.Source {
				restaurantId = dcFile.Services.App.Environment.PHP_POOL_NAME
				restaurantInfo = dcFile
				break
			}
		}
		if restaurantId == "" {
			fmt.Println("restaurant not found")
			c.JSON(404, gin.H{
				"status":  "error",
				"message": "restaurant not found",
			})
			return
		}
		order, err := helpers.GetLastOrderByEmail(restaurantId, feedback.Mail.Destination[0])
		if err != nil {
			fmt.Println("order not found")
			c.JSON(404, gin.H{
				"status":  "error",
				"message": "order not found",
			})
			return
		}
		// send Slack notification
		err = utils.SendSlackWebhookMessage(CreateSlackMessageSnsBounce(feedback, order, restaurantInfo), viper.GetString("slack.webhook"))
		if err != nil {
			fmt.Println(err)
		}

	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func CreateSlackMessageSnsBounce(feedback models.AmazonSesFeedback, order models.RestaurantOrders, restaurant models.DockerCompose) models.SlackWebhookMessage {
	var slackMessage models.SlackWebhookMessage
	ordertype := utils.GetOrderTypeString(order.Type)
	ispreorder := utils.GetDeliveryTypeString(order.IsPreOrder)
	isopen := helpers.IsRestaurantOpen(restaurant.Services.App.Environment.PHP_POOL_NAME)
	openmessage := ""
	if isopen {
		openmessage = "Restaurant is Open"
	} else {
		openmessage = "Restaurant is Closed"
	}
	slackMessage.Text = "Email Bounce\n" + feedback.Mail.Destination[0] + "\n" + order.Customer.Data().Phone + "\n" + order.Customer.Data().Name + " " + order.Customer.Data().Surname + "\n" + restaurant.Services.App.Environment.APP_URL +
		"/order/" + order.PaymentId + "/status" + "\n" + feedback.Bounce.BouncedRecipients[0].DiagnosticCode + "\n" + "Order Number: " + order.OrderNumber + "\n" +
		order.Date + "\n" + ordertype + "-" + ispreorder + "\n" + openmessage
	return slackMessage
}
