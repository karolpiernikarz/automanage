package newAdmin

import "time"

type RestaurantSetting struct {
	Id           int       `json:"id"`
	RestaurantId int       `json:"restaurant_id"`
	Name         string    `json:"name"`
	Value        string    `json:"value"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp"`
}

func (RestaurantSetting) TableName() string {
	return "restaurant_settings"
}

func (RestaurantSetting) PrimaryKey() string {
	return "id"
}

func (restaurantSetting *RestaurantSetting) Create() error {
	if DB.Model(&restaurantSetting).Where("name = ?", restaurantSetting.Name).Updates(&restaurantSetting).RowsAffected == 0 {
		return DB.Create(&restaurantSetting).Error
	} else {
		return nil
	}
}
