package db

import (
	"fmt"
	_ "fmt"
	"testing"
)

// run TestInitDb
func TestInitDB(t *testing.T) {

	fmt.Println("开始测试~")
	newDB := NewDBinfo("81.70.17.60", "3306", "TYServer", "root", "123456")
	newDB.InitDB()
	fmt.Println("连接数据库成功！")

	newDB.SetConnPool(50, 10)
	fmt.Printf("查询数据库连接状态[%v]", Db.Stats())
	newDB.CloseDB()
}

// run TestInitDb
func TestNewMySqlDBInfo(t *testing.T) {

	fmt.Println("开始测试~")
	newDB := NewMySqlDBInfo("192.168.0.111", "3306", "root", "sys123456", "test", "utf8mb4", true)
	db, err := newDB.InitMysql()
	if err != nil {
		fmt.Println("TEST FAIL : ", err)
	}
	fmt.Println("连接数据库成功！", db)
}
