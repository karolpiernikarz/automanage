package newAdmin

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONB map[string]interface{}

// Value Marshal
func (jsonField JSONB) Value() (driver.Value, error) {
	return json.Marshal(jsonField)
}

// Scan Unmarshal
func (jsonField *JSONB) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(data, &jsonField)
}

type JSONBArray []interface{}

// Value Marshal
func (jsonField JSONBArray) Value() (driver.Value, error) {
	return json.Marshal(jsonField)
}

// Scan Unmarshal
func (jsonField *JSONBArray) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(data, &jsonField)
}

type JSONBStringArray []string

// Value Marshal
func (jsonField JSONBStringArray) Value() (driver.Value, error) {
	return json.Marshal(jsonField)
}

// Scan Unmarshal
func (jsonField *JSONBStringArray) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(data, &jsonField)
}
