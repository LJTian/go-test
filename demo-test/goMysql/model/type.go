package model

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Code struct {
	X, Y int
}

func (c *Code) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	return nil
}

func (c Code) GormDataType() string {
	return "Code"
}

func (c Code) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "round(RAND()*?%?/1)",
		Vars: []interface{}{c.X, c.Y},
	}
}
