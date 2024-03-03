package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/karolpiernikarz/automanage/models"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/helpers"
)

func GetRestaurantOrdersFromCompanyId(c *gin.Context) {
	companyId := c.Param("companyId")
	if companyId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "companyid is required"})
		return
	}
	allOrders := helpers.GetOrdersFromCompanyID(companyId)
	c.JSON(http.StatusOK, allOrders)
}

func GetRestaurantTablesFromCompanyId(c *gin.Context) {
	tables := c.PostFormArray("tables")
	companyId := c.Param("companyId")
	if companyId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "companyid is required"})
		return
	}
	restaurants := helpers.GetRestaurantIdsFromCompanyId(companyId)
	var allTables []models.RestaurantTables
	for _, restaurant := range restaurants {
		table := helpers.GetRestaurantTables(restaurant, tables)
		allTables = append(allTables, table)
	}
	c.JSON(http.StatusOK, allTables)
}

func GetRestaurantTablesFromRestaurantId(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	tables := c.PostFormArray("tables")
	if restaurantId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "restaurantId is required"})
		return
	}
	table := helpers.GetRestaurantTables(restaurantId, tables)
	c.JSON(http.StatusOK, table)
}

func GetRestaurantTablesFromAllRestaurants(c *gin.Context) {
	tables := c.PostFormArray("tables")
	restaurantIds := c.Param("restaurantIds")
	restaurants := make([]string, 0)
	if restaurantIds == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "restaurantId is required"})
	} else {
		if restaurantIds == "all" {
			restaurants = helpers.GetAllRestaurantIds()
		} else {
			// split restaurantIds by -
			restaurants = strings.Split(restaurantIds, "-")
		}
	}
	var allTables []models.RestaurantTables
	for _, restaurant := range restaurants {
		table := helpers.GetRestaurantTables(restaurant, tables)
		allTables = append(allTables, table)
	}
	c.JSON(http.StatusOK, allTables)
}

func GetRestaurantCount(c *gin.Context) {
	restaurantCount := helpers.GetRestaurantCount()
	c.JSON(http.StatusOK, restaurantCount)
}

func GetRestaurantsOpen(c *gin.Context) {
	restaurants := helpers.GetAllRestaurants()
	currentTime := time.Now()
	var response []uint
	for _, restaurant := range restaurants {
		restaurantSettings := helpers.GetRestaurantSettingsByName(strconv.Itoa(int(restaurant.ID)))
		openTimes := helpers.GetTheOpenTimes(&restaurantSettings)
		if !openTimes[currentTime.Weekday()].IsOpen {
			continue
		}
		parsedOpenTime, _ := time.Parse("15:04", openTimes[currentTime.Weekday()].Open[0:5])
		parsedCloseTime, _ := time.Parse("15:04", openTimes[currentTime.Weekday()].Open[6:11])
		// if current hour is between open and close time
		if parsedCloseTime.Hour() < parsedOpenTime.Hour() {
			// if parsedCloseTime is less than parsedOpenTime, it means that the restaurant is open until the next day
			// set the parsedCloseTime to the 23:59:59
			parsedCloseTime, _ = time.Parse("15:04", "23:59")
		}
		if currentTime.Hour() >= parsedOpenTime.Hour() && currentTime.Hour() <= parsedCloseTime.Hour() {
			response = append(response, restaurant.ID)
		}
	}
	c.JSON(http.StatusOK, response)
}

func GetSupportResponse(c *gin.Context) {
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
	response, err := helpers.LiveSupportLastOrders(parsedStart, parsedEnd)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, gin.H{
		"message": response,
	})
}

func CourierList(c *gin.Context) {
	restaurants := helpers.GetAllRestaurants()
	var response []map[string]string
	for _, restaurant := range restaurants {
		if restaurant.Info.Data().CurrierType == "" || restaurant.Info.Data().CurrierType == "none" {
			continue
		}
		courier := make(map[string]string)
		courier["restaurantName"] = restaurant.Name
		courier["currier"] = restaurant.Info.Data().CurrierType
		response = append(response, courier)
	}
	// make a map to store the count of each courier
	courierCount := make(map[string]int)
	for _, courier := range response {
		courierCount[courier["currier"]]++
	}
	response = append(response, map[string]string{
		"Total":   strconv.Itoa(len(response)),
		"Wolt":    strconv.Itoa(courierCount["wolt"]),
		"Lastlap": strconv.Itoa(courierCount["lastLap"]),
	})
	c.JSON(200, response)
}

func GetAllRestaurants(c *gin.Context) {
	restaurants := helpers.GetAllRestaurants()
	c.JSON(200, restaurants)
}

func GetAllCompanies(c *gin.Context) {
	companies := helpers.GetCompanies()
	c.JSON(200, companies)
}
