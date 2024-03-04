package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/services/redis"
	"github.com/karolpiernikarz/automanage/utils"
	"gorm.io/datatypes"

	"github.com/karolpiernikarz/automanage/cache"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func GetAllRestaurants() (restaurants []models.Restaurant) {
	gormDB.Table(viper.GetString("db.admindb") + "." + "restaurants").Scan(&restaurants)
	return
}

func GetRestaurants(restaurantIds []string) (restaurants []models.Restaurant) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"restaurants").Where("(id) IN ?", restaurantIds).Scan(&restaurants)
	return
}

func GetRestaurant(restaurantId string) (restaurant models.Restaurant) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"restaurants").Where("id = ?", restaurantId).First(&restaurant)
	return
}

func GetCompanies() (companies []models.Companies) {
	gormDB.Table(viper.GetString("db.admindb") + "." + "companies").Scan(&companies)
	return
}

func GetCompanyNames() (companyNames []string) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"companies").Pluck("name", &companyNames)
	return
}

func GetRestaurantCount() (count int64) {
	gormDB.Table(viper.GetString("db.admindb") + "." + "restaurants").Count(&count)
	return
}

func GetRestaurantURLs() (restaurantURLs []string) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"restaurants").Pluck("website", &restaurantURLs)
	return
}

func GetCompany(companyId string) (company models.Companies) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"companies").Where("id = ?", companyId).First(&company)
	return
}
func GetRestaurantsFromCompanyId(companyId string) (restaurants []models.Restaurant) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"restaurants").Where("company_id = ?", companyId).Scan(&restaurants)
	return
}

func GetRestaurantIdsFromCompanyId(companyId string) (restaurantIds []string) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"restaurants").Where("company_id = ?", companyId).Pluck("id", &restaurantIds)
	return
}

func GetAllRestaurantIds() (restaurantIds []string) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"restaurants").Pluck("id", &restaurantIds)
	return
}

func GetOrdersAndProductFromRestaurants(restaurantIds []string) (response []models.RestaurantOrdersAndProducts) {
	// Get the restaurants from the database
	restaurants := GetRestaurants(restaurantIds)
	var ordersAndProducts models.RestaurantOrdersAndProducts
	// Loop through the restaurants
	for _, restaurant := range restaurants {
		// Get the orders from the restaurant
		ordersAndProducts.Orders = GetOrdersFromRestaurant(strconv.FormatUint(uint64(restaurant.ID), 10))
		// Get the products from the restaurant
		ordersAndProducts.Products = GetRestaurantProducts(strconv.FormatUint(uint64(restaurant.ID), 10))
		// Set the restaurant id
		ordersAndProducts.Id = restaurant.ID
		// append the response
		response = append(response, ordersAndProducts)
	}
	// Return the response
	return
}

func GetOrdersFromCompanyID(companyId string) (allOrdersArray []models.AllRestaurantsOrders) {
	restaurants := GetRestaurantsFromCompanyId(companyId)
	var allOrders models.AllRestaurantsOrders
	for _, restaurant := range restaurants {
		orders := GetOrdersFromRestaurant(strconv.FormatUint(uint64(restaurant.ID), 10))
		allOrders.Restaurant = restaurant
		allOrders.Orders = orders
		allOrdersArray = append(allOrdersArray, allOrders)
	}
	return
}

func GetRestaurantFromToken(restaurantToken string) (restaurant models.Restaurant) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"restaurants").Where("token = ?", restaurantToken).First(&restaurant)
	return
}

// GetRestaurantFromId gets the restaurant from the database
// Returns the restaurant
func GetRestaurantFromId(restaurantId string) (restaurant models.Restaurant) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"restaurants").Where("id = ?", restaurantId).First(&restaurant)
	return
}

func GetTokenFromRestaurantId(restaurantId string) (restauranttoken string, err error) {
	var restaurant models.Restaurant
	// Get the restaurant from the database
	gormDB.Table(viper.GetString("db.admindb")+"."+"restaurants").Where("id = ?", restaurantId).First(&restaurant)
	// Return the token
	return restaurant.Token, err
}

// GetTerminalInfoFromRestaurant
// Get the Terminal info from the restaurant
// Returns the terminal id, terminal username and terminal password
// Returns an error if there is one
// Returns the terminal id as a string
// Returns the terminal username as a string
// Returns the terminal password as a string
func GetTerminalInfoFromRestaurant(restaurantId string) (terminalId string, terminalUsername string, terminalPassword string, err error) {
	var restaurant models.Restaurant
	// Get the restaurant from the database
	gormDB.Table(viper.GetString("db.admindb")+"."+"restaurants").Where("id = ?", restaurantId).First(&restaurant).Table("restaurant")
	// Return the terminal id, terminal username and terminal password
	return fmt.Sprint(restaurant.Info.Data().TerminalId), restaurant.Info.Data().TerminalUsername, restaurant.Info.Data().TerminalPassword, err
}

func GetResponseForGoodcom(restaurantId string, c *gin.Context) (err error) {
	// Get the restaurant url
	restaurantUrl, err := GetRestaurantUrlFromID(restaurantId)
	if err != nil {
		c.String(500, "error getting restaurant url")
		return err
	}
	// Get the response from the restaurant
	resp, err := http.Get(restaurantUrl + c.Request.URL.String())
	if err != nil {
		c.String(500, "error getting response from restaurant")
		return err
	}
	// Read the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.String(500, "error reading body")
		return err
	}
	// Close the body when the function returns
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	c.Header("content-type", resp.Header.Get("content-type"))
	c.Header("X-Order-Number", resp.Header.Get("X-Order-Number"))
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		c.String(resp.StatusCode, string(body))
		return
	}
	// Send the response back to the client
	c.String(resp.StatusCode, string(body))
	return
}

func GetTestResponseForGoodcom(c *gin.Context) (err error) {
	// set the test restaurant url
	restaurantUrl := "http://test.online-takeaway.dk"
	// Get the response from the test restaurant
	resp, err := http.Get(restaurantUrl + c.Request.URL.String())
	if err != nil {
		c.String(500, "error getting response from restaurant")
		return err
	}
	// Read the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.String(500, "error reading body")
		return err
	}
	// Close the body when the function returns
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	c.Header("content-type", resp.Header.Get("content-type"))
	c.Header("X-Order-Number", resp.Header.Get("X-Order-Number"))
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		c.String(resp.StatusCode, string(body))
		return
	}
	// Send the response back to the client
	c.String(resp.StatusCode, string(body))
	return
}

func GetOrderBoxInfo(orderboxId string) (orderboxData *models.OrderboxInfo, err error) {
	// Get the orderbox info from the cache
	var orderboxInfo models.OrderboxInfo
	value, err := cache.GetValueFromKey("orderbox_" + orderboxId)
	if err != nil {
		return nil, err
	}
	// Unmarshal the json
	err = json.Unmarshal([]byte(value), &orderboxInfo)
	if err != nil {
		return nil, err
	}
	// Return the orderbox info
	return &orderboxInfo, err
}

func GetRestaurantUrlFromID(restaurantId string) (restaurantUrl string, err error) {
	var dcFile models.DockerCompose
	// Read the file
	file, err := os.ReadFile(viper.GetString("app.workdir") + "/" + restaurantId + "/docker-compose.yaml")
	if err != nil {
		return "", err
	}
	// Unmarshal the yaml file
	err = yaml.Unmarshal(file, &dcFile)
	// Return the url
	return dcFile.Services.App.Environment.APP_URL, err
}

func GetRestaurantTerminalLogs(restaurantId string) (response string, err error) {
	// set the command and args
	command := "docker compose -f " + viper.GetString("app.workdir") + "/" + restaurantId + "/docker-compose.yaml logs"
	cmd := exec.Command("sh", "-c", command)
	// run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	// send the output back to the client
	return string(output), err
}

func GetRestaurantStorageLogs(restaurantId string) (response string, err error) {
	// set the command and args
	command := "docker compose -f " + viper.GetString("app.workdir") + "/" + restaurantId + "/docker-compose.yaml exec app cat storage/logs/laravel.log"
	cmd := exec.Command("sh", "-c", command)
	// run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	// send the output back to the client
	return string(output), err
}

func RestartRestaurantApp(restaurantId string) (response string, err error) {
	// set the command and args
	command := "docker compose -f " + viper.GetString("app.workdir") + "/" + restaurantId + "/docker-compose.yaml restart"
	cmd := exec.Command("sh", "-c", command)
	// run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	// send the output back to the client
	return string(output), err
}

func GetRestaurantVariables(restaurantId string) (response interface{}, err error) {
	// Create a struct to store the YAML data in
	var dcFile models.DockerCompose
	// Read YAML file
	file, err := os.ReadFile(viper.GetString("app.workdir") + "/" + restaurantId + "/docker-compose.yaml")
	if err != nil {
		return "", err
	}
	// Unmarshal YAML file
	err = yaml.Unmarshal(file, &dcFile)
	// Remove sensitive data
	dcFile.Services.App.Environment.DB_PASSWORD = ""
	dcFile.Services.App.Environment.MAIL_PASSWORD = ""
	dcFile.Services.App.Environment.AWS_SECRET_ACCESS_KEY = ""
	dcFile.Services.App.Environment.APP_KEY = ""
	return dcFile.Services.App.Environment, err
}

func GetRestaurantIdFromPaymentId(paymentId string) (restaurantId uint, err error) {
	// Create a regular expression to match the leftmost integer
	re := regexp.MustCompile(`^(\d+)`)

	// Find the first match in the string
	match := re.FindStringSubmatch(paymentId)

	if len(match) > 1 {
		// Convert the matched string to an uint32
		num, err := strconv.ParseUint(match[1], 10, 32)
		if err != nil {
			fmt.Println("Error converting string to uint32:", err)
			return 0, err
		}

		restaurantId = uint(num)
	} else {
		fmt.Println("No integer found in the string")
	}
	return
}

func LiveSupportLastOrders_Test(timeStart time.Time, timeEnd time.Time) (response []models.SupportLastOrdersResponse, err error) {
	orders := GetOrdersFromRestaurantWithTime_Test(timeStart, timeEnd)
	var r models.SupportLastOrdersResponse
	orderboxStatus := rand.Int() % 2
	for _, order := range orders {
		r.RestaurantId = 2
		r.OrderNumber = order.OrderNumber
		r.RestaurantName = "Test Restaurant"
		r.RestaurantPhone = "+4541414141"
		r.RestaurantAddress = "Test Address"
		r.CustomerName = order.Customer.Data().Name
		r.CustomerPhone = order.Customer.Data().Phone
		r.CustomerAddress = order.Address.Data().Detail
		r.OrderStatus = order.Status
		r.OrderType = order.Type
		r.OrderCreated = order.CreatedAt
		r.OrderDate = order.Date
		r.OrderIsPreOrder = order.IsPreOrder
		r.OrderAddress = order.Address.Data().Name
		r.OrderTotal = string(order.Prices.Data().Total)
		r.OrderLink = "https://test.online-takeaway.dk" + "/order/" + order.PaymentId + "/status"
		r.OrderboxStatus = orderboxStatus
		r.OrderboxPrinted = rand.Int() % 2
		responded := rand.Int() % 3
		if responded == 1 {
			r.OrderboxResponded = "Accepted"
		} else if responded == 0 {
			r.OrderboxResponded = "Rejected"
		} else {
			r.OrderboxResponded = ""
		}
		if responded == 1 {
			r.CurrierStatus = "Success"
		} else if responded == 0 {
			r.CurrierStatus = "Error"
		} else {
			r.CurrierStatus = ""
		}
		if r.CurrierStatus == "Success" {
			r.CurrierReferenceId = "123456789"
			r.CurrierTrackingUrl = "https://test.online-takeaway.dk"
		} else {
			r.CurrierReferenceId = ""
			r.CurrierTrackingUrl = ""
		}
		response = append(response, r)
	}
	return
}

func LiveSupportLastOrders(timeStart time.Time, timeEnd time.Time) (response []models.SupportLastOrdersResponse, err error) {

	restaurants := GetAllRestaurants()
	orders := GetAllRestaurantsOrderWithTime(timeStart, timeEnd)
	oBoxHealth, err := cache.GetKeysFromPrefix("terminal_")
	if err != nil {
		fmt.Println("error getting terminal keys: ", err)
		return nil, err
	}
	// strip the prefix from the keys
	replacer := strings.NewReplacer("terminal_", "")
	for i, key := range oBoxHealth {
		oBoxHealth[i] = replacer.Replace(key)
	}

	var r models.SupportLastOrdersResponse
	for _, order := range orders {
		if order.Orders == nil {
			continue
		}
		for _, restaurant := range restaurants {

			if int(restaurant.ID) == order.Id {
				for _, orderInfo := range order.Orders {
					r = models.SupportLastOrdersResponse{}
					r.RestaurantId = restaurant.ID
					r.OrderNumber = orderInfo.OrderNumber
					r.RestaurantName = restaurant.Name
					r.RestaurantPhone = restaurant.Phone
					r.RestaurantAddress = restaurant.Address
					r.CustomerName = orderInfo.Customer.Data().Name
					r.CustomerPhone = orderInfo.Customer.Data().Phone
					r.CustomerAddress = orderInfo.Address.Data().Detail
					r.OrderStatus = orderInfo.Status
					r.OrderType = orderInfo.Type
					r.OrderCreated = orderInfo.CreatedAt
					r.OrderDate = orderInfo.Date
					r.OrderAddress = orderInfo.Address.Data().Detail
					r.OrderTotal = string(orderInfo.Prices.Data().Total)
					r.OrderIsPreOrder = orderInfo.IsPreOrder
					r.OrderLink = restaurant.Website + "/order/" + orderInfo.PaymentId + "/status"
					if utils.StringInSlice(fmt.Sprintf("%v", restaurant.Info.Data().TerminalId), oBoxHealth) {
						r.OrderboxStatus = 1
					} else {
						r.OrderboxStatus = 0
					}
					if redis.IsExist("order_" + strconv.Itoa(int(r.RestaurantId)) + "_" + orderInfo.OrderNumber + "-printed") {
						r.OrderboxPrinted = 1
					} else {
						r.OrderboxPrinted = 0
					}
					// if redis.IsExist("order_" + strconv.Itoa(int(r.RestaurantId)) + "_" + orderInfo.OrderNumber + "-responded") {
					// 	r.OrderboxResponded, _ = redis.Get("order_" + strconv.Itoa(int(r.RestaurantId)) + "_" + orderInfo.OrderNumber + "-responded")
					// } else {
					// 	r.OrderboxResponded = ""
					// }
					if orderInfo.Other.Data() != nil {
						if orderInfo.Other.Data().Wolt.Data() != nil {
							if orderInfo.Other.Data().Wolt.Data().Data.Data() != nil {
								if orderInfo.Other.Data().Wolt.Data().Data.Data().Status != "" {
									r.CurrierStatus = orderInfo.Other.Data().Wolt.Data().Data.Data().Status
								}
								if orderInfo.Other.Data().Wolt.Data().Data.Data().Data.Data() != nil {
									if orderInfo.Other.Data().Wolt.Data().Data.Data().Data.Data().ReferenceId != "" {
										r.CurrierReferenceId = orderInfo.Other.Data().Wolt.Data().Data.Data().Data.Data().ReferenceId
									}
									if orderInfo.Other.Data().Wolt.Data().Data.Data().Data.Data().TrackingUrl != "" {
										r.CurrierTrackingUrl = orderInfo.Other.Data().Wolt.Data().Data.Data().Data.Data().TrackingUrl
									}
								}
							}
						}
					}

					// if strings.ToLower(restaurant.Info.Data().CurrierType) == "wolt" && orderInfo.Type == 1 {
					// 	currierStatusText := ""
					// 	currierStatusText, _ = redis.Get("order_" + strconv.FormatUint(uint64(restaurant.ID), 10) + "_" + orderInfo.OrderNumber + "_status")
					// 	r.CurrierStatusText = currierStatusText

					// 	CurrierPickupEta := ""
					// 	CurrierPickupEta, _ = redis.Get("order_" + strconv.FormatUint(uint64(restaurant.ID), 10) + "_" + orderInfo.OrderNumber + "_pickupeta")
					// 	r.CurrierPickupEta = CurrierPickupEta

					// 	// we don't track the dropoff eta for now
					// 	//r.CurrierDropoffEta, _ = cache.GetValueFromKey("order_" + strconv.FormatUint(uint64(restaurant.ID), 10) + "_" + orderInfo.OrderNumber + "_dropoff_eta")
					// }
					response = append(response, r)
				}
			}
		}
	}
	return
}

func GetRestaurantTables(restaurantId string, tables []string) (response models.RestaurantTables) {
	for _, table := range tables {
		switch table {
		case "addresses":
			response.Addresses = GetRestaurantAddresses(restaurantId)
		case "admins":
			response.Admins = GetRestaurantAdmins(restaurantId)
		case "attributes":
			response.Attributes = GetRestaurantAttributes(restaurantId)
		case "campaigns":
			response.Campaigns = GetRestaurantCampaigns(restaurantId)
		case "carts":
			response.Carts = GetRestaurantCarts(restaurantId)
		case "categories":
			response.Categories = GetRestaurantCategories(restaurantId)
		case "coupons":
			response.Coupons = GetRestaurantCoupons(restaurantId)
		case "customers":
			response.Customers = GetRestaurantCustomers(restaurantId)
		case "customers_points":
			response.CustomerPoints = GetRestaurantCustomerPoints(restaurantId)
		case "extras":
			response.Extras = GetRestaurantExtras(restaurantId)
		case "extras_groups":
			response.ExtraGroups = GetRestaurantExtraGroups(restaurantId)
		case "failed_jobs":
			response.FailedJobs = GetRestaurantFailedJobs(restaurantId)
		case "jobs":
			response.Jobs = GetRestaurantJobs(restaurantId)
		case "migrations":
			response.Migrations = GetRestaurantMigrations(restaurantId)
		case "orders":
			response.Orders = GetOrdersFromRestaurant(restaurantId)
		case "order_actions":
			response.OrderActions = GetRestaurantOrderActions(restaurantId)
		case "order_products":
			response.OrderProducts = GetRestaurantOrderProducts(restaurantId)
		case "password_resets":
			response.PasswordResets = GetRestaurantPasswordResets(restaurantId)
		case "personal_access_tokens":
			response.PersonalAccessTokens = GetRestaurantPersonalAccessTokens(restaurantId)
		case "phone_verifications":
			response.PhoneVerifications = GetRestaurantPhoneVerifications(restaurantId)
		case "point_actions":
			response.PointActions = GetRestaurantPointActions(restaurantId)
		case "products":
			response.Products = GetRestaurantProducts(restaurantId)
		case "product_extra_groups":
			response.ProductExtraGroups = GetRestaurantProductExtraGroups(restaurantId)
		case "product_variants":
			response.ProductVariants = GetRestaurantProductVariants(restaurantId)
		case "product_variant_options":
			response.ProductVariantOptions = GetRestaurantProductVariantOptions(restaurantId)
		case "reservations":
			response.Reservations = GetRestaurantReservations(restaurantId)
		case "settings":
			response.Settings = GetRestaurantSettings(restaurantId)
		}
	}
	response.Id, _ = strconv.Atoi(restaurantId)
	return
}

func GetCompanyFromId(companyId string) (response models.Companies) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"companies").Where("id = ?", companyId).First(&response)
	return
}

func GetCompanyRestaurants(companyId string) (response []models.Restaurant) {
	gormDB.Table(viper.GetString("db.admindb")+"."+"restaurants").Where("company_id = ?", companyId).Scan(&response)
	return
}

func GetRestaurantFromVenueId(venueId string) (response models.Restaurant) {
	gormDB.Table(viper.GetString("db.admindb") + "." + "restaurants").Where(datatypes.JSONQuery("info").Equals(venueId, "venue_id")).First(&response)
	return
}

func GetRestaurantsFromTermianlId(terminalId string) (response []models.Restaurant) {
	gormDB.Table(viper.GetString("db.admindb") + "." + "restaurants").Where(datatypes.JSONQuery("info").Equals(terminalId, "terminal_id")).Scan(&response)
	return
}

func GetRestaurantFromTerminalId(terminalId string) (response models.Restaurant) {
	gormDB.Table(viper.GetString("db.admindb") + "." + "restaurants").Where(datatypes.JSONQuery("info").Equals(terminalId, "terminal_id")).First(&response)
	return
}
