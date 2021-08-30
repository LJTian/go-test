package model

import (
	"gorm.io/gorm"
)

type ApiRecoid struct {
	gorm.Model
	ReqUrl string
	Timing float64
	Status bool
}
