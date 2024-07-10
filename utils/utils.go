package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v5"
	"github.com/karolpiernikarz/automanage/cache"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/spf13/viper"
)

func ParseToken(tokenString string, tokenSecret string) (claims jwt.Claims, err error) {
	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	// check if token is valid
	if !token.Valid {
		return nil, fmt.Errorf("token not valid")
	}
	// get claims
	claims = token.Claims
	return claims, nil
}

func SendNtfyNotification(message models.NtfyMessage) (err error) {
	// convert message to json
	jsonData, err := json.Marshal(message)
	// check if error
	if err != nil {
		return err
	}
	// convert json to string
	jsonString := string(jsonData)
	// send request
	req, err := http.NewRequest("POST", viper.GetString("ntfy.server"), strings.NewReader(jsonString))
	if err != nil {
		return err
	}
	// set headers
	req.Header.Set("Authorization", "Bearer "+viper.GetString("ntfy.token"))
	// send request
	_, err = http.DefaultClient.Do(req)
	return
}

func Unmarshal(data []byte, v interface{}) (err error) {
	err = json.Unmarshal(data, v)
	return
}

func Marshal(v interface{}) (data []byte, err error) {
	data, err = json.Marshal(v)
	return
}

func IsOrderBoxActive(orderboxId string) (bool, error) {
	// get all keys from cache with terminal_ prefix
	keys, err := cache.GetKeysFromPrefix("terminal_")
	if err != nil {
		return false, err
	}
	// check if orderboxId is in keys
	for i := range keys {
		if keys[i] == "terminal_"+orderboxId {
			return true, nil
		}
	}
	// if not found
	return false, nil
}

// SendSlackWebhookMessage sends message to slack webhook
// message: message to send
// webhook: slack webhook url, after the https://hooks.slack.com/services/ part of the url
// webhook example: 1234567890/1234567890/1234567890
func SendSlackWebhookMessage(message models.SlackWebhookMessage, webhook string) (err error) {
	// convert message to json
	jsonData, err := json.Marshal(message)
	// check if error
	if err != nil {
		return err
	}
	// convert json to string
	jsonString := string(jsonData)
	// send request
	req, err := http.NewRequest("POST", "https://hooks.slack.com/services/"+webhook, strings.NewReader(jsonString))
	if err != nil {
		return err
	}
	// set headers
	// req.Header.Set("Content-Type", "application/json")
	// send request
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetOrderStatusString(orderStatus int) (status string) {
	switch orderStatus {
	case 0:
		return "status"
	case 1:
		return "Waiting for approval"
	case 2:
		return "Preparing"
	case 3:
		return "On the road"
	case 4:
		return "Delivered"
	case 5:
		return "Canceled"
	}
	return "unknown"
}

func GetOrderTypeString(orderType int) (oType string) {
	switch orderType {
	case 0:
		return "pickup"
	case 1:
		return "delivery"
	case 2:
		return "table"
	}
	return "unknown"
}

func GetDeliveryTypeString(deliveryType int) (dType string) {
	switch deliveryType {
	case 0:
		return "asap"
	case 1:
		return "preorder"
	}
	return "unknown"
}

func StringInSlice(str string, list []string) bool {
	for i := range list {
		if list[i] == str {
			return true
		}
	}
	return false
}

func UintInSlice(u uint, list []uint) bool {
	for i := range list {
		if list[i] == u {
			return true
		}
	}
	return false
}

func IntInSlice(i int, list []int) bool {
	for j := range list {
		if list[j] == i {
			return true
		}
	}
	return false
}

func StringToUint(str string) (u uint) {
	for _, s := range str {
		u = u*10 + uint(s-'0')
	}
	return
}

func GetStringInBetweenTwoString(str string, startS string, endS string) (result string) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result
	}
	result = newS[:e]
	return result
}

// SeparateGoodcomPrint separates order info from goodcom print
// not working for every order, needs to be fixed
func SeparateGoodcomPrint(order string) (orderBoxPrinter models.OrderBoxPrinter) {
	firstValue := GetStringInBetweenTwoString(order, "#", "*")
	// remove general info from string
	order = strings.Replace(order, "#"+firstValue, "", 1)
	orderType := GetStringInBetweenTwoString(order, "*", ";")
	// remove order type from string
	order = strings.Replace(order, "*"+orderType, "", 1)
	deliveryType := GetStringInBetweenTwoString(order, ";", ";")
	// remove delivery type from string
	order = strings.Replace(order, ";"+deliveryType, "", 1)
	deliveryTime := GetStringInBetweenTwoString(order, ";", ";")
	// remove delivery time from string
	order = strings.Replace(order, ";"+deliveryTime, "", 1)
	customerName := GetStringInBetweenTwoString(order, ";", ";")
	// remove customer name from string
	order = strings.Replace(order, ";"+customerName, "", 1)
	paymentType := GetStringInBetweenTwoString(order, ";", ";")
	// remove payment type from string
	order = strings.Replace(order, ";"+paymentType, "", 1)
	order = strings.Replace(order, ";", "", 1)
	orderNumber := GetStringInBetweenTwoString(order, "*", "*")
	// remove order id from string
	order = strings.Replace(order, "*"+orderNumber, "", 1)
	orderInfo := GetStringInBetweenTwoString(order, "*", "*")
	// remove order from string
	order = strings.Replace(order, "*"+orderInfo, "", 1)
	subTotal := GetStringInBetweenTwoString(order, "*", ";")
	// remove sub-total from string
	order = strings.Replace(order, "*"+subTotal, "", 1)
	order = strings.Replace(order, ";", "", 1)
	bagFee := GetStringInBetweenTwoString(order, ";", ";")
	// remove bag fee from string
	order = strings.Replace(order, ";"+bagFee, "", 1)
	deliveryFee := GetStringInBetweenTwoString(order, ";", ";")
	// remove delivery fee from string
	order = strings.Replace(order, ";"+deliveryFee, "", 1)
	serviceFee := GetStringInBetweenTwoString(order, ";", ";")
	// remove service fee from string
	order = strings.Replace(order, ";"+serviceFee, "", 1)
	total := GetStringInBetweenTwoString(order, ";", ";")
	// remove total from string
	order = strings.Replace(order, ";"+total, "", 1)
	order = strings.Replace(order, ";*;", "", 1)
	notes := GetStringInBetweenTwoString(order, ";", ";")
	// remove notes from string
	order = strings.Replace(order, ";"+notes, "", 1)
	kundeInfo := GetStringInBetweenTwoString(order, ";", ";")
	// remove kunde info from string
	order = strings.Replace(order, ";"+kundeInfo, "", 1)
	customerName = GetStringInBetweenTwoString(order, ";", ";")
	// remove customer name from string
	order = strings.Replace(order, ";"+customerName, "", 1)
	customerDetails := GetStringInBetweenTwoString(order, ";", ";")
	// remove customer details from string
	order = strings.Replace(order, ";"+customerDetails, "", 1)
	customerPhone := GetStringInBetweenTwoString(order, ";", ";")
	// remove customer phone from string
	order = strings.Replace(order, ";"+customerPhone, "", 1)
	order = strings.Replace(order, ";", "", 1)
	customerCard := GetStringInBetweenTwoString(order, ";", ";")
	// remove customer card from string
	order = strings.Replace(order, ";"+customerCard, "", 1)
	orderBoxPrinter.FirstValue = firstValue
	orderBoxPrinter.OrderType = orderType
	orderBoxPrinter.DeliveryType = deliveryType
	orderBoxPrinter.DeliveryTime = deliveryTime
	orderBoxPrinter.CustomerName = customerName
	orderBoxPrinter.PaymentType = paymentType
	orderBoxPrinter.OrderNumber = orderNumber
	orderBoxPrinter.Order = order
	orderBoxPrinter.SubTotal = subTotal
	orderBoxPrinter.BagFee = bagFee
	orderBoxPrinter.DeliveryFee = deliveryFee
	orderBoxPrinter.ServiceFee = serviceFee
	orderBoxPrinter.Total = total
	orderBoxPrinter.Notes = notes
	orderBoxPrinter.KundeInfo = kundeInfo
	orderBoxPrinter.CustomerDetails = customerDetails
	orderBoxPrinter.CustomerPhone = customerPhone
	orderBoxPrinter.CustomerCard = customerCard
	return
}

func ExtractNumericPartFromString(str string) (int, error) {
	var numericPart string

	for _, char := range str {
		// Check if the character is a digit
		if unicode.IsDigit(char) {
			numericPart += string(char)
		} else {
			// Break the loop if a non-digit character is encountered
			break
		}
	}

	// Parse the numeric part into an integer
	num, err := strconv.Atoi(numericPart)
	return num, err
}

func IsPathWritable(path string) (Writable bool, err error) {
	Writable = false
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	err = nil
	if !info.IsDir() {
		return false, err
	}

	// Check if the user bit is enabled in file permission
	if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		return false, err
	}

	var stat syscall.Stat_t
	if err = syscall.Stat(path, &stat); err != nil {
		return false, err
	}

	err = nil
	if uint32(os.Geteuid()) != stat.Uid {
		Writable = false
		return false, err
	}

	Writable = true
	return
}

func GetPickupTime(settings models.RestaurantSettingsByName, order models.RestaurantOrders) (eta time.Time, err error) {
	deliveryTime, err := time.Parse("2006-01-02 15:04:05", order.Date)
	if err != nil {
		return eta, err
	}

	minPrepTime, err := strconv.Atoi(settings.OrderPickupTime)
	if err != nil {
		return eta, err
	}

	return deliveryTime.Add(time.Minute * time.Duration(-minPrepTime)), nil
}

// NotifySlackError constructs and sends a Slack notification message
func NotifySlackError(err error, to string, subject string, from string) {
	message := models.SlackWebhookMessage{
		Text: fmt.Sprintf("Failed to send email: \nError: %s\nTo: %s\nSubject: %s", err.Error(), to, subject+"\nFrom: "+from),
	}
	if slackErr := SendSlackWebhookMessage(message, viper.GetString("slack.webhook")); slackErr != nil {
		fmt.Println("Failed to send Slack notification:", slackErr)
	}
}
