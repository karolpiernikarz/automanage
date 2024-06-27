package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/cache/filecache"
	"github.com/karolpiernikarz/automanage/helpers"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/karolpiernikarz/automanage/services/redis"
	"github.com/karolpiernikarz/automanage/utils"
	"github.com/spf13/viper"
)

type SnsMessage struct {
	Type           string `json:"Type"`
	MessageId      string `json:"MessageId"`
	Token          string `json:"Token"`
	TopicArn       string `json:"TopicArn"`
	Message        string `json:"Message"`
	SubscribeURL   string `json:"SubscribeURL"`
	Timestamp      string `json:"Timestamp"`
	Signature      string `json:"Signature"`
	SigningCertURL string `json:"SigningCertURL"`
	UnsubscribeURL string `json:"UnsubscribeURL"`
}

func AwsSesWebhook(c *gin.Context) {
	var snsMessage SnsMessage

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	fmt.Printf("Raw request body: %s\n", string(body))

	err = json.Unmarshal(body, &snsMessage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse SNS message"})
		return
	}

	fmt.Printf("Parsed SNS message: %+v\n", snsMessage)

	if snsMessage.Type == "SubscriptionConfirmation" {
		resp, err := http.Get(snsMessage.SubscribeURL)
		if err != nil {
			fmt.Println("Error confirming subscription:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm subscription"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Println("Successfully confirmed subscription")
			c.JSON(http.StatusOK, gin.H{"message": "Subscription confirmed"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm subscription"})
		}
		return
	}

	// Handle regular SES feedback
	var feedback models.AmazonSesFeedback
	err = json.Unmarshal(body, &feedback)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	fmt.Printf("Received feedback: %+v\n", feedback)
	fmt.Printf("Feedback EventType: %s\n", feedback.EventType)

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

	fmt.Printf("Received feedback: %+v\n", feedback)
	fmt.Printf("Feedback EventType: %s\n", feedback.EventType)

	switch feedback.EventType {
	case "Bounce":
		fmt.Println("Processing Bounce event")
		fmt.Printf("Bounce details: %+v\n", feedback)

		dcFiles, err := filecache.DockerComposeFiles()
		if err != nil {
			fmt.Println("Error fetching Docker Compose files:", err)
			break
		}
		fmt.Printf("Docker Compose files: %+v\n", dcFiles)

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
			fmt.Println("Restaurant not found for source:", feedback.Mail.Source)
			c.JSON(404, gin.H{
				"status":  "error",
				"message": "restaurant not found",
			})
			return
		}
		fmt.Println("Found restaurant ID:", restaurantId)

		order, err := helpers.GetLastOrderByEmail(restaurantId, feedback.Mail.Destination[0])
		if err != nil {
			fmt.Println("Order not found for email:", feedback.Mail.Destination[0])
			c.JSON(404, gin.H{
				"status":  "error",
				"message": "order not found",
			})
			return
		}
		fmt.Printf("Order details: %+v\n", order)

		// send Slack notification
		err = utils.SendSlackWebhookMessage(CreateSlackMessageSnsBounce(feedback, order, restaurantInfo), viper.GetString("slack.webhook"))
		if err != nil {
			fmt.Println("Error sending Slack notification:", err)
		} else {
			fmt.Println("Slack notification sent successfully")
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
