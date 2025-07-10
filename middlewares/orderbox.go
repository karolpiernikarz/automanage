package middlewares

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/karolpiernikarz/automanage/cache/dbcache"
	"github.com/karolpiernikarz/automanage/services/redis"
	"github.com/karolpiernikarz/automanage/utils"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/cache"
)

func SetOrderBoxAlive() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		if c.Writer.Status() == 200 {
			terminalId := c.Query("a")
			go setOrderboxLive(terminalId)
			if c.Writer.Size() > 0 {
				restaurantId := c.GetString("restaurantId")
				go handleOrderPrinted(blw, &restaurantId)
			}
		}
	}
}

func CallbackHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic in handleCallback:", r)
				}
			}()
			handleCallback(c)
		}()
	}
}

func SetRestaurantIdFromTerminalId() gin.HandlerFunc {
	return func(c *gin.Context) {
		terminalId := c.Query("a")
		restaurants, _ := dbcache.Restaurants()
		for _, r := range restaurants {
			if r.Info.Data().TerminalId == terminalId {
				c.Set("restaurantId", strconv.FormatUint(uint64(r.ID), 10))
				c.Next()
				return
			}
		}
		c.Next()
	}
}

func handleOrderPrinted(blw *bodyLogWriter, restaurantId *string) {
	orderNumber := blw.ResponseWriter.Header().Get("X-Order-Number")
	if orderNumber == "" {
		order := utils.SeparateGoodcomPrint(blw.body.String())
		orderNumber = order.OrderNumber
	}
	setOrderAsPrinted(*restaurantId, orderNumber)
}

func handleCallback(c *gin.Context) {
	if c.Writer.Status() != 200 {
		return
	}
	orderNumber := c.Query("o")
	if orderNumber == "" {
		fmt.Println("Empty order number in callback")
		return
	}

	underscoreIndex := strings.Index(orderNumber, "_")
	if underscoreIndex >= 0 {
		orderNumber = orderNumber[underscoreIndex+1:]
	}

	ak := c.Query("ak")
	reason := c.Query("m")
	restaurantId := c.GetString("restaurantId")
	// remove '_minutes' suffix from reason safely and convert to int
	message := ""
	if ak == "Accepted" {
		if strings.HasSuffix(reason, "_minutes") {
			reason = strings.TrimSuffix(reason, "_minutes")
		} else {
			fmt.Printf("handleCallback warning: expected suffix '_minutes' not found in reason '%s'\n", reason)
		}

		if reason == "" {
			// nothing to add, treat as plain acceptance
			message = "Accepted"
		} else {
			minutesInt, err := strconv.Atoi(reason)
			if err != nil {
				fmt.Printf("handleCallback parse error: unable to convert reason '%s' to int: %v\n", reason, err)
				return
			}
			if minutesInt > 0 {
				message = "Accepted and added " + reason + " minutes"
			} else {
				message = "Accepted"
			}
		}
	}
	if ak == "Rejected" {
		message = "Rejected with reason: " + reason
	}
	setOrderAsResponded(restaurantId, orderNumber, []byte(message))
}

func setOrderAsPrinted(restaurantId, orderNumber string) {
	err := redis.Set("order_"+restaurantId+"_"+orderNumber+"-printed", "1", 76*time.Hour)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func setOrderAsResponded(restaurantId, orderNumber string, message []byte) {
	err := redis.Set("order_"+restaurantId+"_"+orderNumber+"-responded", message, 76*time.Hour)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func setOrderboxLive(terminalId string) {
	err := cache.SetKeyValue([]byte("terminal_"+terminalId), []byte("1"), 5*time.Minute)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
