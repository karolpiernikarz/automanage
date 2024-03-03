package jobs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/karolpiernikarz/automanage/cache"
	"github.com/karolpiernikarz/automanage/helpers"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/karolpiernikarz/automanage/utils"
)

func CreateOrderBoxCache() (err error) {
	restaurants := helpers.GetAllRestaurants()
	var orderboxIds []string
	var i uint
	for _, s := range restaurants {
		i = 0

		if s.Info.Data().TerminalId == nil {
			continue
		}
		for utils.StringInSlice(fmt.Sprint(s.Info.Data().TerminalId)+"_"+fmt.Sprint(i), orderboxIds) {
			i++
		}
		orderboxIds = append(orderboxIds, fmt.Sprint(s.Info.Data().TerminalId)+"_"+fmt.Sprint(i))

		orderboxInfo := models.OrderboxInfo{}
		orderboxInfo.ID = s.ID
		orderboxInfo.TerminalPassword = s.Info.Data().TerminalPassword
		orderboxInfo.TerminalType = s.Info.Data().TerminalType
		orderboxInfo.TerminalUsername = s.Info.Data().TerminalUsername
		orderboxJson, err := json.Marshal(orderboxInfo)
		if err != nil {
			continue
		}
		key := []byte("orderbox_" + fmt.Sprint(s.Info.Data().TerminalId) + "_" + fmt.Sprint(i))
		// if orderboxInfo.ID already exist in orderboxIds, then change add key name to _+1
		err = cache.SetKeyValue(key, orderboxJson, 10*time.Minute)
		if err != nil {
			return err
		}
	}
	return err
}
