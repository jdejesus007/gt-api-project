package models

type Order struct {
	BaseModel
	UUID         string          `gorm:"size:255;index,unique" json:"UUID"`
	CustomerUUID string          `gorm:"size:255;" json:"customerUUID"`
	BookUUIDs    StringSliceType `gorm:"type:VARCHAR(255)" json:"bookUUIDs"`
}
