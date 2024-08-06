package models

import (
	"database/sql/driver"
	"strings"
	"time"
)

type BaseModel struct {
	ID         uint           `gorm:"primaryKey" json:"-"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  *time.Time     `gorm:"index" json:"-"`
	BaseStatus BaseStatusEnum `gorm:"index" default:"0" json:"-"`
}

type StringSliceType []string

func (o *StringSliceType) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		*o = strings.Split(src.(string), ",")
		return nil
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
