package dbcache

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/karolpiernikarz/automanage/cache"
	"github.com/karolpiernikarz/automanage/helpers"
	"github.com/karolpiernikarz/automanage/models"
)

func Companies() (companies []models.Companies, err error) {
	// check if key exist
	var value string
	if value, err = cache.GetValueFromKey("dbCache_companies"); err == nil {
		// if key exist, return value
		err = json.Unmarshal([]byte(value), &companies)
		return
	}
	// if key does not exist, create the key
	if errors.Is(badger.ErrKeyNotFound, err) {
		companies = helpers.GetCompanies()
		// marshal to json
		var companiesJson []byte
		companiesJson, _ = json.Marshal(companies)
		err = cache.SetKeyValue([]byte("dbCache_companies"), companiesJson, 5*time.Minute)
		if err != nil {
			return
		}
	}
	return
}

func Restaurants() (restaurants []models.Restaurant, err error) {
	// check if key exist
	var value string
	if value, err = cache.GetValueFromKey("dbCache_restaurants"); err == nil {
		// if key exist, return value
		err = json.Unmarshal([]byte(value), &restaurants)
		return
	}
	// if key does not exist, create the key
	if errors.Is(badger.ErrKeyNotFound, err) {
		restaurants = helpers.GetAllRestaurants()
		// marshal to json
		var restaurantsJson []byte
		restaurantsJson, _ = json.Marshal(restaurants)
		err = cache.SetKeyValue([]byte("dbCache_restaurants"), restaurantsJson, 5*time.Minute)
		if err != nil {
			return
		}
	}
	return
}
