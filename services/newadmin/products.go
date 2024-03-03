package newAdmin

import "time"

type Product struct {
	Id                   int       `json:"id"`
	MenuId               int       `json:"menu_id" gorm:"type:int(11)"`
	CategoryId           int       `json:"category_id" gorm:"type:int(11)"`
	Name                 string    `json:"name" gorm:"type:varchar(255)"`
	Slug                 string    `json:"slug" gorm:"type:varchar(255)"`
	Image                string    `json:"image" gorm:"type:varchar(255)"`
	WithoutDiscountPrice float64   `json:"without_discount_price" gorm:"type:double"`
	Price                float64   `json:"price" gorm:"type:double"`
	Description          string    `json:"description" gorm:"type:longtext"`
	Sort                 int       `json:"sort" gorm:"type:int(11)"`
	IsActive             int       `json:"is_active" gorm:"type:tinyint(1)"`
	Keywords             string    `json:"keywords" gorm:"type:longtext"`
	Ingredients          string    `json:"ingredients" gorm:"type:longtext"`
	Discount             JSONB     `json:"discount" gorm:"type:longtext"`
	Type                 string    `json:"type" gorm:"type:longtext"`
	Unit                 JSONB     `json:"unit" gorm:"type:longtext"`
	Allergen             string    `json:"allergen" gorm:"type:longtext"`
	CreatedAt            time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt            time.Time `json:"updated_at" gorm:"type:timestamp"`
}

func (Product) TableName() string {
	return "products"
}

func (Product) PrimaryKey() string {
	return "id"
}

func (Product *Product) Create() error {
	return DB.Create(Product).Error
}

type ProductExtraGroups struct {
	Id                     int       `json:"id" gorm:"primary_key"`
	ProductId              int       `json:"product_id" gorm:"type:int(11)"`
	ProductVariantOptionId int       `json:"product_variant_option_id" gorm:"type:int(11)"`
	ExtraGroupId           int       `json:"extra_group_id" gorm:"type:int(11)"`
	Sort                   int       `json:"sort" gorm:"type:int(11)"`
	CreatedAt              time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt              time.Time `json:"updated_at" gorm:"type:timestamp"`
}

func (ProductExtraGroups) TableName() string {
	return "product_extra_groups"
}

func (ProductExtraGroups) PrimaryKey() string {
	return "id"
}

func (ProductExtraGroups *ProductExtraGroups) Create() error {
	return DB.Create(ProductExtraGroups).Error
}
