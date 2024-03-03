package controllers

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/helpers"
	"github.com/karolpiernikarz/automanage/models"
)

func CreateRestaurantOrderReport(c *gin.Context) {
	data := c.Query("restaurants")
	if data == "" {
		c.JSON(400, "restaurants can't be empty")
		return
	}
	var orders []models.AllRestaurantsOrders
	if data == "all" {
		orders = helpers.GetOrdersFromAllRestaurants()
	} else {
		restaurants := strings.Split(data, "-")
		orders = helpers.GetOrdersFromRestaurants(restaurants)
	}
	filename, err := helpers.CreateXlsxFileFromAllOrders(orders)
	if err != nil {
		c.JSON(500, err)
		return
	}
	//c.JSON(http.StatusOK, filename)
	c.File(filename)
}

func CreateRestaurantOrderReportWithTime(c *gin.Context) {
	data := c.Query("restaurants")
	timeStart := c.Query("timestart")
	timeEnd := c.Query("timeend")
	dateFilter := c.Query("dateFilter")
	if data == "" {
		c.JSON(400, "restaurants can't be empty")
		return
	}
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
	var orders []models.AllRestaurantsOrders
	if dateFilter == "date" {
		if data == "all" {
			orders = helpers.GetOrdersFromAllRestaurantsWithTimeUsingDate(parsedStart, parsedEnd)
		} else {
			restaurants := strings.Split(data, "-")
			orders = helpers.GetOrdersFromRestaurantsWithTimeUsingDate(restaurants, parsedStart, parsedEnd)
		}
	} else {
		if data == "all" {
			orders = helpers.GetOrdersFromAllRestaurantsWithTime(parsedStart, parsedEnd)
		} else {
			restaurants := strings.Split(data, "-")
			orders = helpers.GetOrdersFromRestaurantsWithTime(restaurants, parsedStart, parsedEnd)
		}
	}
	filename, err := helpers.CreateXlsxFileFromAllOrders(orders)
	if err != nil {
		c.JSON(500, err)
		return
	}

	//c.JSON(200, filename)
	c.File(filename)
}

func CreateRestaurantReports(c *gin.Context) {
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
	restaruants := helpers.GetAllRestaurants()

	filename, err := helpers.CreateXlsxFileFromRestaruantsReportWithTime(restaruants, parsedStart, parsedEnd)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.File(filename)
}
