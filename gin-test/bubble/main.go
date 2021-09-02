package main

import (
	"fmt"
	"gin_test/bubble/conf"
	"gin_test/bubble/dao"
	"gin_test/bubble/models"
	"gin_test/bubble/routers"
	"gin_test/bubble/utils/tools"
)

func main() {
	// 读取配置文件
	conf := conf.NewConf(tools.GetCurrentDirectory() + "/conf/cfg.ini")
	conf.InitConf(tools.GetCurrentDirectory() + "/conf/cfg.ini")
	fmt.Printf("%s\n", conf.PrintConf())
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
