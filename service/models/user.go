package models

type User struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `json:"name" binding:"required,lte=30"`
	Key    string `gorm:"unique" json:"key" binding:"required,lte=20"`
	Secret string `json:"secret" binding:"required,lte=10"`
}

type UserOne struct {
	Key  string `json:"key"  binding:"required,lte=20"`
	Sign string `json:"sign"  binding:"required"`
}
