package models

type Book struct {
	BaseModel
	UUID string `gorm:"size:255;index:,unique" json:"uuid"`
	Name string `gorm:"size:255" json:"name"`
}
