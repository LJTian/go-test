package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Commodity struct {
	gorm.Model
	Name       string
	CommodCode string
}

type User struct {
	gorm.Model
	Name   string
	Age    uint8
	IcCard string
}

type Order struct {
	gorm.Model
	Name      string
	UserId    string
	Commodity Commodity `gorm:"foreignKey:ID;references:ID"`
}

// 钩子函数
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	fmt.Println("hook func~")
	if u.Name == "123" {
		u.Name = "name" + "123"
	}
	return
}
