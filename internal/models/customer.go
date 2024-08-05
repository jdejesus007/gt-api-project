package models

type Customer struct {
	BaseModel
	UUID      string `gorm:"size:255;index:,unique" json:"UUID"`
	Email     string `gorm:"size:255;index:,unique" json:"email"`
	FirstName string `gorm:"size:255" json:"firstName"`
	LastName  string `gorm:"size:255" json:"lastName"`
}
