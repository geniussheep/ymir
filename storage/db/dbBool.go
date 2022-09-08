package db

import (
	"database/sql/driver"
	"errors"
)

type DbBool bool

func (b DbBool) Value() (driver.Value, error) {
	result := make([]byte, 1)
	if b {
		result[0] = byte(1)
	} else {
		result[0] = 0
	}
	return result, nil
}
func (b *DbBool) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	*b = v[0] == 1

	return nil
}
