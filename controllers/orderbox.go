package controllers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/karolpiernikarz/automanage/cache"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/helpers"
	log "github.com/sirupsen/logrus"
)

func OrderBoxList(c *gin.Context) {
	username := c.Query("u")
	password := c.Query("p")
	terminalId := c.Query("a")
	if terminalId == "11111" {
		err := helpers.GetTestResponseForGoodcom(c)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		return
	}

	i := 0
	for cache.IsKeyExist([]byte("orderbox_" + terminalId + "_" + strconv.Itoa(i))) {
		orderboxInfo, err := helpers.GetOrderBoxInfo(terminalId + "_" + strconv.Itoa(i))
		if err != nil {
			log.WithFields(log.Fields{"terminalid": terminalId, "url": c.Request.URL, "client_ip": c.ClientIP()}).Debug(err.Error())
			c.String(404, err.Error())
			return
		}

		if orderboxInfo.TerminalType == "NewAdmin" {
			orderboxnewadminresponse(c)
			return
		}

		if orderboxInfo.TerminalPassword != password {
			c.String(401, "wrong username or password")
			return
		}
		if orderboxInfo.TerminalUsername != username {
			c.String(401, "wrong username or password")
			return
		}
		restaurantUrl, err := helpers.GetRestaurantUrlFromID(strconv.FormatUint(uint64(orderboxInfo.ID), 10))
		if err != nil {
			log.WithFields(log.Fields{"terminalId": terminalId, "url": c.Request.URL, "client_ip": c.ClientIP()}).Debug(err.Error())
			c.String(500, "error getting restaurant url")
			return
		}
		resp, err := http.Get(restaurantUrl + c.Request.URL.String())
		if err != nil {
			log.WithFields(log.Fields{"terminalId": terminalId, "url": c.Request.URL, "client_ip": c.ClientIP()}).Debug(err.Error())
			c.String(500, "error getting response from restaurant")
			return
		}
		// Read the body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.String(500, "error reading body")
			return
		}
		// Close the body when the function returns
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(resp.Body)

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			c.String(resp.StatusCode, string(body))
			return
		}
		c.Header("content-type", resp.Header.Get("content-type"))
		c.Header("X-Order-Number", resp.Header.Get("X-Order-Number"))

		if len(string(body)) == 0 {
			i++
			continue
		}
		c.Set("restaurantId", strconv.FormatUint(uint64(orderboxInfo.ID), 10))
		err = cache.SetKeyValue([]byte("orderboxcallback_"+terminalId+"_"+resp.Header.Get("X-Order-Number")), []byte(strconv.Itoa(i)), 3*time.Hour*24)
		if err != nil {
			fmt.Println(err)
		}
		c.String(resp.StatusCode, string(body))
		return
	}

	if !cache.IsKeyExist([]byte("orderbox_" + terminalId + "_" + strconv.Itoa(0))) {
		c.String(404, "orderbox not found")
		return
	}

	c.String(200, "")
}

func OrderBoxCallBack(c *gin.Context) {
	username := c.Query("u")
	password := c.Query("p")
	orderNumber := c.Query("o")
	// remove all strings before first "_" in orderNumber
	orderNumber = orderNumber[strings.Index(orderNumber, "_")+1:]
	terminalId := c.Query("a")
	if terminalId == "11111" {
		err := helpers.GetTestResponseForGoodcom(c)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		return
	}
	i := "0"
	var err error
	if cache.IsKeyExist([]byte("orderboxcallback_" + terminalId + "_" + orderNumber)) {
		i, err = cache.GetValueFromKey("orderboxcallback_" + terminalId + "_" + orderNumber)
		if err != nil {
			i = "0"
		}
	}
	orderboxInfo, err := helpers.GetOrderBoxInfo(terminalId + "_" + i)
	if err != nil {
		log.WithFields(log.Fields{"terminalid": terminalId, "url": c.Request.URL, "client_ip": c.ClientIP()}).Debug(err.Error())
		c.String(404, terminalId+"_"+i)
		return
	}

	if orderboxInfo.TerminalType == "NewAdmin" {
		orderboxnewadminresponse(c)
		c.Set("restaurantId", strconv.FormatUint(uint64(orderboxInfo.ID), 10))
		return
	}

	if orderboxInfo.TerminalPassword != password {
		c.String(401, "wrong username or password")
		return
	}
	if orderboxInfo.TerminalUsername != username {
		c.String(401, "wrong username or password")
		return
	}
	restaurantUrl, err := helpers.GetRestaurantUrlFromID(strconv.FormatUint(uint64(orderboxInfo.ID), 10))
	if err != nil {
		log.WithFields(log.Fields{"terminalId": terminalId, "url": c.Request.URL, "client_ip": c.ClientIP()}).Debug(err.Error())
		c.String(500, "error getting restaurant url")
		return
	}
	resp, err := http.Get(restaurantUrl + c.Request.URL.String())
	if err != nil {
		log.WithFields(log.Fields{"terminalId": terminalId, "url": c.Request.URL, "client_ip": c.ClientIP()}).Debug(err.Error())
		c.String(500, "error getting response from restaurant")
		return
	}
	// Read the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.String(500, "error reading body")
		return
	}
	// Close the body when the function returns
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	c.Set("restaurantId", strconv.FormatUint(uint64(orderboxInfo.ID), 10))
	c.String(resp.StatusCode, string(body))
}

func OrderBoxHealthcheck(c *gin.Context) {
	// get all keys from cache with prefix "terminal_"
	keys, err := cache.GetKeysFromPrefix("terminal_")
	if err != nil {
		c.JSON(500, err)
	}
	// remove prefix from keys
	response := make([]string, 0)
	replacer := strings.NewReplacer("terminal_", "")
	// append keys to response
	for _, key := range keys {
		response = append(response, replacer.Replace(key))
	}
	// return response
	c.JSON(200, response)
}

func orderboxnewadminresponse(c *gin.Context) {
	resp, err := http.Get(viper.GetString("newadmin.url") + c.Request.URL.String())
	if err != nil {
		log.WithFields(log.Fields{"terminalId": c.Query("a"), "url": c.Request.URL, "client_ip": c.ClientIP()}).Debug(err.Error())
		c.String(500, "error getting response from restaurant")
		return
	}
	// Read the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.String(500, "error reading body from restaurant newadmin")
		return
	}
	// Close the body when the function returns
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		c.Header("content-type", resp.Header.Get("content-type"))
		c.String(resp.StatusCode, string(body))
		return
	}
	c.Header("content-type", resp.Header.Get("content-type"))
	c.Header("X-Order-Number", resp.Header.Get("X-Order-Number"))
	c.String(resp.StatusCode, string(body))
}
