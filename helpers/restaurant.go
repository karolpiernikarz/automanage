package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"gorm.io/datatypes"

	"github.com/karolpiernikarz/automanage/utils"

	"github.com/google/uuid"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/spf13/viper"
	"github.com/xuri/excelize/v2"
)

func GetRestaurantExtras(restaurantId string) (extras []models.RestaurantExtras) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "extras").Scan(&extras)
	return
}

func getOrderStatusString(orderStatus int) (status string) {
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

func getOrderTypeString(orderType int) (oType string) {
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

func getDeliveryTypeString(deliveryType int) (dType string) {
	switch deliveryType {
	case 0:
		return "asap"
	case 1:
		return "preorder"
	}
	return "unknown"
}

func getPaymentType(payment string) (paymentType string) {
	switch payment {
	case "IN_RESTAURANT":
		return "CASH"
	case "IN_DOOR":
		return "CASH"
	}
	return "ONLINE"
}

func SetColumnValuesForReportsXlsxFile(f *excelize.File) *excelize.File {
	sheetName := "Orders"
	index, _ := f.NewSheet(sheetName)
	columnIndex := 1
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Resid")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "ResName")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "OrderId")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "OrderNo")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Created Date")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Delivery Date")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "C.Id")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "C.Name")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "C.Email")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "C.Phone")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Address")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Status")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "OrderType")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "DeliveryType")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "CurrierType")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "SubTotal")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Discount")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "DeliveryFee")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "BagFee")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "ServiceFee")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Before Tax")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Tax%")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Tax")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Item Count")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Coupon Code")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Allawonce Used")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Deduction")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total After Deduction")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Bonus")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Payment Type")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Payment Id")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Notes")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Distance")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "OrderURL")

	//set width
	f.SetColWidth(sheetName, "G", "G", 4)
	f.SetColWidth(sheetName, "B", "B", 18)
	f.SetColWidth(sheetName, "C", "C", 18)
	f.SetColWidth(sheetName, "E", "F", 18)
	f.SetColWidth(sheetName, "H", "H", 15)
	f.SetColWidth(sheetName, "I", "I", 25)
	f.SetColWidth(sheetName, "AC", "AC", 10)
	f.SetColWidth(sheetName, "O", "R", 4)
	f.SetColWidth(sheetName, "T", "U", 5)
	f.SetActiveSheet(index)
	return f
}

func AddOrderToXlsxFile(count int, f *excelize.File, r models.Restaurant, o models.RestaurantOrders) (int, *excelize.File) {
	columNumber := strconv.Itoa(count)
	columnIndex := 1
	sheetName := "Orders"
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, r.ID)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, r.Name)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, o.PaymentId)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, o.OrderNumber)
	columnIndex++
	createdAt, _ := time.Parse("2006-01-02 15:04:05", o.CreatedAt)
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, createdAt)
	columnIndex++
	orderDate, _ := time.Parse("2006-01-02 15:04:05", o.Date)
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, orderDate)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, o.CustomerId)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, o.Customer.Data().Name)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, o.Customer.Data().Email)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, o.Customer.Data().Phone)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, o.Address.Data().Detail)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, getOrderStatusString(o.Status))
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, getOrderTypeString(o.Type))
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, getDeliveryTypeString(o.IsPreOrder))
	columnIndex++
	currierType := ""
	if o.Type == 1 {
		if r.Info.Data().CurrierType == "none" {
			currierType = "restaurant"
		} else {
			currierType = r.Info.Data().CurrierType
		}
	}
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, currierType)
	columnIndex++
	subTotal, _ := strconv.ParseFloat(o.Prices.Data().SubTotal.String(), 64)
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, subTotal)
	columnIndex++
	totalDiscount := 0.0
	if o.Prices.Data().Discount.Data().Total.String() != "null" {
		totalDiscount, _ = strconv.ParseFloat(o.Prices.Data().Discount.Data().Total.String(), 64)
	}
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalDiscount)
	columnIndex++
	deliveryFee, _ := strconv.ParseFloat(strings.ReplaceAll(o.Prices.Data().Delivery.String(), "\"", ""), 64)
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, deliveryFee)
	columnIndex++
	bagFee, _ := strconv.ParseFloat(strings.ReplaceAll(o.Prices.Data().Bag.String(), "\"", ""), 64)
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, bagFee)
	columnIndex++
	paymentFee, _ := strconv.ParseFloat(strings.ReplaceAll(o.Prices.Data().PaymentFee.String(), "\"", ""), 64)
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, paymentFee)
	columnIndex++
	total, _ := strconv.ParseFloat(o.Prices.Data().Total.String(), 64)
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, total/1.25)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, "25")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, total-total/1.25)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, total)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, "")
	columnIndex++
	if o.Prices.Data().Discount.Data().Code.String() != "null" {
		f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, o.Prices.Data().Discount.Data().Code.String())
	} else {
		f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, "")
	}
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, "")
	columnIndex++
	deduction := 0.0
	if o.Prices.Data().Deduction.String() != "null" {
		deduction, _ = strconv.ParseFloat(strings.ReplaceAll(o.Prices.Data().Deduction.String(), "\"", ""), 64)
	}
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, deduction)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, total)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, "")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, getPaymentType(o.PaymentType))
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, o.Payment.Data().Id)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, o.Note)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, "")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, r.Website+"/order/"+o.PaymentId+"/status")
	count++
	return count, f
}

func AddRestaurantReportsToXlsxFile(count int, f *excelize.File, r models.Restaurant, orders []models.RestaurantOrders) (int, *excelize.File) {
	columNumber := strconv.Itoa(count)
	columnIndex := 1
	sheetName := "Restaurants"

	totalOrders := len(orders)
	totalPreorder := 0
	totalAsap := 0
	totalDiscount := 0.0
	totalBag := 0.0
	totalService := 0.0
	totalSubtotal := 0.0
	total := 0.0
	totalOnline := 0.0
	totalCash := 0.0
	totalDelivery := 0
	totalPickup := 0
	totalTable := 0
	totalCouponUsed := 0

	for _, order := range orders {
		Bag := 0.0
		if order.Prices.Data().Bag.String() != "null" {
			Bag, _ = strconv.ParseFloat(strings.ReplaceAll(order.Prices.Data().Bag.String(), "\"", ""), 64)
		}
		Service := 0.0
		if order.Prices.Data().PaymentFee.String() != "null" {
			Service, _ = strconv.ParseFloat(strings.ReplaceAll(order.Prices.Data().PaymentFee.String(), "\"", ""), 64)
		}
		Discount := 0.0
		if order.Prices.Data().Discount.Data().Total.String() != "null" {
			Discount, _ = strconv.ParseFloat(order.Prices.Data().Discount.Data().Total.String(), 64)
		}
		Subtotal := 0.0
		if order.Prices.Data().SubTotal.String() != "null" {
			Subtotal, _ = strconv.ParseFloat(order.Prices.Data().SubTotal.String(), 64)
		}
		totalP := 0.0
		if order.Prices.Data().Total.String() != "null" {
			totalP, _ = strconv.ParseFloat(order.Prices.Data().Total.String(), 64)
		}
		cash := 0.0
		online := 0.0
		if order.PaymentType == "IN_RESTAURANT" || order.PaymentType == "IN_DOOR" {
			cash, _ = strconv.ParseFloat(order.Prices.Data().Total.String(), 64)
		} else {
			online, _ = strconv.ParseFloat(order.Prices.Data().Total.String(), 64)
		}

		if order.IsPreOrder == 1 {
			totalPreorder++
		} else {
			totalAsap++
		}

		totalDiscount += Discount
		totalBag += Bag
		totalService += Service
		totalSubtotal += Subtotal
		total += totalP
		if order.PaymentType == "IN_RESTAURANT" || order.PaymentType == "IN_DOOR" {
			totalCash += cash
		} else {
			totalOnline += online
		}
		if order.Type == 1 {
			totalDelivery++
		} else if order.Type == 0 {
			totalPickup++
		} else {
			totalTable++
		}
		if order.Prices.Data().Discount.Data().Code.String() != "null" {
			totalCouponUsed++
		}
	}

	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, r.ID)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, r.Name)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalOrders)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalPreorder)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalAsap)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalDiscount)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalBag)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalService)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalSubtotal)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, total)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalOnline)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalCash)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalDelivery)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalPickup)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalTable)
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+columNumber, totalCouponUsed)

	count++
	return count, f
}

func SetColumnValuesForRestaurantReportsXlsxFile(f *excelize.File) *excelize.File {
	sheetName := "Restaurants"
	index, _ := f.NewSheet(sheetName)
	columnIndex := 1
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Resid")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "ResName")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Orders")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Preorder")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Asap")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Discount")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Bag")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Service")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Subtotal")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Online")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Cash")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Delivery")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Pickup")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Table")
	columnIndex++
	f.SetCellValue(sheetName, convertToColumnName(columnIndex)+"1", "Total Coupon Used")

	//set width
	f.SetColWidth(sheetName, "B", "B", 20)
	f.SetActiveSheet(index)
	return f
}

func CreateXlsxFileFromRestaruantsReportWithTime(restaruants []models.Restaurant, startTime time.Time, endTime time.Time) (file string, err error) {
	sheetName := "Restaurants"
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, _ := f.NewSheet(sheetName)

	f = SetColumnValuesForRestaurantReportsXlsxFile(f)
	var count = 2
	for _, restaurant := range restaruants {
		orders := GetOrdersFromRestaurantWithTime(strconv.FormatUint(uint64(restaurant.ID), 10), startTime, endTime)
		count, f = AddRestaurantReportsToXlsxFile(count, f, restaurant, orders)
	}
	f.SetActiveSheet(index)
	filename := uuid.New()
	file = filepath.Join(viper.GetString("app.datadir"), filename.String()+".xlsx")
	if err = f.SaveAs(file); err != nil {
		return
	}
	return
}

func CreateXlsxFileFromAllOrders(allOrdersArray []models.AllRestaurantsOrders) (file string, err error) {
	sheetName := "Orders"
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, _ := f.NewSheet(sheetName)

	f = SetColumnValuesForReportsXlsxFile(f)
	var count = 2
	for _, allOrders := range allOrdersArray {
		for _, order := range allOrders.Orders {
			count, f = AddOrderToXlsxFile(count, f, allOrders.Restaurant, order)
		}
	}
	f.SetActiveSheet(index)
	filename := uuid.New()
	file = filepath.Join(viper.GetString("app.datadir"), filename.String()+".xlsx")
	if err = f.SaveAs(file); err != nil {
		return
	}
	return
}

func GetOrdersFromRestaurant(restaurantId string) (orders []models.RestaurantOrders) {
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"orders").Where("status <> ?", 0).Scan(&orders)
	return
}

func GetOrdersFromRestaurantWithTime(restaurantId string, timeStart time.Time, timeEnd time.Time) (orders []models.RestaurantOrders) {
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"orders").Where("created_at >= ? AND created_at <= ?",
		timeStart.Format("2006-01-02 15:04:05"),
		timeEnd.Format("2006-01-02 15:04:05")).Where("status <> ?", 0).Find(&orders)
	return
}

func GetOrdersFromRestaurantWithTimeUsingDate(restaurantId string, timeStart time.Time, timeEnd time.Time) (orders []models.RestaurantOrders) {
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"orders").Where("date >= ? AND date <= ?",
		timeStart.Format("2006-01-02 15:04:05"),
		timeEnd.Format("2006-01-02 15:04:05")).Where("status <> ?", 0).Find(&orders)
	return
}

func GetOrdersFromRestaurantWithTime_Test(timeStart time.Time, timeEnd time.Time) (orders []models.RestaurantOrders) {
	gormDB.Table(viper.GetString("test.database")+"."+"orders").Where("created_at >= ? AND created_at <= ?",
		timeStart.Format("2006-01-02 15:04:05"),
		timeEnd.Format("2006-01-02 15:04:05")).Where("status <> ?", 0).Find(&orders)
	return
}

func GetOrdersFromAllRestaurants() (allOrdersArray []models.AllRestaurantsOrders) {
	restaurants := GetAllRestaurants()
	var allOrders models.AllRestaurantsOrders
	for _, restaurant := range restaurants {
		orders := GetOrdersFromRestaurant(strconv.FormatUint(uint64(restaurant.ID), 10))
		allOrders.Restaurant = restaurant
		allOrders.Orders = orders
		allOrdersArray = append(allOrdersArray, allOrders)
	}
	return
}

func GetOrdersFromAllRestaurantsWithTime(timeStart time.Time, timeEnd time.Time) (allOrdersArray []models.AllRestaurantsOrders) {
	restaurants := GetAllRestaurants()
	var allOrders models.AllRestaurantsOrders
	for _, restaurant := range restaurants {
		orders := GetOrdersFromRestaurantWithTime(strconv.FormatUint(uint64(restaurant.ID), 10), timeStart, timeEnd)
		allOrders.Restaurant = restaurant
		allOrders.Orders = orders
		allOrdersArray = append(allOrdersArray, allOrders)
	}
	return
}

func GetOrdersFromAllRestaurantsWithTimeUsingDate(timeStart time.Time, timeEnd time.Time) (allOrdersArray []models.AllRestaurantsOrders) {
	restaurants := GetAllRestaurants()
	var allOrders models.AllRestaurantsOrders
	for _, restaurant := range restaurants {
		orders := GetOrdersFromRestaurantWithTimeUsingDate(strconv.FormatUint(uint64(restaurant.ID), 10), timeStart, timeEnd)
		allOrders.Restaurant = restaurant
		allOrders.Orders = orders
		allOrdersArray = append(allOrdersArray, allOrders)
	}
	return
}

func GetOrdersFromRestaurants(restaurants []string) (allOrdersArray []models.AllRestaurantsOrders) {
	restaurantInfos := GetRestaurants(restaurants)
	var allOrders models.AllRestaurantsOrders
	for _, restaurant := range restaurantInfos {
		orders := GetOrdersFromRestaurant(strconv.FormatUint(uint64(restaurant.ID), 10))
		allOrders.Restaurant = restaurant
		allOrders.Orders = orders
		allOrdersArray = append(allOrdersArray, allOrders)
	}
	return
}

func GetOrdersFromRestaurantsWithTime(restaurants []string, timeStart time.Time, timeEnd time.Time) (allOrdersArray []models.AllRestaurantsOrders) {
	restaurantInfos := GetRestaurants(restaurants)
	var allOrders models.AllRestaurantsOrders
	for _, restaurant := range restaurantInfos {
		orders := GetOrdersFromRestaurantWithTime(strconv.FormatUint(uint64(restaurant.ID), 10), timeStart, timeEnd)
		allOrders.Restaurant = restaurant
		allOrders.Orders = orders
		allOrdersArray = append(allOrdersArray, allOrders)
	}
	return
}

func GetOrdersFromRestaurantsWithTimeUsingDate(restaurants []string, timeStart time.Time, timeEnd time.Time) (allOrdersArray []models.AllRestaurantsOrders) {
	restaurantInfos := GetRestaurants(restaurants)
	var allOrders models.AllRestaurantsOrders
	for _, restaurant := range restaurantInfos {
		orders := GetOrdersFromRestaurantWithTimeUsingDate(strconv.FormatUint(uint64(restaurant.ID), 10), timeStart, timeEnd)
		allOrders.Restaurant = restaurant
		allOrders.Orders = orders
		allOrdersArray = append(allOrdersArray, allOrders)
	}
	return
}

func GetRestaurantCustomers(restaurantId string) (customers []models.RestaurantCustomers) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "customers").Scan(&customers)
	return
}

func GetRestaurantAddresses(restaurantId string) (addresses []models.RestaurantAddresses) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "addresses").Scan(&addresses)
	return
}

func GetRestaurantAttributes(restaurantId string) (attributes []models.RestaurantAttributes) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "attributes").Scan(&attributes)
	return
}

func GetRestaurantCarts(restaurantId string) (carts []models.RestaurantCarts) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "carts").Scan(&carts)
	return
}

func GetRestaurantCoupons(restaurantId string) (coupons []models.RestaurantCoupons) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "coupons").Scan(&coupons)
	return
}

func GetRestaurantCategories(restaurantId string) (categories []models.RestaurantCategories) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "categories").Scan(&categories)
	return
}

func GetRestaurantFailedJobs(restaurantId string) (failedJobs []models.RestaurantFailed_Jobs) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "failed_jobs").Scan(&failedJobs)
	return
}

func GetRestaurantMigrations(restaurantId string) (migrations []models.RestaurantMigrations) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "migrations").Scan(&migrations)
	return
}

func GetRestaurantOrderActions(restaurantId string) (orderActions []models.RestaurantOrder_Actions) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "order_actions").Scan(&orderActions)
	return
}

func GetRestaurantOrderProducts(restaurantId string) (orderProducts []models.RestaurantOrder_Products) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "order_products").Scan(&orderProducts)
	return
}

func GetRestaurantPasswordResets(restaurantId string) (passwordResets []models.RestaurantPassword_Resets) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "password_resets").Scan(&passwordResets)
	return
}

func GetRestaurantProducts(restaurantId string) (products []models.RestaurantProducts) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "products").Scan(&products)
	return
}

func GetRestaurantProductExtraGroups(restaurantId string) (productExtraGroups []models.RestaurantProduct_Extra_Groups) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "product_extra_groups").Scan(&productExtraGroups)
	return
}

func GetRestaurantProductVariants(restaurantId string) (productVariants []models.RestaurantProduct_Variants) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "product_variants").Scan(&productVariants)
	return
}

func GetRestaurantProductVariantOptions(restaurantId string) (productVariantOptions []models.RestaurantProduct_Variant_Options) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "product_variant_options").Scan(&productVariantOptions)
	return
}

func GetRestaurantSettings(restaurantId string) (restaurantSettings []models.RestaurantSettings) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "settings").Scan(&restaurantSettings)
	return
}

func GetRestaurantOrderFromId(restaurantId string, orderId string) (order models.RestaurantOrders) {
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"orders").Where("id = ?", orderId).Find(&order)
	return
}

func GetRestaurantProductFromId(restaurantId string, productId string) (product models.RestaurantProducts) {
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"products").Where("id = ?", productId).Find(&product)
	return
}

func GetRestaurantAddressFromId(restaurantId string, addressId string) (address models.RestaurantAddresses) {
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"addresses").Where("id = ?", addressId).Find(&address)
	return
}

func GetRestaurantAttributeFromName(restaurantId string, attributeName string) (attribute models.RestaurantAttributes) {
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"attributes").Where("name = ?", attributeName).Find(&attribute)
	return
}

func GetRestaurantCartFromCartId(restaurantId string, cartId string) (cart []models.RestaurantCarts) {
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"carts").Where("cart_id = ?", cartId).Find(&cart)
	return
}

func GetRestaurantCategoryFromId(restaurantId string, categoryId string) (category models.RestaurantCategories) {
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"categories").Where("id = ?", categoryId).Find(&category)
	return
}

func GetRestaurantAdmins(restaurantId string) (admins []models.RestaurantAdmins) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "admins").Scan(&admins)
	return
}

func GetRestaurantCampaigns(restaurantId string) (campaigns []models.RestaurantCampaigns) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "campaigns").Scan(&campaigns)
	return
}

func GetRestaurantCustomerPoints(restaurantId string) (customerPoints []models.RestaurantCustomer_Points) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "customer_points").Scan(&customerPoints)
	return
}

func GetRestaurantExtraGroups(restaurantId string) (extraGroups []models.RestaurantExtra_Groups) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "extra_groups").Scan(&extraGroups)
	return
}

func GetRestaurantJobs(restaurantId string) (jobs []models.RestaurantJobs) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "jobs").Scan(&jobs)
	return
}

func GetRestaurantPersonalAccessTokens(restaurantId string) (personalAccessTokens []models.RestaurantPersonal_Access_Tokens) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "personal_access_tokens").Scan(&personalAccessTokens)
	return
}

func GetRestaurantPhoneVerifications(restaurantId string) (phoneVerifications []models.RestaurantPhone_Verifications) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "phone_verifications").Scan(&phoneVerifications)
	return
}

func GetRestaurantPointActions(restaurantId string) (pointActions []models.RestaurantPoint_Actions) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "point_actions").Scan(&pointActions)
	return
}

func GetRestaurantReservations(restaurantId string) (reservations []models.RestaurantReservations) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "reservations").Scan(&reservations)
	return
}

func GetRestaurantOrderFromOrderNumber(restaurantId string, orderNumber string) (order models.RestaurantOrders) {
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"orders").Where("order_number = ?", orderNumber).Find(&order)
	return
}

// GetLastNOrdersWithStatusNone0 need testing
func GetLastNOrdersWithStatusNone0(restaurantId string, n int) (orders []models.RestaurantOrders) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "orders").Where("status != 0").Order("id desc").Limit(n).Find(&orders)
	return
}

func GetLastOrderByEmail(restaurantId string, email string) (order models.RestaurantOrders, err error) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "orders").Where(datatypes.JSONQuery("customer").Equals(email, "email")).Order("id desc").First(&order)
	// return error if no record found
	if order.Id == 0 {
		return order, sql.ErrNoRows
	}
	return
}

func SearchLastNOrderByEmail(restaurantId string, email string, n int) (order models.RestaurantOrders, err error) {
	// get orders from restaurant with descending order
	orders := GetLastNOrdersWithStatusNone0(restaurantId, n)
	// search the email
	for _, order := range orders {
		if order.Customer.Data().Email == email {
			return order, nil
		}
	}
	return order, sql.ErrNoRows
}

func GetRestaurantCoupon(restaurantId string, couponName string, couponId string) (coupon models.RestaurantCoupons, exist bool) {
	if couponName != "" {
		result := gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"coupons").Where("name = ?", couponName).Find(&coupon)
		exist = result.RowsAffected != 0
	} else {
		result := gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"coupons").Where("id = ?", couponId).Find(&coupon)
		exist = result.RowsAffected != 0
	}
	return
}

func GetRestaurantCustomer(restaurantId string, customerId string, customerEmail string) (customer models.RestaurantCustomers, exist bool) {
	if customerId != "" {
		result := gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"customers").Where("id = ?", customerId).Find(&customer)
		exist = result.RowsAffected != 0
	} else {
		result := gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"customers").Where("email = ?", customerEmail).Find(&customer)
		exist = result.RowsAffected != 0
	}
	return
}

func IsRestaurantOpen(restaurantId string) (isOpen bool) {
	restaurantSettings := GetRestaurantSettings(restaurantId)
	// get the current time
	currentTime := time.Now()
	// get the timezone from settings and convert the time
	loc, _ := time.LoadLocation(viper.GetString("app.timezone"))
	currentTime = currentTime.In(loc)
	// get the current day
	currentDay := currentTime.Weekday()
	// search the "activeDays" from settings
	for _, settings := range restaurantSettings {
		if settings.Name == "activeDays" {
			// separate ActiveDays
			activeDays := strings.Split(settings.Value, ",")
			changeActiveDaysName(&activeDays)
			// compare the day
			open := utils.StringInSlice(currentDay.String(), activeDays)
			if !open {
				return false
			}
			break
		}
	}
	// search the "time_"+currentDay.String()+"_open" from settings
	for _, settings := range restaurantSettings {
		if settings.Name == "time_"+currentDay.String()+"_open" {
			// get the current time
			currentTime := currentTime.Format("15:04")
			// get the open time
			openAndCloseTime := settings.Value
			// separate the open and close time
			openTime := openAndCloseTime[0:5]
			closeTime := openAndCloseTime[6:11]
			if closeTime == "00:00" {
				closeTime = "23:59"
			}
			if closeTime < openTime {
				closeTime = "23:59"
			}
			// compare the time
			if currentTime >= openTime && currentTime <= closeTime {
				isOpen = true
				return
			}
		}
	}
	return false
}

func GetTheOpenTimes(restaurantSettings *models.RestaurantSettingsByName) []models.RestaurantOpenTimes {
	// separate ActiveDays
	activeDays := strings.Split(restaurantSettings.ActiveDays, ",")
	changeActiveDaysName(&activeDays)
	openTimes := make([]models.RestaurantOpenTimes, 7)
	openTimes[0].Day = "Sunday"
	openTimes[1].Day = "Monday"
	openTimes[2].Day = "Tuesday"
	openTimes[3].Day = "Wednesday"
	openTimes[4].Day = "Thursday"
	openTimes[5].Day = "Friday"
	openTimes[6].Day = "Saturday"
	for i := 0; i < 7; i++ {
		openTimes[i].IsOpen = utils.StringInSlice(openTimes[i].Day, activeDays)
		openTimes[i].Open = getOpenTimeOpen(restaurantSettings, i)
		openTimes[i].Delivery = getOpenTimeDelivery(restaurantSettings, i)
		openTimes[i].Pickup = getOpenTimePickup(restaurantSettings, i)
	}
	return openTimes
}

func WhenRestaurantOpen(restaurantId string) (openingTime time.Time, closingTime time.Time, isOpen bool) {
	restaurantSettings := GetRestaurantSettingsByName(restaurantId)
	if restaurantSettings.OrderStatus == "0" {
		isOpen = false
		return
	}
	openTimes := GetTheOpenTimes(&restaurantSettings)
	currentTime := time.Now()
	loc, _ := time.LoadLocation(viper.GetString("app.timezone"))
	currentTime = currentTime.In(loc)

	for i := 0; i < 7; i++ {
		workingTime := currentTime
		if i != 0 {
			workingTime = currentTime.Add(time.Hour * 24)
		}
		if openTimes[workingTime.Weekday()].IsOpen {
			timeStr := openTimes[workingTime.Weekday()].Open

			openHour, err := strconv.Atoi(timeStr[0:2])
			if err != nil {
				fmt.Printf("WhenRestaurantOpen: unable to parse open hour from '%s': %v\n", timeStr, err)
				continue
			}

			openMinute, err := strconv.Atoi(timeStr[3:5])
			if err != nil {
				fmt.Printf("WhenRestaurantOpen: unable to parse open minute from '%s': %v\n", timeStr, err)
				continue
			}

			closeHour, err := strconv.Atoi(timeStr[6:8])
			if err != nil {
				fmt.Printf("WhenRestaurantOpen: unable to parse close hour from '%s': %v\n", timeStr, err)
				continue
			}

			closeMinute, err := strconv.Atoi(timeStr[9:11])
			if err != nil {
				fmt.Printf("WhenRestaurantOpen: unable to parse close minute from '%s': %v\n", timeStr, err)
				continue
			}
			closingTime = time.Date(workingTime.Year(), workingTime.Month(), workingTime.Day(), closeHour, closeMinute, 0, 0, workingTime.Location())
			openingTime = time.Date(workingTime.Year(), workingTime.Month(), workingTime.Day(), openHour, openMinute, 0, 0, workingTime.Location())
			currentTime := currentTime.Format("15:04")
			// get the open time
			openAndCloseTime := openTimes[workingTime.Weekday()].Open
			// separate the open and close time
			openTime := openAndCloseTime[0:5]
			closeTime := openAndCloseTime[6:11]
			if closeTime == "00:00" {
				closeTime = "23:59"
			}
			if closeTime < openTime {
				closeTime = "23:59"
				closingTime = closingTime.Add(time.Hour * 24)
			}
			// compare the time
			if currentTime >= openTime && currentTime <= closeTime {
				isOpen = true
			}
			return
		}
	}
	return
}

func GetOrdersByEmail(restaurantId string, email string) (orders []models.RestaurantOrders) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "orders").Where(datatypes.JSONQuery("customer").Equals(email, "email")).Scan(&orders)
	return
}

func GetOrdersByPhone(restaurantId string, phone string) (orders []models.RestaurantOrders) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "orders").Where(datatypes.JSONQuery("customer").Equals(phone, "phone")).Scan(&orders)
	return
}

func GetOrdersByAddress(restaurantId string, address string) (orders []models.RestaurantOrders) {
	gormDB.Table(viper.GetString("db.prefix") + restaurantId + "." + "orders").Where(datatypes.JSONQuery("customer").Equals(address, "address")).Scan(&orders)
	return
}

func GetAllOrdersByEmail(email string) (orders []models.RestaurantOrders) {
	restaurants := GetAllRestaurants()
	for _, restaurant := range restaurants {
		orders = append(orders, GetOrdersByEmail(strconv.FormatUint(uint64(restaurant.ID), 10), email)...)
	}
	return
}

func GetRestaurantSettingsByName(restaurantId string) (settings models.RestaurantSettingsByName) {
	var names []sql.NullString
	var values []sql.NullString
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"settings").Select("name").Pluck("name", &names)
	gormDB.Table(viper.GetString("db.prefix")+restaurantId+"."+"settings").Select("value").Pluck("value", &values)
	// create a key value map
	settingsMap := make(map[string]string)
	for i := 0; i < len(names); i++ {
		settingsMap[names[i].String] = values[i].String
	}
	// settingMap to json
	settingsJson, _ := json.Marshal(settingsMap)
	// json to struct
	json.Unmarshal(settingsJson, &settings)
	return
}

func GetAllRestaurantsOrderWithTime(startTime time.Time, endTime time.Time) (orders []models.RestaurantTables) {
	restaurantIds := GetAllRestaurantIds()
	// don't wait for the goroutines to finish
	var wg sync.WaitGroup
	// create a channel to receive the results
	restaurantCount := len(restaurantIds)
	var result = make([]models.RestaurantTables, restaurantCount)

	// start the goroutines
	for i, restaurantId := range restaurantIds {
		wg.Add(1)
		i := i
		go func(restaurantId string) {
			orders := GetOrdersFromRestaurantWithTime(restaurantId, startTime, endTime)
			id, _ := strconv.Atoi(restaurantId)
			result[i].Id = id
			result[i].Orders = orders
			wg.Done()
		}(restaurantId)
	}
	wg.Wait()
	return result
}

func convertToColumnName(index int) (columnName string) {

	for index > 0 {
		index--
		columnName = string(rune('A'+index%26)) + columnName
		index /= 26
	}

	return
}

func changeActiveDaysName(activeDays *[]string) {
	for i, day := range *activeDays {
		switch day {
		case "Pazartesi":
			(*activeDays)[i] = "Monday"
		case "Salı":
			(*activeDays)[i] = "Tuesday"
		case "Çarşamba":
			(*activeDays)[i] = "Wednesday"
		case "Perşembe":
			(*activeDays)[i] = "Thursday"
		case "Cuma":
			(*activeDays)[i] = "Friday"
		case "Cumartesi":
			(*activeDays)[i] = "Saturday"
		case "Pazar":
			(*activeDays)[i] = "Sunday"
		}
	}
}

func getOpenTimeDelivery(restaurantSettings *models.RestaurantSettingsByName, i int) (response string) {
	switch i {
	case 0:
		response = restaurantSettings.TimeSundayDelivery
	case 1:
		response = restaurantSettings.TimeMondayDelivery
	case 2:
		response = restaurantSettings.TimeTuesdayDelivery
	case 3:
		response = restaurantSettings.TimeWednesdayDelivery
	case 4:
		response = restaurantSettings.TimeThursdayDelivery
	case 5:
		response = restaurantSettings.TimeFridayDelivery
	case 6:
		response = restaurantSettings.TimeSaturdayDelivery
	}
	return
}

func getOpenTimeOpen(restaurantSettings *models.RestaurantSettingsByName, i int) (response string) {
	switch i {
	case 0:
		response = restaurantSettings.TimeSundayOpen
	case 1:
		response = restaurantSettings.TimeMondayOpen
	case 2:
		response = restaurantSettings.TimeTuesdayOpen
	case 3:
		response = restaurantSettings.TimeWednesdayOpen
	case 4:
		response = restaurantSettings.TimeThursdayOpen
	case 5:
		response = restaurantSettings.TimeFridayOpen
	case 6:
		response = restaurantSettings.TimeSaturdayOpen
	}
	return
}

func getOpenTimePickup(restaurantSettings *models.RestaurantSettingsByName, i int) (response string) {
	switch i {
	case 0:
		response = restaurantSettings.TimeSundayPickup
	case 1:
		response = restaurantSettings.TimeMondayPickup
	case 2:
		response = restaurantSettings.TimeTuesdayPickup
	case 3:
		response = restaurantSettings.TimeWednesdayPickup
	case 4:
		response = restaurantSettings.TimeThursdayPickup
	case 5:
		response = restaurantSettings.TimeFridayPickup
	case 6:
		response = restaurantSettings.TimeSaturdayPickup
	}
	return
}
