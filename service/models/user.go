package models

type User struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `json:"name"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type UserOne struct {
	Key  string `json:"key"  binding:"required,lte=20"`
	Sign string `json:"sign"  binding:"required"`
}
