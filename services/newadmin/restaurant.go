package newAdmin

import (
	"database/sql"
	"time"
)

type Restaurant struct {
	Id               int          `json:"id" gorm:"type:int(11)"`
	CompanyId        int          `json:"company_id" gorm:"type:int(11)"`
	MenuId           int          `json:"menu_id" gorm:"type:int(11)"`
	Name             string       `json:"name" gorm:"type:varchar(255)"`
	Slug             string       `json:"slug" gorm:"type:varchar(255)"`
	Logo             string       `json:"logo" gorm:"type:varchar(255)"`
	Address          JSONB        `json:"address" gorm:"type:longtext"`
	Phone            string       `json:"phone" gorm:"type:varchar(255)"`
	Commission       JSONB        `json:"commission" gorm:"type:longtext"`
	Bank             JSONB        `json:"bank" gorm:"type:longtext"`
	PlaceId          string       `json:"place_id" gorm:"type:varchar(255)"`
	Lat              string       `json:"lat" gorm:"type:varchar(255)"`
	Long             string       `json:"long" gorm:"type:varchar(255)"`
	IsActive         int          `json:"is_active" gorm:"type:tinyint(1)"`
	PlatformIsActive int          `json:"platform_is_active" gorm:"type:tinyint(1)"`
	WebIsActive      int          `json:"web_is_active" gorm:"type:tinyint(1)"`
	CreatedAt        time.Time    `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt        time.Time    `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt        sql.NullTime `json:"deleted_at" gorm:"type:timestamp"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (Restaurant) PrimaryKey() string {
	return "id"
}

func (restaurant *Restaurant) Create() error {
	return DB.Create(restaurant).Error
}
