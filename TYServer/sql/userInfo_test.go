package sql

import (
	"TYServer/utils/db"
	"fmt"
	"testing"
)

func TestAddUsertoDB(t *testing.T) {

	fmt.Println("开始测试~")
	newDB := db.NewDBinfo("81.70.17.60", "3306", "TYServer", "root", "123456")
	newDB.InitDB()
	fmt.Println("连接数据库成功！")
	userinfo := NewUserInfo("田利军测试1", "17613811921", 1)
	//userinfo.AddUsertoDB()
	//user, err := userinfo.GetUserInfoByIn()
	//if err != nil {
	//	fmt.Printf("获取数据失败[%s]\n", err.Error())
	//}
	//fmt.Printf("user is [%v]\n", user)

	userdata := NewUserInfo("修改测试", "17619871312", 0)
	err := userinfo.UpdateUserInfoById(userdata, 100)
	if err != nil {
		fmt.Printf("修改数据失败[%s]", err.Error())
	}
	//userDatas, err := userinfo.GetUserInfoAll()
	//if err != nil {
	//	fmt.Printf("获取数据失败[%s]", err.Error())
	//}
	//for _, userdata := range userDatas {
	//	fmt.Printf("user is [%v]\n", userdata)
	//}
	newDB.CloseDB()
}
