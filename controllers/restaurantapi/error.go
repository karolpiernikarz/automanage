package restaurantapi

import (
	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/email"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Error(c *gin.Context) {
	messageType := c.PostForm("type")
	switch messageType {
	case "currier":
		currierError(c)
		return
	}
	c.JSON(404, gin.H{
		"message": "type not found",
	})
}

func currierError(c *gin.Context) {
	data := c.PostForm("data")
	errorResponse := c.PostForm("response")
	restaurantId := c.GetString("restaurantid")
	err := email.Send(viper.GetString("app.notifyemail"), restaurantId+"\n"+data+"\n"+errorResponse, "Currier Error")
	if err != nil {
		log.WithFields(log.Fields{"type": "currier", "restaurantId": restaurantId}).Warn("Error while sending email")
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "ok",
	})
	return
}
