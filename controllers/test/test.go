package controllers_test

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/helpers"
)

func GetAllRestaurantOrdersWithTime(c *gin.Context) {
	timeStart := c.Query("timestart")
	timeEnd := c.Query("timeend")
	parsedStart, err := time.Parse("2006-01-02T15:04", timeStart)
	if err != nil {
		c.JSON(400, err)
		return
	}
	parsedEnd, err := time.Parse("2006-01-02T15:04", timeEnd)
	if err != nil {
		c.JSON(400, err)
		return
	}
	response, err := helpers.LiveSupportLastOrders_Test(parsedStart, parsedEnd)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, gin.H{
		"message": response,
	})
}

func OrderBoxHealthById(c *gin.Context) {
	var response []int
	i1 := rand.Intn(100)
	i2 := rand.Intn(3)
	if i1 == 15 {
		i1 = 52
	}
	if i2 == 0 {
		i2 = 15
	}
	response = append(response, i1)
	response = append(response, i2)
	c.JSON(200, response)
}

func RestaurantTables(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	tables := c.PostFormArray("tables")
	if restaurantId != "5" {
		c.JSON(404, "not found")
		return
	}
	response := helpers.GetRestaurantTables(restaurantId, tables)
	c.JSON(200, response)
}
