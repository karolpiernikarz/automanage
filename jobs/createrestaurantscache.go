package jobs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/karolpiernikarz/automanage/cache"
	"github.com/karolpiernikarz/automanage/helpers"
)

func CreateRestaurantsCache() {
	restaurants := helpers.GetAllRestaurants()
	orderboxJson, _ := json.Marshal(restaurants)
	err := cache.SetKeyValue([]byte("restaurants"), orderboxJson, 10*time.Minute)
	if err != nil {
		fmt.Println("error, restaurant cache setkeyvalue:", err)
		return
	}
}
