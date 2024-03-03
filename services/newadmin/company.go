package newAdmin

import (
	"time"
)

type Company struct {
	Id        int       `json:"id,omitempty" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	ChainName string    `json:"chain_name" gorm:"type:varchar(255)"`
	Notes     string    `json:"notes" gorm:"type:longtext"`
	TaxNumber string    `json:"tax_number" gorm:"type:varchar(255)"`
	Contact   JSONB     `json:"contact" gorm:"type:longtext"`
	Domain    string    `json:"domain" gorm:"type:varchar(255)"`
	Social    JSONB     `json:"social" gorm:"type:longtext"`
	Settings  JSONB     `json:"settings" gorm:"type:longtext"`
	IsActive  int       `json:"is_active,omitempty" gorm:"type:tinyint(1)"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
	DeleteAt  time.Time `json:"delete_at,omitempty" gorm:"type:timestamp"`
}

func (Company) TableName() string {
	return "companies"
}

func (Company) PrimaryKey() string {
	return "id"
}

func (company *Company) Create() error {
	return DB.Create(company).Error
}
