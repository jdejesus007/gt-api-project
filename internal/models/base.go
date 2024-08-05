package models

import (
	"time"
)

type BaseModel struct {
	ID         uint           `gorm:"primaryKey" json:"-"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  *time.Time     `gorm:"index" json:"-"`
	BaseStatus BaseStatusEnum `gorm:"index" default:"0" json:"base_status"`
}
