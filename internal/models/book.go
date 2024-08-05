package models

type Book struct {
	BaseModel
	UUID string `gorm:"size:255;index:,unique" json:"UUID"`
	Name string `gorm:"size:255" json:"name"`
}
