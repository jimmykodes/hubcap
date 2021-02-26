package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrInvalidDataType = errors.New("invalid type for Service Data field")
)

type ServiceData map[string]interface{}

func (s *ServiceData) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	switch t := src.(type) {
	case []byte:
		return json.Unmarshal(t, s)
	default:
		return fmt.Errorf("%w: %T", ErrInvalidDataType, t)
	}
}

func (s ServiceData) Value() (driver.Value, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return driver.Value(string(data)), nil
}
