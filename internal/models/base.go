package models

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

type BaseModel struct {
	ID         uint           `gorm:"primaryKey" json:"-"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  *time.Time     `gorm:"index" json:"-"`
	BaseStatus BaseStatusEnum `gorm:"index" default:"0" json:"-"`
}

type StringSliceType []string

func (o *StringSliceType) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}
	*o = strings.Split(string(bytes), ",")
	return nil
}

func (o StringSliceType) Value() (driver.Value, error) {
	if len(o) == 0 {
		return nil, nil
	}
	return strings.Join(o, ","), nil
}
