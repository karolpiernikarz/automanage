package restaurantapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/helpers"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/karolpiernikarz/automanage/utils"
	"github.com/spf13/viper"
)

func Order(c *gin.Context) {
	messageType := c.PostForm("type")
	switch messageType {
	case "created":
		if c.GetString("restaurantid") == "test" {
			orderCreatedTest(c)
			return
		} else {
			orderCreated(c)
			return
		}
	case "updated":
		fmt.Println("order updated")
		return
	}
	c.JSON(400, "bad request")
}

func orderCreated(c *gin.Context) {
	restaurantId := c.GetString("restaurantid")
	restaurant := helpers.GetRestaurant(restaurantId)
	boxActive, err := utils.IsOrderBoxActive(fmt.Sprintf("%v", restaurant.Info.Data().TerminalId))
	if err != nil {
		c.JSON(500, err)
		return
	}
	orderData := c.PostForm("order")

	var order models.RestaurantOrders
	err = json.Unmarshal([]byte(orderData), &order)
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(422, err)
		return
	}
	open := helpers.IsRestaurantOpen(restaurantId)
	if open {
		if !boxActive {
			err := utils.SendSlackWebhookMessage(createSlackMessageForOrderBoxNotActive(order, restaurant), viper.GetString("slack.webhook"))
			if err != nil {
				fmt.Println(err)
			}
		} else {
			go time.AfterFunc(5*time.Minute, func() {
				pendingOrderCheck(restaurant, order.OrderNumber)
			})
		}
	}

	var message models.NtfyMessage
	message.Topic = "test"
	message.Title = "New Order"
	message.Message = restaurant.Website + "/order/" + order.PaymentId + "/status"
	err = utils.SendNtfyNotification(message)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "ok")
}

func pendingOrderCheck(restaurant models.Restaurant, orderNumber string) {
	order := helpers.GetRestaurantOrderFromOrderNumber(strconv.Itoa(int(restaurant.ID)), orderNumber)
	if order.Status == 1 {
		err := utils.SendSlackWebhookMessage(CreateSlackMessageForOrderNotAccepted(order, restaurant), viper.GetString("slack.webhook"))
		if err != nil {
			log.WithFields(log.Fields{"restaurantId": restaurant.ID}).Warn("Error while sending email")
			fmt.Println(err)
		}
	}
}

func CreateSlackMessageForOrderNotAccepted(order models.RestaurantOrders, restaurant models.Restaurant) models.SlackWebhookMessage {
	var slackMessage models.SlackWebhookMessage
	slackMessage.Text = "*Order Not Accepted In 5 minutes*\n" + "Order with the number " + order.OrderNumber + " is not accepted in 5 minutes" + "\n" + restaurant.Website + "/order/" + order.PaymentId + "/status" + "" +
		"\n" + "*Customer Details:* " + order.Customer.Data().Name + " " + order.Customer.Data().Phone + " " + order.Address.Data().Detail + "" +
		"\n" + "*Restaurant Information:* " + restaurant.Name + " " + restaurant.Phone + " " + restaurant.Address
	return slackMessage
}

func createSlackMessageForOrderBoxNotActive(order models.RestaurantOrders, restaurant models.Restaurant) models.SlackWebhookMessage {
	var slackMessage models.SlackWebhookMessage
	slackMessage.Text = "Order Box Not Active"
	slackMessage.Blocks = make([]models.SlackWebhookMessageBlocks, 2)

	// Order Information
	slackMessage.Blocks[0].Type = "section"
	slackMessage.Blocks[0].Text.Type = "mrkdwn"
	slackMessage.Blocks[0].Text.Text = "*Orderbox Not Active For : " + restaurant.Name + "*"
	slackMessage.Blocks[0].Fields = make([]models.SlackWebhookMessageBlocksField, 2)
	i := 0
	slackMessage.Blocks[0].Fields[i].Type = "mrkdwn"
	slackMessage.Blocks[0].Fields[i].Text = "*Order URL:*\n" + restaurant.Website + "/order/" + order.PaymentId + "/status"
	i++
	deliveryAt, _ := time.Parse(time.RFC3339, order.Date)
	loc, _ := time.LoadLocation(viper.GetString("app.timezone"))
	deliveryAt = deliveryAt.In(loc)
	slackMessage.Blocks[0].Fields[i].Type = "mrkdwn"
	slackMessage.Blocks[0].Fields[i].Text = "*Delivery Time:*\n" + deliveryAt.Format("02-01-06 15:04")
	// Customer Information
	i = 0
	slackMessage.Blocks[1].Type = "section"
	slackMessage.Blocks[1].Text.Type = "mrkdwn"
	slackMessage.Blocks[1].Text.Text = "*Details:*"
	slackMessage.Blocks[1].Fields = make([]models.SlackWebhookMessageBlocksField, 2)
	slackMessage.Blocks[1].Fields[i].Type = "mrkdwn"
	slackMessage.Blocks[1].Fields[i].Text = "*Customer Details:*\n" + order.Customer.Data().Name + "\n" + order.Customer.Data().Phone + "\n" + order.Address.Data().Detail
	i++
	slackMessage.Blocks[1].Fields[i].Type = "mrkdwn"
	slackMessage.Blocks[1].Fields[i].Text = "*Restaurant Information:*\n" + restaurant.Name + "\n" + restaurant.Phone + "\n" + restaurant.Address

	return slackMessage
}

func orderCreatedTest(c *gin.Context) {
	boxActive, err := utils.IsOrderBoxActive(viper.GetString("test.orderboxid"))
	if err != nil {
		c.JSON(500, err)
		return
	}
	if !boxActive {
		fmt.Println("test orderbox not active")
	}
	var message models.NtfyMessage
	message.Topic = "test"
	message.Title = "New Order"
	message.Message = "new order"
	err = utils.SendNtfyNotification(message)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "test order")
}
