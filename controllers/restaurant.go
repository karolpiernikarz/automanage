package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/cache/dbcache"
	"github.com/karolpiernikarz/automanage/helpers"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/karolpiernikarz/automanage/utils"
)

func GetRestaurantOrder(c *gin.Context) {
	restaurantId := c.Query("id")
	orderId := c.Query("orderid")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	if orderId == "" {
		c.JSON(400, "orderid can't be empty")
		return
	}
	response := helpers.GetRestaurantOrderFromId(restaurantId, orderId)
	c.JSON(200, response)
}

func GetRestaurantProduct(c *gin.Context) {
	restaurantId := c.Query("id")
	productId := c.Query("productid")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	if productId == "" {
		c.JSON(400, "productid can't be empty")
		return
	}
	response := helpers.GetRestaurantProductFromId(productId, productId)
	c.JSON(200, response)
}

func GetRestaurantAddress(c *gin.Context) {
	restaurantId := c.Query("id")
	address := c.Query("adressid")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	if address == "" {
		c.JSON(400, "adressid can't be empty")
		return
	}
	response := helpers.GetRestaurantAddressFromId(restaurantId, address)
	c.JSON(200, response)
}

func GetRestaurantAttribute(c *gin.Context) {
	restaurantId := c.Query("id")
	attributeName := c.Query("name")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	if attributeName == "" {
		c.JSON(400, "name can't be empty")
		return
	}
	response := helpers.GetRestaurantAttributeFromName(restaurantId, attributeName)
	c.JSON(200, response)
}

func GetRestaurantCart(c *gin.Context) {
	restaurantId := c.Query("id")
	cartId := c.Query("cartid")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	if cartId == "" {
		c.JSON(400, "cartid can't be empty")
		return
	}
	response := helpers.GetRestaurantCartFromCartId(restaurantId, cartId)
	c.JSON(200, response)
}

func GetRestaurantCategory(c *gin.Context) {
	restaurantId := c.Query("id")
	categoryId := c.Query("categoryid")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	if categoryId == "" {
		c.JSON(400, "categoryid can't be empty")
		return
	}
	response := helpers.GetRestaurantCategoryFromId(restaurantId, categoryId)
	c.JSON(200, response)
}

func GetRestaurantCoupon(c *gin.Context) {
	restaurantId := c.Query("id")
	couponId := c.Query("couponid")
	couponName := c.Query("couponname")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	if (couponId == "") && (couponName == "") {
		c.JSON(400, "couponid or couponname must be set")
		return
	}
	response, exist := helpers.GetRestaurantCoupon(restaurantId, couponName, couponId)
	if !exist {
		c.JSON(404, "can't find the coupon")
		return
	}
	c.JSON(200, response)
}

func GetRestaurantCustomer(c *gin.Context) {
	restaurantId := c.Query("id")
	customerId := c.Query("customerid")
	customerEmail := c.Query("customeremail")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	if (customerId == "") && (customerEmail == "") {
		c.JSON(400, "customerid or customeremail must be set")
		return
	}
	response, exist := helpers.GetRestaurantCustomer(restaurantId, customerId, customerEmail)
	if !exist {
		c.JSON(404, "can't find the coupon")
		return
	}
	c.JSON(200, response)
}

func GetRestaurantOrders(c *gin.Context) {
	restaurantId := c.Query("id")
	timeStart := c.Query("timestart")
	timeEnd := c.Query("timeend")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
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
	response := helpers.GetOrdersFromRestaurantWithTime(restaurantId, parsedStart, parsedEnd)
	c.JSON(200, response)
}

func GetRestaurantCustomers(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantCustomers(restaurantId)
	for i := range response {
		response[i].Password = ""
	}
	c.JSON(200, response)
}

func GetRestaurantAddresses(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantAddresses(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantAttributes(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantAttributes(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantCarts(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantCarts(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantCategories(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantCategories(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantCoupons(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantCoupons(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantFailedJobs(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantFailedJobs(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantMigrations(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantMigrations(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantOrderActions(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantOrderActions(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantOrderProducts(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantOrderProducts(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantPasswordResets(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantPasswordResets(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantProducts(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantProducts(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantProductExtraGroups(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantProductExtraGroups(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantProductVariants(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantProductVariants(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantProductVariantOptions(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantProductVariantOptions(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantSettings(c *gin.Context) {
	restaurantId := c.Query("id")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantSettings(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantIsOpen(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
	}
	response := helpers.IsRestaurantOpen(restaurantId)
	c.JSON(200, response)
}

func GetRestaurantLastNOrders(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	n := c.Query("n")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	// n must be a number
	var nInt int
	if n == "" {
		nInt = 10
	} else {
		var err error
		nInt, err = strconv.Atoi(n)
		if err != nil {
			c.JSON(400, "n must be a number")
			return
		}
	}
	response := helpers.GetLastNOrdersWithStatusNone0(restaurantId, nInt)
	c.JSON(200, response)
}

func GetRestaurantSettingsByName(c *gin.Context) {
	restaurantId := c.Param("restaurantId")
	if restaurantId == "" {
		c.JSON(400, "id can't be empty")
		return
	}
	response := helpers.GetRestaurantSettingsByName(restaurantId)
	c.JSON(200, response)
}

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
	response := helpers.GetAllRestaurantsOrderWithTime(parsedStart, parsedEnd)
	c.JSON(200, response)
}

func GetOrdersFromEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(400, "email can't be empty")
		return
	}
	orders := helpers.GetAllOrdersByEmail(email)

	if len(orders) == 0 {
		c.JSON(404, "can't find any orders")
		return
	}

	var response []struct {
		RestaurantName string `json:"restaurant_name"`
		RestaurantId   int    `json:"restaurant_id"`
		OrderId        int    `json:"order_id"`
		OrderNumber    string `json:"order_number"`
		OrderStatus    int    `json:"order_status"`
		OrderCreatedAt string `json:"order_created_at"`
		OrderLink      string `json:"order_link"`
	}

	restaurants, err := dbcache.Restaurants()
	if err != nil {
		c.JSON(500, err)
		return
	}

	for _, order := range orders {
		if order.PaymentId == "" {
			continue
		}
		restaurantId, err := utils.ExtractNumericPartFromString(order.PaymentId)
		if err != nil {
			continue
		}
		restaurant := models.Restaurant{}
		for _, r := range restaurants {
			if r.ID == uint(restaurantId) {
				restaurant = r
				break
			}
		}

		if restaurant.ID == 0 {
			continue
		}

		response = append(response, struct {
			RestaurantName string `json:"restaurant_name"`
			RestaurantId   int    `json:"restaurant_id"`
			OrderId        int    `json:"order_id"`
			OrderNumber    string `json:"order_number"`
			OrderStatus    int    `json:"order_status"`
			OrderCreatedAt string `json:"order_created_at"`
			OrderLink      string `json:"order_link"`
		}{
			RestaurantName: restaurant.Name,
			RestaurantId:   int(restaurant.ID),
			OrderId:        order.Id,
			OrderNumber:    order.OrderNumber,
			OrderStatus:    order.Status,
			OrderCreatedAt: order.CreatedAt,
			OrderLink:      restaurant.Website + "/order/" + order.PaymentId + "/status",
		})

	}

	if len(response) == 0 {
		c.JSON(404, "can't find any orders")
		return
	}

	c.JSON(200, response)
}
