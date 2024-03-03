package models

import (
	"gorm.io/datatypes"
)

// RestaurantAttributes not tested
type RestaurantAttributes struct {
	Id        int    `db:"id" gorm:"id" json:"id"`
	Name      string `db:"name" gorm:"name" json:"name"`
	Value     string `db:"value" gorm:"value" json:"value"`
	CreatedAt string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantOrders struct {
	Id          int                                           `db:"id" gorm:"id" json:"id"`                               // order id
	OrderNumber string                                        `db:"order_number" gorm:"order_number" json:"order_number"` // order number
	PaymentId   string                                        `db:"payment_id" gorm:"payment_id" json:"payment_id"`       // customer id
	CustomerId  int                                           `db:"customer_id" gorm:"customer_id" json:"customer_id"`    // customer id
	Customer    datatypes.JSONType[*RestaurantOrdersCustomer] `db:"customer" gorm:"customer" json:"customer"`
	Address     datatypes.JSONType[*RestaurantOrdersAddress]  `db:"address" gorm:"address" json:"address"`
	Prices      datatypes.JSONType[*RestaurantOrdersPrices]   `db:"prices" gorm:"prices" json:"prices"`
	Payment     datatypes.JSONType[*RestaurantOrdersPayment]  `db:"payment" gorm:"payment" json:"payment"`
	PaymentType string                                        `db:"payment_type" gorm:"payment_type" json:"payment_type"` // payment type. example: "IN_RESTAURANT", "IN_DOOR"
	Status      int                                           `db:"status" gorm:"status" json:"status"`                   // order status. example: 2 = "preparing", 4 = "delivered"
	Type        int                                           `db:"type" gorm:"type" json:"type"`                         // order type. example: 0 = "pickup", 1 = "delivery"
	Note        string                                        `db:"note" gorm:"note" json:"note"`                         // notes from order
	Other       datatypes.JSONType[*RestaurantOrdersOther]    `db:"other" gorm:"other" json:"other"`
	Date        string                                        `db:"date" gorm:"date" json:"date"` // time of the delivery
	CreatedAt   string                                        `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt   string                                        `db:"updated_at" gorm:"updated_at" json:"updated_at"`
	IsPreOrder  int                                           `db:"is_pre_order" gorm:"is_pre_order" json:"is_pre_order"` // 1 is preorder
}

type AllRestaurantsOrders struct {
	Restaurant Restaurant
	Orders     []RestaurantOrders
}

type RestaurantOrdersCustomer struct {
	Name    string `db:"name" gorm:"name" json:"name"`
	Surname string `db:"surname" gorm:"surname" json:"surname"`
	Email   string `db:"email" gorm:"email" json:"email"`
	Phone   string `db:"phone" gorm:"phone" json:"phone"`
}

type RestaurantOrdersAddress struct {
	Name   string `db:"name" gorm:"name" json:"name"`
	Phone  string `db:"phone" gorm:"phone" json:"phone"`
	Email  string `db:"email" gorm:"email" json:"email"`
	Detail string `db:"detail" gorm:"detail" json:"detail"`
	Lat    any    `db:"lat" gorm:"lat" json:"lat"`
	Long   any    `db:"long" gorm:"long" json:"long"`
}

type RestaurantOrdersPrices struct {
	SubTotal   datatypes.JSON                                      `db:"sub_total" gorm:"sub_total" json:"sub_total"`
	Delivery   datatypes.JSON                                      `db:"delivery" gorm:"delivery" json:"delivery"`
	Discount   datatypes.JSONType[*RestaurantOrdersPricesDiscount] `db:"discount" gorm:"discount" json:"discount"`
	Bag        datatypes.JSON                                      `db:"bag" gorm:"bag" json:"bag"`
	PaymentFee datatypes.JSON                                      `db:"payment_fee" gorm:"payment_fee" json:"payment_fee"`
	Total      datatypes.JSON                                      `db:"total" gorm:"total" json:"total"`
	BankFee    datatypes.JSON                                      `db:"bankFee" gorm:"bankFee" json:"bankFee"`
	ServiceFee datatypes.JSON                                      `db:"serviceFee" gorm:"serviceFee" json:"serviceFee"`
	Currier    datatypes.JSON                                      `db:"currier" gorm:"currier" json:"currier"`
	Deduction  datatypes.JSON                                      `db:"deduction,omitempty" gorm:"deduction,omitempty" json:"deduction,omitempty"`
}

type RestaurantOrdersPayment struct {
	Id       string                                           `db:"id" gorm:"id" json:"id,omitempty"`
	Amount   any                                              `db:"amount" gorm:"amount" json:"amount,omitempty"`
	Currency string                                           `db:"currency" gorm:"currency" json:"currency,omitempty"`
	Card     datatypes.JSONType[*RestaurantOrdersPaymentCard] `db:"card" gorm:"card" json:"card,omitempty"`
	Type     string                                           `db:"type" gorm:"type" json:"type,omitempty"`
}

type RestaurantOrdersOther struct {
	Wolt            datatypes.JSONType[*RestaurantOrdersOtherWolt] `db:"wolt,omitempty" gorm:"wolt,omitempty" json:"wolt,omitempty"`
	CurrierType     string                                         `db:"currier_type" gorm:"currier_type" json:"currier_type"`
	CurrierCallDate string                                         `db:"currier_call_date" gorm:"currier_call_date" json:"currier_call_date"`
}

type RestaurantOrdersOtherWolt struct {
	Data datatypes.JSONType[*RestaurantOrdersOtherWoltData] `db:"data,omitempty" gorm:"data,omitempty" json:"data,omitempty"`
	//Shipment datatypes.JSONType[*RestaurantOrdersOtherWoltShipment] `db:"shipment,omitempty" gorm:"shipment,omitempty" json:"shipment,omitempty"`
}

type RestaurantOrdersOtherWoltData struct {
	Status string                                                 `db:"status" gorm:"status" json:"status"`
	Data   datatypes.JSONType[*RestaurantOrdersOtherWoltDataData] `db:"data" gorm:"data" json:"data"`
}

type RestaurantOrdersOtherWoltDataData struct {
	ReferenceId string `db:"reference_id" gorm:"reference_id" json:"reference_id"`
	TrackingUrl string `db:"tracking_url" gorm:"tracking_url" json:"tracking_url"`
}

type RestaurantOrdersOtherWoltShipment struct {
	Id         string `db:"id" gorm:"id" json:"id"`
	CreatedAt  string `db:"created_at" gorm:"created_at" json:"created_at"`
	ValidUntil string `db:"valid_until" gorm:"valid_until" json:"valid_until"`
}

type RestaurantOrdersPaymentCard struct {
	Name       string `db:"name" gorm:"name" json:"name,omitempty"`
	ExpMonth   string `db:"exp_month" gorm:"exp_month" json:"exp_month,omitempty"`
	ExpYear    string `db:"exp_year" gorm:"exp_year" json:"exp_year,omitempty"`
	Last4      string `db:"last4" gorm:"last4" json:"last4,omitempty"`
	Type       string `db:"type" gorm:"type" json:"type,omitempty"`
	ReceiptUrl string `db:"receipt_url" gorm:"receipt_url" json:"receipt_url,omitempty"`
	Status     string `db:"status" gorm:"status" json:"status,omitempty"`
	Refunded   int    `db:"refunded" gorm:"refunded" json:"refunded,omitempty"`
}

type RestaurantOrdersPricesDiscount struct {
	Type     datatypes.JSON `db:"type" gorm:"type" json:"type"`
	CouponId datatypes.JSON `db:"coupon_id" gorm:"coupon_id" json:"coupon_id"`
	Code     datatypes.JSON `db:"code" gorm:"code" json:"code"`
	Discount datatypes.JSON `db:"discount" gorm:"discount" json:"discount"`
	Total    datatypes.JSON `db:"total" gorm:"total" json:"total"`
}

type RestaurantAddresses struct {
	Id          int    `db:"id" gorm:"id" json:"id"`
	CustomerId  int    `db:"customer_id" gorm:"customer_id" json:"customer_id"` // customer id
	Title       string `db:"title" gorm:"title" json:"title"`
	Description string `db:"description" gorm:"description" json:"description"`
	PlaceId     string `db:"place_id" gorm:"place_id" json:"place_id"`
	Detail      string `db:"detail" gorm:"detail" json:"detail"`
	Lat         string `db:"lat" gorm:"lat" json:"lat"`
	Long        string `db:"long" gorm:"long" json:"long"`
	IsActive    int    `db:"is_active" gorm:"is_active" json:"is_active"`
	CreatedAt   string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantCarts struct {
	Id         int                                          `db:"id" gorm:"id" json:"id"`
	CartId     string                                       `db:"cart_id" gorm:"cart_id" json:"cart_id"`
	ProductId  int                                          `db:"product_id" gorm:"product_id" json:"product_id"`
	Qty        int                                          `db:"qty" gorm:"qty" json:"qty"`
	Variants   string                                       `db:"variants" gorm:"variants" json:"variants"`
	Extras     datatypes.JSONType[*[]RestaurantCartsExtras] `db:"extras" gorm:"extras" json:"extras"`
	CreatedAt  string                                       `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt  string                                       `db:"updated_at" gorm:"updated_at" json:"updated_at"`
	CategoryId string                                       `db:"category_id" gorm:"category_id" json:"category_id"`
}

type RestaurantCartsExtras struct {
	Id             int                                             `db:"id" gorm:"id" json:"id"`
	Name           string                                          `db:"name" gorm:"name" json:"name"`
	Price          float64                                         `db:"price" gorm:"price" json:"price"`
	FormattedPrice string                                          `db:"formatted_price" gorm:"formatted_price" json:"formatted_price"`
	Group          datatypes.JSONType[*RestaurantCartsExtrasGroup] `db:"group" gorm:"group" json:"group"`
}

type RestaurantCartsExtrasGroup struct {
	Id   int    `db:"id" gorm:"id" json:"id"`
	Name string `db:"name" gorm:"name" json:"name"`
}

type RestaurantCategories struct {
	Id          int    `db:"id" gorm:"id" json:"id"`
	Name        string `db:"name" gorm:"name" json:"name"`
	Slug        string `db:"slug" gorm:"slug" json:"slug"`
	Icon        string `db:"icon" gorm:"icon" json:"icon"`
	Banner      string `db:"banner" gorm:"banner" json:"banner"`
	Sort        int    `db:"sort" gorm:"sort" json:"sort"`
	IsActive    int    `db:"is_active" gorm:"is_active" json:"is_active"`
	Description string `db:"description" gorm:"description" json:"description"`
	CreatedAt   string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
	Sunday      string `db:"sunday" gorm:"sunday" json:"sunday"`
	Monday      string `db:"monday" gorm:"monday" json:"monday"`
	Tuesday     string `db:"tuesday" gorm:"tuesday" json:"tuesday"`
	Wednesday   string `db:"wednesday" gorm:"wednesday" json:"wednesday"`
	Thursday    string `db:"thursday" gorm:"thursday" json:"thursday"`
	Friday      string `db:"friday" gorm:"friday" json:"friday"`
	Saturday    string `db:"saturday" gorm:"saturday" json:"saturday"`
}

type RestaurantCoupons struct {
	Id           int     `db:"id" gorm:"id" json:"id"`
	Name         string  `db:"name" gorm:"name" json:"name"`
	Code         string  `db:"code" gorm:"code" json:"code"`
	Status       int     `db:"status" gorm:"status" json:"status"`
	StartDate    string  `db:"start_date" gorm:"start_date" json:"start_date"`
	EndDate      string  `db:"end_date" gorm:"end_date" json:"end_date"`
	MaxUse       int     `db:"max_use" gorm:"max_use" json:"max_use"`
	UserMaxUse   int     `db:"user_max_use" gorm:"user_max_use" json:"user_max_use"`
	TotalUse     int     `db:"total_use" gorm:"total_use" json:"total_use"`
	SaleType     int     `db:"sale_type" gorm:"sale_type" json:"sale_type"`
	SaleAmount   float64 `db:"sale_amount" gorm:"sale_amount" json:"sale_amount"`
	MinCartTotal float64 `db:"min_cart_total" gorm:"min_cart_total" json:"min_cart_total"`
	UserGroup    int     `db:"user_group" gorm:"user_group" json:"user_group"`
	UserId       int     `db:"user_id" gorm:"user_id" json:"user_id"`
	RegisterDate int     `db:"register_date" gorm:"register_date" json:"register_date"`
	OrderSource  int     `db:"order_source" gorm:"order_source" json:"order_source"`
	FreeCargo    int     `db:"free_cargo" gorm:"free_cargo" json:"free_cargo"`
	OrderCount   int     `db:"order_count" gorm:"order_count" json:"order_count"`
	DeliveryType int     `db:"delivery_type" gorm:"delivery_type" json:"delivery_type"`
	CreatedAt    string  `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt    string  `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantCustomers struct {
	Id            int    `db:"id" gorm:"id" json:"id"`
	Name          string `db:"name" gorm:"name" json:"name"`
	Surname       string `db:"surname" gorm:"surname" json:"surname"`
	Email         string `db:"email" gorm:"email" json:"email"`
	Phone         string `db:"phone" gorm:"phone" json:"phone"`
	Password      string `db:"password" gorm:"password" json:"password"`
	Gender        int    `db:"gender" gorm:"gender" json:"gender"`
	IsVerified    int    `db:"is_verified" gorm:"is_verified" json:"is_verified"`
	DeletedAt     string `db:"deleted_at" gorm:"deleted_at" json:"deleted_at"`
	RememberToken string `db:"remember_token" gorm:"remember_token" json:"remember_token"`
	CreatedAt     string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt     string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
	Lang          string `db:"lang" gorm:"lang" json:"lang"`
}

type RestaurantCustomer_Points struct {
	Id         int     `db:"id" gorm:"id" json:"id"`
	CustomerId int     `db:"customer_id" gorm:"customer_id" json:"customer_id"` // customer id
	Point      float64 `db:"point" gorm:"point" json:"point"`
	CreatedAt  string  `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt  string  `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantExtras struct {
	Id         int     `db:"id" gorm:"id" json:"id"`
	GroupId    int     `db:"group_id" gorm:"group_id" json:"group_id"`
	Name       string  `db:"name" gorm:"name" json:"name"`
	Price      float64 `db:"price" gorm:"price" json:"price"`
	IsDisabled int     `db:"is_disabled" gorm:"is_disabled" json:"is_disabled"`
	IsDefault  int     `db:"is_default" gorm:"is_default" json:"is_default"`
	CreatedAt  string  `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt  string  `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantExtra_Groups struct {
	Id          int    `db:"id" gorm:"id" json:"id"`
	Name        string `db:"name" gorm:"name" json:"name"`
	DisplayName string `db:"display_name" gorm:"display_name" json:"display_name"`
	Limit       string `db:"limit" gorm:"limit" json:"limit"`
	Order       int    `db:"order" gorm:"order" json:"order"`
	CreatedAt   string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantFailed_Jobs struct {
	Id         int    `db:"id" gorm:"id" json:"id"`
	Uuid       string `db:"uuid" gorm:"uuid" json:"uuid"`
	Connection string `db:"connection" gorm:"connection" json:"connection"`
	Queue      string `db:"queue" gorm:"queue" json:"queue"`
	Payload    string `db:"payload" gorm:"payload" json:"payload"` // need to be json
	Exception  string `db:"exception" gorm:"exception" json:"exception"`
	FailedAt   string `db:"failed_at" gorm:"failed_at" json:"failed_at"`
}

type RestaurantMigrations struct {
	Id        int    `db:"id" gorm:"id" json:"id"`
	Migration string `db:"migration" gorm:"migration" json:"migration"`
	Batch     int    `db:"batch" gorm:"id" batch:"batch"`
}

type RestaurantOrder_Actions struct {
	Id        int    `db:"id" gorm:"id" json:"id"`
	OrderId   int    `db:"order_id" gorm:"order_id" json:"order_id"`
	Text      string `db:"text" gorm:"text" json:"text"`
	CreatedAt string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantOrder_Products struct {
	Id        int                                                     `db:"id" gorm:"id" json:"id"`
	OrderId   int                                                     `db:"order_id" gorm:"order_id" json:"order_id"`
	ProductId int                                                     `db:"product_id" gorm:"product_id" json:"product_id"`
	Qty       int                                                     `db:"qty" gorm:"qty" json:"qty"`
	UnitPrice float64                                                 `db:"unit_price" gorm:"unit_price" json:"unit_price"`
	Price     float64                                                 `db:"price" gorm:"price" json:"price"`
	Variants  datatypes.JSONType[*[]RestaurantOrder_ProductsVariants] `db:"variants" gorm:"variants" json:"variants"`
	Extras    datatypes.JSONType[*[]RestaurantOrder_ProductsExtras]   `db:"extras" gorm:"extras" json:"extras"`
	Comment   string                                                  `db:"comment" gorm:"comment" json:"comment"`
}
type RestaurantOrder_ProductsVariants struct {
	Variant   datatypes.JSONType[*RestaurantOrder_ProductsVariantsVariant]   `db:"variant" gorm:"variant" json:"variant"`
	Option    datatypes.JSONType[*RestaurantOrder_ProductsVariantsVariant]   `db:"option" gorm:"option" json:"option"`
	Price     string                                                         `db:"price" gorm:"price" json:"price"`
	Formatted datatypes.JSONType[*RestaurantOrder_ProductsVariantsFormatted] `db:"formatted" gorm:"formatted" json:"formatted"`
}

type RestaurantOrder_ProductsVariantsVariant struct {
	Id   int    `db:"id" gorm:"id" json:"id"`
	Name string `db:"name" gorm:"name" json:"name"`
}

type RestaurantOrder_ProductsVariantsFormatted struct {
	NonPrice string `db:"non_price" gorm:"non_price" json:"non_price"`
	AddPrice string `db:"add_price" gorm:"add_price" json:"add_price"`
}

type RestaurantOrder_ProductsExtras struct {
	Id             int                                                      `db:"id" gorm:"id" json:"id"`
	Name           string                                                   `db:"name" gorm:"name" json:"name"`
	Price          float64                                                  `db:"price" gorm:"price" json:"price"`
	FormattedPrice string                                                   `db:"formatted_price" gorm:"formatted_price" json:"formatted_price"`
	Group          datatypes.JSONType[*RestaurantOrder_ProductsExtrasGroup] `db:"group" gorm:"group" json:"group"`
}

type RestaurantOrder_ProductsExtrasGroup struct {
	Id          int    `db:"id" gorm:"id" json:"id"`
	Name        string `db:"name" gorm:"name" json:"name"`
	DisplayName string `db:"display_name" gorm:"display_name" json:"display_name"`
	Limit       string `db:"limit" gorm:"limit" json:"limit"`
}

type RestaurantPassword_Resets struct {
	Email     string `db:"email" gorm:"email" json:"email"`
	Token     string `db:"token" gorm:"token" json:"token"`
	CreatedAt string `db:"created_at" gorm:"created_at" json:"created_at"`
}

type RestaurantProducts struct {
	Id                   int                                             `db:"id" gorm:"id" json:"id"`
	CategoryId           int                                             `db:"category_id" gorm:"category_id" json:"category_id"`
	Name                 string                                          `db:"name" gorm:"name" json:"name"`
	Slug                 string                                          `db:"slug" gorm:"slug" json:"slug"`
	Image                string                                          `db:"image" gorm:"image" json:"image"`
	WithoutDiscountPrice float64                                         `db:"without_discount_price" gorm:"without_discount_price" json:"without_discount_price"`
	Price                float64                                         `db:"price" gorm:"price" json:"price"`
	Description          string                                          `db:"description" gorm:"description" json:"description"`
	Sort                 int                                             `db:"sort" gorm:"sort" json:"sort"`
	IsActive             int                                             `db:"is_active" gorm:"is_active" json:"is_active"`
	Keywords             string                                          `db:"keywords" gorm:"keywords" json:"keywords"`
	Materials            string                                          `db:"materials" gorm:"materials" json:"materials"`
	Discount             datatypes.JSONType[*RestaurantProductsDiscount] `db:"discount" gorm:"discount" json:"discount"`
	CreatedAt            string                                          `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt            string                                          `db:"updated_at" gorm:"updated_at" json:"updated_at"`
	Type                 string                                          `db:"type" gorm:"type" json:"type"`
}

type RestaurantProductsDiscount struct {
	Type                string `db:"type" gorm:"type" json:"type"`
	Percent             string `db:"percent" gorm:"percent" json:"percent"`
	FixedDiscountAmount string `db:"fixed_discount_amount" gorm:"fixed_discount_amount" json:"fixed_discount_amount"`
}

type RestaurantProduct_Extra_Groups struct {
	Id                     int    `db:"id" gorm:"id" json:"id"`
	ProductId              int    `db:"product_id" gorm:"product_id" json:"product_id"`
	ProductVariantOptionId int    `db:"product_variant_option_id" gorm:"product_variant_option_id" json:"product_variant_option_id"`
	ExtraGroupId           int    `db:"extra_group_id" gorm:"extra_group_id" json:"extra_group_id"`
	CreatedAt              string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt              string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantProduct_Variants struct {
	Id        int    `db:"id" gorm:"id" json:"id"`
	ProductId int    `db:"product_id" gorm:"product_id" json:"product_id"`
	Name      string `db:"name" gorm:"name" json:"name"`
	CreatedAt string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantProduct_Variant_Options struct {
	Id               int     `db:"id" gorm:"id" json:"id"`
	ProductVariantId int     `db:"product_variant_id" gorm:"product_variant_id" json:"product_variant_id"`
	Name             string  `db:"name" gorm:"name" json:"name"`
	Price            float64 `db:"price" gorm:"price" json:"price"`
	CreatedAt        string  `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt        string  `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantSettings struct {
	Id        int    `db:"id" gorm:"id" json:"id"`
	Name      string `db:"name" gorm:"name" json:"name"`
	Value     string `db:"value" gorm:"value" json:"value"`
	CreatedAt string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantOrdersAndProducts struct {
	Id       uint
	Orders   []RestaurantOrders
	Products []RestaurantProducts
}

type RestaurantCampaigns struct {
	Id        int    `db:"id" gorm:"id" json:"id"`
	Title     string `db:"title" gorm:"title" json:"title"`
	Slug      string `db:"slug" gorm:"slug" json:"slug"`
	Image     string `db:"image" gorm:"image" json:"image"`
	Detail    string `db:"detail" gorm:"detail" json:"detail"`
	EndDate   string `db:"end_date" gorm:"end_date" json:"end_date"`
	CreatedAt string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantJobs struct {
	Id          int    `db:"id" gorm:"id" json:"id"`
	Queue       string `db:"queue" gorm:"queue" json:"queue"`
	Payload     string `db:"payload" gorm:"payload" json:"payload"`
	Attempts    int    `db:"attempts" gorm:"attempts" json:"attempts"`
	ReservedAt  string `db:"reserved_at" gorm:"reserved_at" json:"reserved_at"`
	AvailableAt string `db:"available_at" gorm:"available_at" json:"available_at"`
	CreatedAt   string `db:"created_at" gorm:"created_at" json:"created_at"`
}

type RestaurantPersonal_Access_Tokens struct {
	Id            int    `db:"id" gorm:"id" json:"id"`
	TokenableType string `db:"tokenable_type" gorm:"tokenable_type" json:"tokenable_type"`
	TokenableId   int    `db:"tokenable_id" gorm:"tokenable_id" json:"tokenable_id"`
	Name          string `db:"name" gorm:"name" json:"name"`
	Token         string `db:"token" gorm:"token" json:"token"`
	Abilities     string `db:"abilities" gorm:"abilities" json:"abilities"`
	LastUsedAt    string `db:"last_used_at" gorm:"last_used_at" json:"last_used_at"`
	ExpiresAt     string `db:"expires_at" gorm:"expires_at" json:"expires_at"`
	CreatedAt     string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt     string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantPhone_Verifications struct {
	Id         int    `db:"id" gorm:"id" json:"id"`
	CustomerId int    `db:"customer_id" gorm:"customer_id" json:"customer_id"`
	Phone      string `db:"phone" gorm:"phone" json:"phone"`
	Code       string `db:"code" gorm:"code" json:"code"`
	CreatedAt  string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt  string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantPoint_Actions struct {
	Id         int     `db:"id" gorm:"id" json:"id"`
	CustomerId int     `db:"customer_id" gorm:"customer_id" json:"customer_id"`
	OrderId    int     `db:"order_id" gorm:"order_id" json:"order_id"`
	Amount     float64 `db:"amount" gorm:"amount" json:"amount"`
	Type       string  `db:"type" gorm:"type" json:"type"`
	CreatedAt  string  `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt  string  `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantReservations struct {
	Id         int    `db:"id" gorm:"id" json:"id"`
	CustomerId int    `db:"customer_id" gorm:"customer_id" json:"customer_id"`
	TableId    int    `db:"table_id" gorm:"table_id" json:"table_id"`
	Person     int    `db:"person" gorm:"person" json:"person"`
	Date       string `db:"date" gorm:"date" json:"date"`
	Time       string `db:"time" gorm:"time" json:"time"`
	Status     int    `db:"status" gorm:"status" json:"status"`
	CreatedAt  string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt  string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}
type RestaurantAdmins struct {
	Id        int    `db:"id" gorm:"id" json:"id"`
	Name      string `db:"name" gorm:"name" json:"name"`
	Surname   string `db:"surname" gorm:"surname" json:"surname"`
	Email     string `db:"email" gorm:"email" json:"email"`
	Phone     string `db:"phone" gorm:"phone" json:"phone"`
	Password  string `db:"password" gorm:"password" json:"password"`
	CreatedAt string `db:"created_at" gorm:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" gorm:"updated_at" json:"updated_at"`
}

type RestaurantTables struct {
	Id                    int                                 `db:"id" gorm:"id" json:"id,omitempty"`
	Addresses             []RestaurantAddresses               `db:"addresses" gorm:"addresses" json:"addresses,omitempty"`
	Admins                []RestaurantAdmins                  `db:"admins" gorm:"admins" json:"admins,omitempty"`
	Attributes            []RestaurantAttributes              `db:"attributes" gorm:"attributes" json:"attributes,omitempty"`
	Campaigns             []RestaurantCampaigns               `db:"campaigns" gorm:"campaigns" json:"campaigns,omitempty"`
	Carts                 []RestaurantCarts                   `db:"carts" gorm:"carts" json:"carts,omitempty"`
	Categories            []RestaurantCategories              `db:"categories" gorm:"categories" json:"categories,omitempty"`
	Coupons               []RestaurantCoupons                 `db:"coupons" gorm:"coupons" json:"coupons,omitempty"`
	Customers             []RestaurantCustomers               `db:"customers" gorm:"customers" json:"customers,omitempty"`
	CustomerPoints        []RestaurantCustomer_Points         `db:"customer_points" gorm:"customer_points" json:"customer_points,omitempty"`
	Extras                []RestaurantExtras                  `db:"extras" gorm:"extras" json:"extras,omitempty"`
	ExtraGroups           []RestaurantExtra_Groups            `db:"extra_groups" gorm:"extra_groups" json:"extra_groups,omitempty"`
	FailedJobs            []RestaurantFailed_Jobs             `db:"failed_jobs" gorm:"failed_jobs" json:"failed_jobs,omitempty"`
	Jobs                  []RestaurantJobs                    `db:"jobs" gorm:"jobs" json:"jobs,omitempty"`
	Migrations            []RestaurantMigrations              `db:"migrations" gorm:"migrations" json:"migrations,omitempty"`
	Orders                []RestaurantOrders                  `db:"orders" gorm:"orders" json:"orders,omitempty"`
	OrderActions          []RestaurantOrder_Actions           `db:"order_actions" gorm:"order_actions" json:"order_actions,omitempty"`
	OrderProducts         []RestaurantOrder_Products          `db:"order_products" gorm:"order_products" json:"order_products,omitempty"`
	PasswordResets        []RestaurantPassword_Resets         `db:"password_resets" gorm:"password_resets" json:"password_resets,omitempty"`
	PersonalAccessTokens  []RestaurantPersonal_Access_Tokens  `db:"personal_access_tokens" gorm:"personal_access_tokens" json:"personal_access_tokens,omitempty"`
	PhoneVerifications    []RestaurantPhone_Verifications     `db:"phone_verifications" gorm:"phone_verifications" json:"phone_verifications,omitempty"`
	PointActions          []RestaurantPoint_Actions           `db:"point_actions" gorm:"point_actions" json:"point_actions,omitempty"`
	Products              []RestaurantProducts                `db:"products" gorm:"products" json:"products,omitempty"`
	ProductExtraGroups    []RestaurantProduct_Extra_Groups    `db:"product_extra_groups" gorm:"product_extra_groups" json:"product_extra_groups,omitempty"`
	ProductVariants       []RestaurantProduct_Variants        `db:"product_variants" gorm:"product_variants" json:"product_variants,omitempty"`
	ProductVariantOptions []RestaurantProduct_Variant_Options `db:"product_variant_options" gorm:"product_variant_options" json:"product_variant_options,omitempty"`
	Reservations          []RestaurantReservations            `db:"reservations" gorm:"reservations" json:"reservations,omitempty"`
	Settings              []RestaurantSettings                `db:"settings" gorm:"settings" json:"settings,omitempty"`
}

type RestaurantSettingsByName struct {
	Name                                      string `db:"name" gorm:"name" json:"name"`
	Logo                                      string `db:"logo" gorm:"logo" json:"logo"`
	Title                                     string `db:"title" gorm:"title" json:"title"`
	Favicon                                   string `db:"favicon" gorm:"favicon" json:"favicon"`
	ActiveDays                                string `db:"activeDays" gorm:"activeDays" json:"activeDays"`
	OrderStatus                               string `db:"order_status" gorm:"order_status" json:"order_status"`
	MinOrderAmount                            string `db:"min_order_amount" gorm:"min_order_amount" json:"min_order_amount"`
	DeliveryTime                              string `db:"delivery_time" gorm:"delivery_time" json:"delivery_time"`
	RestaurantLocationLat                     string `db:"restaurant_location_lat" gorm:"restaurant_location_lat" json:"restaurant_location_lat"`
	RestaurantLocationLong                    string `db:"restaurant_location_long" gorm:"restaurant_location_long" json:"restaurant_location_long"`
	Description                               string `db:"description" gorm:"description" json:"description"`
	Banner                                    string `db:"banner" gorm:"banner" json:"banner"`
	Address                                   string `db:"address" gorm:"address" json:"address"`
	DeliveryPrice                             string `db:"delivery_price" gorm:"delivery_price" json:"delivery_price"`
	FreeDeliveryLimit                         string `db:"free_delivery_limit" gorm:"free_delivery_limit" json:"free_delivery_limit"`
	OrderBonus                                string `db:"order_bonus" gorm:"order_bonus" json:"order_bonus"`
	CurrierType                               string `db:"currier_type" gorm:"currier_type" json:"currier_type"`
	TemplateModuleSearchTitle                 string `db:"template_module_search_title" gorm:"template_module_search_title" json:"template_module_search_title"`
	TemplateModuleSearchDetail                string `db:"template_module_search_detail" gorm:"template_module_search_detail" json:"template_module_search_detail"`
	TemplateModuleSearchInputText             string `db:"template_module_search_input_text" gorm:"template_module_search_input_text" json:"template_module_search_input_text"`
	TemplateModuleSearchImage                 string `db:"template_module_search_image" gorm:"template_module_search_image" json:"template_module_search_image"`
	TemplateModuleSearchIsActive              string `db:"template_module_search_is_active" gorm:"template_module_search_is_active" json:"template_module_search_is_active"`
	TemplateModuleSliderIsActive              string `db:"template_module_slider_is_active" gorm:"template_module_slider_is_active" json:"template_module_slider_is_active"`
	TemplateModuleSliderImage                 string `db:"template_module_slider_image" gorm:"template_module_slider_image" json:"template_module_slider_image"`
	TemplateModuleSliderButtonText            string `db:"template_module_slider_button_text" gorm:"template_module_slider_button_text" json:"template_module_slider_button_text"`
	TemplateModuleSliderTitle                 string `db:"template_module_slider_title" gorm:"template_module_slider_title" json:"template_module_slider_title"`
	TemplateModuleSliderButtonUrl             string `db:"template_module_slider_button_url" gorm:"template_module_slider_button_url" json:"template_module_slider_button_url"`
	TemplateModuleCategoryLength              string `db:"template_module_category_length" gorm:"template_module_category_length" json:"template_module_category_length"`
	TemplateModuleCategoryActiveColor         string `db:"template_module_category_active_color" gorm:"template_module_category_active_color" json:"template_module_category_active_color"`
	TemplateModuleCategoryIsActive            string `db:"template_module_category_is_active" gorm:"template_module_category_is_active" json:"template_module_category_is_active"`
	TemplateModuleProductsIsActive            string `db:"template_module_products_is_active" gorm:"template_module_products_is_active" json:"template_module_products_is_active"`
	TemplateModuleProductLength               string `db:"template_module_product_length" gorm:"template_module_product_length" json:"template_module_product_length"`
	TemplateModuleProductsShowSalePrice       string `db:"template_module_products_show_sale_price" gorm:"template_module_products_show_sale_price" json:"template_module_products_show_sale_price"`
	TemplateModuleProductsShowTime            string `db:"template_module_products_show_time" gorm:"template_module_products_show_time" json:"template_module_products_show_time"`
	TemplateModuleSliderShowNoDiscountPrice   string `db:"template_module_slider_show_no_discount_price" gorm:"template_module_slider_show_no_discount_price" json:"template_module_slider_show_no_discount_price"`
	TemplateModuleProductsShowCommentCount    string `db:"template_module_products_show_comment_count" gorm:"template_module_products_show_comment_count" json:"template_module_products_show_comment_count"`
	TemplateModuleProductsShowNoDiscountPrice string `db:"template_module_products_show_no_discount_price" gorm:"template_module_products_show_no_discount_price" json:"template_module_products_show_no_discount_price"`
	TemplateModuleProductsShowType            string `db:"template_module_products_show_type" gorm:"template_module_products_show_type" json:"template_module_products_show_type"`
	TemplateModuleCampaignIsActive            string `db:"template_module_campaign_is_active" gorm:"template_module_campaign_is_active" json:"template_module_campaign_is_active"`
	TemplateModuleFeaturedProductsIsActive    string `db:"template_module_featured_products_is_active" gorm:"template_module_featured_products_is_active" json:"template_module_featured_products_is_active"`
	TemplateModuleFeaturedProductsLength      string `db:"template_module_featured_products_length" gorm:"template_module_featured_products_length" json:"template_module_featured_products_length"`
	TemplateModuleFeaturedProductsTitle       string `db:"template_module_featured_products_title" gorm:"template_module_featured_products_title" json:"template_module_featured_products_title"`
	TemplateModuleCampaignTitle               string `db:"template_module_campaign_title" gorm:"template_module_campaign_title" json:"template_module_campaign_title"`
	TemplateModuleGeneralBackground           string `db:"template_module_general_background" gorm:"template_module_general_background" json:"template_module_general_background"`
	TemplateModuleGeneralHeaderBackground     string `db:"template_module_general_header_background" gorm:"template_module_general_header_background" json:"template_module_general_header_background"`
	TemplateModuleProductsBoxBackground       string `db:"template_module_products_box_background" gorm:"template_module_products_box_background" json:"template_module_products_box_background"`
	TemplateModuleSearchInputBackground       string `db:"template_module_search_input_background" gorm:"template_module_search_input_background" json:"template_module_search_input_background"`
	TemplateModuleSearchInputTextColor        string `db:"template_module_search_input_text_color" gorm:"template_module_search_input_text_color" json:"template_module_search_input_text_color"`
	TemplateModuleFooterCampaignIsActive      string `db:"template_module_footer_campaign_is_active" gorm:"template_module_footer_campaign_is_active" json:"template_module_footer_campaign_is_active"`
	TemplateModuleFooterCampaignTitle         string `db:"template_module_footer_campaign_title" gorm:"template_module_footer_campaign_title" json:"template_module_footer_campaign_title"`
	TemplateModuleFooterCampaignButtonUrl     string `db:"template_module_footer_campaign_button_url" gorm:"template_module_footer_campaign_button_url" json:"template_module_footer_campaign_button_url"`
	TemplateModuleFooterCampaignButtonText    string `db:"template_module_footer_campaign_button_text" gorm:"template_module_footer_campaign_button_text" json:"template_module_footer_campaign_button_text"`
	TemplateModuleFooterCampaignImage         string `db:"template_module_footer_campaign_image" gorm:"template_module_footer_campaign_image" json:"template_module_footer_campaign_image"`
	TemplateModuleFooterCampaignDetail        string `db:"template_module_footer_campaign_detail" gorm:"template_module_footer_campaign_detail" json:"template_module_footer_campaign_detail"`
	Phone                                     string `db:"phone" gorm:"phone" json:"phone"`
	DeliveryArea                              string `db:"delivery_area" gorm:"delivery_area" json:"delivery_area"`
	PaidDeliveryStart                         string `db:"paid_delivery_start" gorm:"paid_delivery_start" json:"paid_delivery_start"`
	KmDeliveryPrice                           string `db:"km_delivery_price" gorm:"km_delivery_price" json:"km_delivery_price"`
	DefaultProductImage                       string `db:"default_product_image" gorm:"default_product_image" json:"default_product_image"`
	TimeMondayOpen                            string `db:"time_Monday_open" gorm:"time_Monday_open" json:"time_Monday_open"`
	TimeMondayDelivery                        string `db:"time_Monday_delivery" gorm:"time_Monday_delivery" json:"time_Monday_delivery"`
	TimeMondayPickup                          string `db:"time_Monday_pickup" gorm:"time_Monday_pickup" json:"time_Monday_pickup"`
	TimeTuesdayOpen                           string `db:"time_Tuesday_open" gorm:"time_Tuesday_open" json:"time_Tuesday_open"`
	TimeTuesdayDelivery                       string `db:"time_Tuesday_delivery" gorm:"time_Tuesday_delivery" json:"time_Tuesday_delivery"`
	TimeTuesdayPickup                         string `db:"time_Tuesday_pickup" gorm:"time_Tuesday_pickup" json:"time_Tuesday_pickup"`
	TimeWednesdayOpen                         string `db:"time_Wednesday_open" gorm:"time_Wednesday_open" json:"time_Wednesday_open"`
	TimeWednesdayDelivery                     string `db:"time_Wednesday_delivery" gorm:"time_Wednesday_delivery" json:"time_Wednesday_delivery"`
	TimeWednesdayPickup                       string `db:"time_Wednesday_pickup" gorm:"time_Wednesday_pickup" json:"time_Wednesday_pickup"`
	TimeThursdayOpen                          string `db:"time_Thursday_open" gorm:"time_Thursday_open" json:"time_Thursday_open"`
	TimeThursdayDelivery                      string `db:"time_Thursday_delivery" gorm:"time_Thursday_delivery" json:"time_Thursday_delivery"`
	TimeThursdayPickup                        string `db:"time_Thursday_pickup" gorm:"time_Thursday_pickup" json:"time_Thursday_pickup"`
	TimeFridayOpen                            string `db:"time_Friday_open" gorm:"time_Friday_open" json:"time_Friday_open"`
	TimeFridayDelivery                        string `db:"time_Friday_delivery" gorm:"time_Friday_delivery" json:"time_Friday_delivery"`
	TimeFridayPickup                          string `db:"time_Friday_pickup" gorm:"time_Friday_pickup" json:"time_Friday_pickup"`
	TimeSaturdayOpen                          string `db:"time_Saturday_open" gorm:"time_Saturday_open" json:"time_Saturday_open"`
	TimeSaturdayDelivery                      string `db:"time_Saturday_delivery" gorm:"time_Saturday_delivery" json:"time_Saturday_delivery"`
	TimeSaturdayPickup                        string `db:"time_Saturday_pickup" gorm:"time_Saturday_pickup" json:"time_Saturday_pickup"`
	TimeSundayOpen                            string `db:"time_Sunday_open" gorm:"time_Sunday_open" json:"time_Sunday_open"`
	TimeSundayDelivery                        string `db:"time_Sunday_delivery" gorm:"time_Sunday_delivery" json:"time_Sunday_delivery"`
	TimeSundayPickup                          string `db:"time_Sunday_pickup" gorm:"time_Sunday_pickup" json:"time_Sunday_pickup"`
	MaxOrderDay                               string `db:"max_order_day" gorm:"max_order_day" json:"max_order_day"`
	OrderDeliveryTime                         string `db:"order_delivery_time" gorm:"order_delivery_time" json:"order_delivery_time"`
	OrderPickupTime                           string `db:"order_pickup_time" gorm:"order_pickup_time" json:"order_pickup_time"`
	OrderRestaurantTime                       string `db:"order_restaurant_time" gorm:"order_restaurant_time" json:"order_restaurant_time"`
	Theme                                     string `db:"theme" gorm:"theme" json:"theme"`
	Email                                     string `db:"email" gorm:"email" json:"email"`
	HealthReportUrl                           string `db:"health_report_url" gorm:"health_report_url" json:"health_report_url"`
	TraditionalBackgroundImage                string `db:"traditional_background_image" gorm:"traditional_background_image" json:"traditional_background_image"`
	Facebook                                  string `db:"facebook" gorm:"facebook" json:"facebook"`
	Twitter                                   string `db:"twitter" gorm:"twitter" json:"twitter"`
	Instagram                                 string `db:"instagram" gorm:"instagram" json:"instagram"`
	HeaderTags                                string `db:"header_tags" gorm:"header_tags" json:"header_tags"`
	FooterTags                                string `db:"footer_tags" gorm:"footer_tags" json:"footer_tags"`
	Paymentfee                                string `db:"paymentfee" gorm:"paymentfee" json:"paymentfee"`
	BagFee                                    string `db:"bagfee" gorm:"bagfee" json:"bagfee"`
	PaymentfeeOther                           string `db:"paymentfee_other" gorm:"paymentfee_other" json:"paymentfee_other"`
	PaymentfeeDelivery                        string `db:"paymentfee_delivery" gorm:"paymentfee_delivery" json:"paymentfee_delivery"`
	PaymentfeeTable                           string `db:"paymentfee_table" gorm:"paymentfee_table" json:"paymentfee_table"`
	OrderDeliveryPickupTime                   string `db:"order_delivery_pickup_time" gorm:"order_delivery_pickup_time" json:"order_delivery_pickup_time"`
	PopupMessage                              string `db:"popup_message" gorm:"popup_message" json:"popup_message"`
}

type RestaurantOpenTimes struct {
	Day      string `json:"day"`
	IsOpen   bool   `json:"is_open"`
	Open     string `json:"open"`
	Delivery string `json:"delivery"`
	Pickup   string `json:"pickup"`
}
