package models

type Order struct {
	BaseModel
	UUID         string `gorm:"size:255;index,unique" json:"UUID"`
	CustomerUUID string `json:"size:255;" json:"customerUUID"`
}
