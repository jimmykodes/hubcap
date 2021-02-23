package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

var (
	ErrInvalidDataType = errors.New("invalid type for Service Data field")
)

type ServiceData map[string]interface{}

func (s *ServiceData) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		return json.Unmarshal([]byte(src.(string)), s)
	default:
		return ErrInvalidDataType
	}
}

func (s *ServiceData) Value() (driver.Value, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return driver.Value(string(data)), nil
}
