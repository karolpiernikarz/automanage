package middlewares

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/cache"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/spf13/viper"
)

func RestaurantTokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateRestaurantToken(c)
	}
}

func validateRestaurantToken(c *gin.Context) {
	restaurantsByte, err := cache.GetValueFromKey("restaurants")
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(500)
		return
	}
	var restaurants []models.Restaurant
	err = json.Unmarshal([]byte(restaurantsByte), &restaurants)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(500)
		return
	}
	token := c.GetHeader("X-Api-Key")
	if token == "" {
		c.AbortWithStatus(401)
		return
	}
	if token == viper.GetString("test.restauranttoken") {
		c.Set("restaurantid", "test")
		c.Next()
		return
	}
	for _, r := range restaurants {
		if r.Token == token {
			c.Set("restaurantid", strconv.FormatUint(uint64(r.ID), 10))
			c.Next()
			return
		}
	}
	c.AbortWithStatus(401)
}
