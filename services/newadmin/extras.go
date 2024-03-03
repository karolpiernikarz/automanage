package newAdmin

import "time"

type ExtraGroup struct {
	Id          int       `json:"id"`
	MenuId      int       `json:"menu_id" gorm:"type:int(11)"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	DisplayName string    `json:"display_name" gorm:"type:varchar(255)"`
	Limit       int       `json:"limit" gorm:"type:int(11)"`
	Sort        int       `json:"sort" gorm:"type:int(11)"`
	IsActive    int       `json:"is_active" gorm:"type:tinyint(1)"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp"`
}

func (ExtraGroup) TableName() string {
	return "extra_groups"
}

func (ExtraGroup) PrimaryKey() string {
	return "id"
}

func (ExtraGroup *ExtraGroup) Create() error {
	return DB.Create(ExtraGroup).Error
}

type Extra struct {
	Id         int       `json:"id"`
	GroupId    int       `json:"group_id" gorm:"type:int(11)"`
	Name       string    `json:"name" gorm:"type:varchar(255)"`
	Price      float64   `json:"price" gorm:"type:double"`
	Sort       int       `json:"sort" gorm:"type:int(11)"`
	IsDisabled int       `json:"is_disabled" gorm:"type:tinyint(1)"`
	IsDefault  int       `json:"is_default" gorm:"type:tinyint(1)"`
	Limit      int       `json:"limit" gorm:"type:int(11)"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp"`
}

func (Extra) TableName() string {
	return "extras"
}

func (Extra) PrimaryKey() string {
	return "id"
}

func (Extra *Extra) Create() error {
	return DB.Create(Extra).Error
}
