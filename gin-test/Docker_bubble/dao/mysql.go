package dao

import (
	"bubble/conf"
	"bubble/utils/db"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitMysqlDB() (err error) {
	// 初始化数据库
	cfg := conf.CFG
	mysqlDbInfo := db.NewMySqlDBInfo(cfg.Ip, cfg.Port, cfg.UserName, cfg.PassWord, cfg.Name, "utf8mb4", true)
	Db, err = mysqlDbInfo.InitMysql()
	if err != nil {
		return err
	}
	return
}
