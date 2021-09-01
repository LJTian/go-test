package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitMysqlDB() (err error) {
	// 初始化数据库
	sTmpBuf := "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(sTmpBuf), &gorm.Config{})
	if err != nil {
		return err
	}
	return
}
