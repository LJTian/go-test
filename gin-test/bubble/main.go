package main

import (
	"gin_test/bubble/dao"
	"gin_test/bubble/models"
	"gin_test/bubble/routers"
)

func main() {
	// 初始化数据库
	if err := dao.InitMysqlDB(); err != nil {
		panic(err)
	}
	// 迁移
	dao.Db.AutoMigrate(&models.Todo{})

	//启动Gin
	r := routers.StartRouter()
	r.Run(":8005")
}
