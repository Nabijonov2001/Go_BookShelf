package models

type User struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `json:"name" binding:"required,gte=3,lte=30"`
	Key    string `gorm:"unique" json:"key" binding:"required,gte=3,lte=20"`
	Secret string `json:"secret" binding:"required,gte=4,lte=10"`
}

type UserOne struct {
	Key  string `json:"key"  binding:"required,gte=3,lte=20"`
	Sign string `json:"sign"  binding:"required"`
}
