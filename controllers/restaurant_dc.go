package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/helpers"
)

func RestartRestaurantApp(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	response, err := helpers.RestartRestaurantApp(restaurantId)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.String(200, response)
}

func GetRestaurantVariables(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	response, err := helpers.GetRestaurantVariables(restaurantId)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, response)
}

func GetRestaurantTerminalLogs(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	response, err := helpers.GetRestaurantTerminalLogs(restaurantId)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.String(200, response)
}

func GetRestaurantStorageLogs(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	response, err := helpers.GetRestaurantStorageLogs(restaurantId)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.String(200, response)
}
