package fake

import (
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	newAdmin "github.com/karolpiernikarz/automanage/services/newadmin"
)

func init() {
	gofakeit.Seed(0)
}

// GetFakeOrder returns a fake order
func GetFakeOrder() (order newAdmin.Order) {
	order.OrderNumber = strconv.Itoa(gofakeit.IntRange(1000000000, 9999999999))
	order.RestaurantId = 1
	order.CustomerId = 0
	order.Type = gofakeit.RandomInt([]int{0, 1})
	order.Customer = newAdmin.JSONB{
		"name":    gofakeit.FirstName(),
		"surname": gofakeit.LastName(),
		"phone":   gofakeit.Phone(),
		"email":   gofakeit.Email(),
	}

	addressdetail := ""
	if order.Type == 1 {
		addressdetail = gofakeit.Address().Address
	} else {
		addressdetail = ""
	}

	order.Address = newAdmin.JSONB{
		"name":    gofakeit.FirstName(),
		"surname": gofakeit.LastName(),
		"email":   gofakeit.Email(),
		"placeid": gofakeit.UUID(),
		"detail":  addressdetail,
		"phone":   gofakeit.Phone(),
		"lat":     gofakeit.Latitude(),
		"long":    gofakeit.Longitude(),
	}

	deliveryprice := 0
	sub_total := gofakeit.Number(100, 1000)
	bagfee := gofakeit.Number(0, 10)
	payment_fee := gofakeit.Number(0, 10)
	if order.Type == 1 {
		deliveryprice = gofakeit.Number(20, 150)
	} else {
		deliveryprice = 0
	}
	total := sub_total + bagfee + payment_fee + deliveryprice

	order.Prices = newAdmin.JSONB{
		"total":       total,
		"sub_total":   sub_total,
		"delivery":    deliveryprice,
		"bag":         bagfee,
		"payment_fee": payment_fee,
		"serviceFee":  0,
		"discount": newAdmin.JSONB{
			"coupon_id": "0.0",
			"type":      "percent",
			"code":      "",
			"discount":  0,
			"total":     0.0,
		},
	}

	order.Payment = newAdmin.JSONB{
		"amount":   total,
		"id":       gofakeit.Number(100, 1000),
		"currency": 208,
		"card": newAdmin.JSONB{
			"last4":     gofakeit.Number(1000, 9999),
			"name":      "",
			"exp_month": gofakeit.Number(1, 12),
			"exp_year":  gofakeit.Number(2021, 2025),
			"type":      "visa",
		},
		"receipt_url": "",
		"status":      "paid",
		"refunded":    false,
	}

	order.PaymentId = gofakeit.UUID()

	order.Status = gofakeit.RandomInt([]int{1, 2, 3, 4, 5})
	order.PaymentType = gofakeit.RandomString([]string{"CARD", "IN_RESTAURANT"})
	order.Note = gofakeit.Sentence(10)
	order.Date = time.Now()
	order.Currier = newAdmin.JSONBArray{}
	order.IsPreOrder = gofakeit.RandomInt([]int{0, 1})

	return
}
