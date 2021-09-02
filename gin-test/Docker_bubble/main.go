package main

import (
	"bubble/conf"
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"bubble/utils/tools"
	"fmt"
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
