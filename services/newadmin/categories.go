package newAdmin

import "time"

type Category struct {
	Id          int       `json:"id" gorm:"primary_key"`
	MenuId      int       `json:"menu_id" gorm:"type:int(11)"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:longtext"`
	Slug        string    `json:"slug" gorm:"type:varchar(255)"`
	Icon        string    `json:"icon" gorm:"type:varchar(255)"`
	Banner      string    `json:"banner" gorm:"type:varchar(255)"`
	Sort        int       `json:"sort" gorm:"type:int(11)"`
	IsActive    int       `json:"is_active" gorm:"type:tinyint(1)"`
	Discount    int       `json:"discount" gorm:"type:int(11)"`
	Hours       JSONB     `json:"hours" gorm:"type:longtext"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp"`
}

func (Category) TableName() string {
	return "categories"
}

func (Category) PrimaryKey() string {
	return "id"
}

func (Category *Category) Create() error {
	return DB.Create(Category).Error
}
