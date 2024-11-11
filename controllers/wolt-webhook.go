package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/helpers"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/karolpiernikarz/automanage/services/redis"
	"github.com/karolpiernikarz/automanage/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func WoltWebHook(c *gin.Context) {
	var tokenString models.WoltWebHookBody
	err := c.BindJSON(&tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	log.Info(tokenString)

	// parse jwt
	claims, err := utils.ParseToken(tokenString.Token, viper.GetString("wolt.jwtsecret"))

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse jwt"})
		return
	}

	log.Info(claims)
	var body models.WoltWebHook

	// unmarshal it to struct
	jsonString, err := json.Marshal(claims)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal jwt claims"})
		return
	}

	err = json.Unmarshal(jsonString, &body)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal jwt claims"})
		return
	}

	restaurant := helpers.GetRestaurantFromVenueId(body.Details.VenueId)
	if restaurant.ID == 0 {
		log.Error("wolt restaurant not found")
		message := models.SlackWebhookMessage{
			Text: "Wolt Restaurant Not Found" + "\n" + body.Details.Id + "\n" + body.Details.MerchantOrderReferenceId + "\n" + body.Details.VenueId,
		}
		utils.SendSlackWebhookMessage(message, viper.GetString("slack.sandbox"))
		c.JSON(200, "ok")
		return
	}

	order := helpers.GetRestaurantOrderFromOrderNumber(strconv.FormatUint(uint64(restaurant.ID), 10), body.Details.MerchantOrderReferenceId)
	// restaurantSettings := helpers.GetRestaurantSettingsByName(strconv.FormatUint(uint64(restaurant.ID), 10))

	if order.Id == 0 {
		log.Error("wolt order not found")
		message := models.SlackWebhookMessage{
			Text: "Wolt Order Not Found" + "\n" + body.Details.Id + "\n" + body.Details.MerchantOrderReferenceId + "\n" + restaurant.Name,
		}
		utils.SendSlackWebhookMessage(message, viper.GetString("slack.sandbox"))
		c.JSON(200, "ok")
		return
	}

	switch body.Type {
	case "order.pickup_eta_updated":
		redis.Set("order_"+strconv.FormatUint(uint64(restaurant.ID), 10)+"_"+body.Details.MerchantOrderReferenceId+"_pickupeta", body.Details.Pickup.Eta.String(), 0)
		// utils.SendSlackWebhookMessage(CreateWoltPickupEtaUpdatedMessage(body, restaurant, order, restaurantSettings), viper.GetString("slack.sandbox"))
	case "order.dropoff_eta_updated":
		redis.Set("order_"+strconv.FormatUint(uint64(restaurant.ID), 10)+"_"+body.Details.MerchantOrderReferenceId+"_dropoffeta", body.Details.Dropoff.Eta.String(), 0)
		// utils.SendSlackWebhookMessage(CreateWoltDropoffEtaUpdatedMessage(body), viper.GetString("slack.sandbox"))
	case "order.received":
		redis.Set("order_"+strconv.FormatUint(uint64(restaurant.ID), 10)+"_"+body.Details.MerchantOrderReferenceId+"_status", "received", 0)
		// utils.SendSlackWebhookMessage(CreateWoltOrderReceivedMessage(body), viper.GetString("slack.sandbox"))
	case "order.picked_up":
		redis.Set("order_"+strconv.FormatUint(uint64(restaurant.ID), 10)+"_"+body.Details.MerchantOrderReferenceId+"_status", "picked_up", 0)
		// utils.SendSlackWebhookMessage(CreateWoltOrderPickedUpMessage(body), viper.GetString("slack.sandbox"))
	case "order.delivered":
		redis.Set("order_"+strconv.FormatUint(uint64(restaurant.ID), 10)+"_"+body.Details.MerchantOrderReferenceId+"_status", "delivered", 0)
		// utils.SendSlackWebhookMessage(CreateWoltOrderDeliveredMessage(body), viper.GetString("slack.sandbox"))
	case "order.rejected":
		redis.Set("order_"+strconv.FormatUint(uint64(restaurant.ID), 10)+"_"+body.Details.MerchantOrderReferenceId+"_status", "rejected", 0)
		// utils.SendSlackWebhookMessage(CreateWoltOrderRejectedMessage(body), viper.GetString("slack.sandbox"))
	case "order.pickup_started":
		redis.Set("order_"+strconv.FormatUint(uint64(restaurant.ID), 10)+"_"+body.Details.MerchantOrderReferenceId+"_status", "pickup_started", 0)
		// utils.SendSlackWebhookMessage(CreateWoltOrderPickupStartedMessage(body), viper.GetString("slack.sandbox"))
	case "order.customer_no_show":
		// utils.SendSlackWebhookMessage(CreateWoltOrderCustomerNoShowMessage(body), viper.GetString("slack.sandbox"))
	case "order.location_updated":
		// utils.SendSlackWebhookMessage(CreateWoltOrderLocationUpdatedMessage(body), viper.GetString("slack.sandbox"))
	}

	c.JSON(200, "ok")
}

func WoltWebHookTest(c *gin.Context) {
	var tokenString models.WoltWebHookBody
	err := c.BindJSON(&tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	log.Info(tokenString)

	// parse jwt
	claims, err := utils.ParseToken(tokenString.Token, viper.GetString("wolt.jwtsecret_test"))

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse jwt"})
		return
	}

	log.Info(claims)
	var body models.WoltWebHook

	// unmarshal it to struct
	jsonString, err := json.Marshal(claims)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal jwt claims"})
		return
	}

	err = json.Unmarshal(jsonString, &body)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal jwt claims"})
		return
	}

	switch body.Type {
	case "order.pickup_eta_updated":
		//utils.SendSlackWebhookMessage(CreateWoltPickupEtaUpdatedMessage(body), viper.GetString("slack.sandbox"))
	case "order.dropoff_eta_updated":
		utils.SendSlackWebhookMessage(CreateWoltDropoffEtaUpdatedMessage(body), viper.GetString("slack.sandbox"))
	case "order.received":
		utils.SendSlackWebhookMessage(CreateWoltOrderReceivedMessage(body), viper.GetString("slack.sandbox"))
	case "order.picked_up":
		utils.SendSlackWebhookMessage(CreateWoltOrderPickedUpMessage(body), viper.GetString("slack.sandbox"))
	case "order.delivered":
		utils.SendSlackWebhookMessage(CreateWoltOrderDeliveredMessage(body), viper.GetString("slack.sandbox"))
	case "order.rejected":
		utils.SendSlackWebhookMessage(CreateWoltOrderRejectedMessage(body), viper.GetString("slack.sandbox"))
	case "order.pickup_started":
		utils.SendSlackWebhookMessage(CreateWoltOrderPickupStartedMessage(body), viper.GetString("slack.sandbox"))
	case "order.customer_no_show":
		utils.SendSlackWebhookMessage(CreateWoltOrderCustomerNoShowMessage(body), viper.GetString("slack.sandbox"))
	case "order.location_updated":
		utils.SendSlackWebhookMessage(CreateWoltOrderLocationUpdatedMessage(body), viper.GetString("slack.sandbox"))
	}

	c.JSON(200, "ok")
}

func CreateWoltPickupEtaUpdatedMessage(webhhook models.WoltWebHook, restaurant models.Restaurant, order models.RestaurantOrders, restaurantSettings models.RestaurantSettingsByName) (slackMessage models.SlackWebhookMessage) {
	realPickupEta, _ := utils.GetPickupTime(restaurantSettings, order)
	slackMessage.Text = "Wolt Order Pickup ETA Updated" + "\n" + webhhook.Details.Id + "\n" + webhhook.Details.Pickup.Eta.Format("2006-01-02 15:04:05") + "\n" + restaurant.Name +
		"\n" + order.OrderNumber + "\n" + "Real Pickup ETA: " + realPickupEta.Format("2006-01-02 15:04:05") + "\n" + "Pickup ETA Difference: " + strconv.FormatInt(webhhook.Details.Pickup.Eta.Unix()-realPickupEta.Unix(), 10) + " seconds"
	return
}

func CreateWoltDropoffEtaUpdatedMessage(order models.WoltWebHook) (slackMessage models.SlackWebhookMessage) {
	slackMessage.Text = "Wolt Order Dropoff ETA Updated" + "\n" + order.Details.Id + "\n" + order.Details.Dropoff.Eta.Format("2006-01-02 15:04:05")
	return
}

func CreateWoltOrderReceivedMessage(order models.WoltWebHook) (slackMessage models.SlackWebhookMessage) {
	slackMessage.Text = "Wolt Order Received" + "\n" + order.Details.Id
	return
}

func CreateWoltOrderPickedUpMessage(order models.WoltWebHook) (slackMessage models.SlackWebhookMessage) {
	slackMessage.Text = "Wolt Order Picked Up" + "\n" + order.Details.Id
	return
}

func CreateWoltOrderDeliveredMessage(order models.WoltWebHook) (slackMessage models.SlackWebhookMessage) {
	slackMessage.Text = "Wolt Order Delivered" + "\n" + order.Details.Id
	return
}

func CreateWoltOrderRejectedMessage(order models.WoltWebHook) (slackMessage models.SlackWebhookMessage) {
	slackMessage.Text = "Wolt Order Rejected" + "\n" + order.Details.Id
	return
}

func CreateWoltOrderPickupStartedMessage(order models.WoltWebHook) (slackMessage models.SlackWebhookMessage) {
	slackMessage.Text = "Wolt Order Pickup Started" + "\n" + order.Details.Id
	return
}

func CreateWoltOrderCustomerNoShowMessage(order models.WoltWebHook) (slackMessage models.SlackWebhookMessage) {
	slackMessage.Text = "Wolt Order Customer No Show" + "\n" + order.Details.Id
	return
}

func CreateWoltOrderLocationUpdatedMessage(order models.WoltWebHook) (slackMessage models.SlackWebhookMessage) {
	slackMessage.Text = "Wolt Order Location Updated" + "\n" + order.Details.Id
	return
}
