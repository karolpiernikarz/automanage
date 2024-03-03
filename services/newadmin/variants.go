package newAdmin

import "time"

type Variant struct {
	Id        int       `json:"id"`
	ProductId int       `json:"product_id" gorm:"type:int(11)"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}

func (Variant) TableName() string {
	return "variants"
}

func (Variant) PrimaryKey() string {
	return "id"
}

func (Variant *Variant) Create() error {
	return DB.Create(Variant).Error
}

type VariantOption struct {
	Id        int       `json:"id"`
	VariantId int       `json:"variant_id" gorm:"type:int(11)"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Price     float64   `json:"price" gorm:"type:double"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}

func (VariantOption) TableName() string {
	return "variant_options"
}

func (VariantOption) PrimaryKey() string {
	return "id"
}

func (VariantOption *VariantOption) Create() error {
	return DB.Create(VariantOption).Error
}
