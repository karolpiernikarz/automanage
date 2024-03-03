package newAdmin

import (
	"database/sql"
	"time"
)

type Order struct {
	Id           int          `json:"id"`
	RestaurantId int          `json:"restaurant_id" gorm:"type:int(11)"`
	OrderNumber  string       `json:"order_number" gorm:"type:varchar(255)"`
	CustomerId   int          `json:"customer_id" gorm:"type:int(11)"`
	Customer     JSONB        `gorm:"type:longtext"`
	Address      JSONB        `gorm:"type:longtext"`
	Prices       JSONB        `gorm:"type:longtext"`
	Payment      JSONB        `gorm:"type:longtext"`
	PaymentId    string       `json:"payment_id" gorm:"type:varchar(255)"`
	PaymentType  string       `json:"payment_type" gorm:"type:varchar(255)"`
	Status       int          `json:"status" gorm:"type:int(11)"`
	Type         int          `json:"type" gorm:"type:int(11)"`
	Note         string       `json:"note"`
	Date         time.Time    `json:"date" gorm:"type:timestamp"`
	Currier      JSONBArray   `gorm:"type:longtext"`
	PickupDate   sql.NullTime `json:"pickup_date" gorm:"type:timestamp,omitempty" db:"pickup_date,omitempty"`
	DeliveryDate sql.NullTime `json:"delivery_date" gorm:"type:timestamp,omitempty" db:"delivery_date,omitempty"`
	IsPreOrder   int          `json:"is_pre_order" gorm:"type:tinyint(1)"`
	CreatedAt    time.Time    `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt    time.Time    `json:"updated_at" gorm:"type:timestamp"`
}

func (Order) TableName() string {
	return "orders"
}

func (Order) PrimaryKey() string {
	return "id"
}

func (Order *Order) Create() error {
	return DB.Create(Order).Error
}

type OrderProduct struct {
	Id          int        `json:"id" gorm:"primary_key"`
	OrderId     int        `json:"order_id" gorm:"type:int(11)"`
	ProductId   int        `json:"product_id" gorm:"type:int(11)"`
	Qty         int        `json:"qty" gorm:"type:int(11)"`
	UnitPrice   float64    `json:"unit_price" gorm:"type:double"`
	Price       float64    `json:"price" gorm:"type:double"`
	Variants    JSONBArray `json:"variants" gorm:"type:longtext"`
	Extras      JSONBArray `json:"extras" gorm:"type:longtext"`
	Ingredients JSONB      `json:"ingredients" gorm:"type:longtext"`
	Note        string     `json:"note" gorm:"type:longtext"`
	CreatedAt   time.Time  `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"type:timestamp"`
}

func (OrderProduct) TableName() string {
	return "order_products"
}

func (OrderProduct) PrimaryKey() string {
	return "id"
}

func (OrderProduct *OrderProduct) Create() error {
	return DB.Create(OrderProduct).Error
}

func (Order *Order) Get(id int) error {
	return DB.Where("id = ?", id).First(Order).Error
}
