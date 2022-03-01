package db

import (
	"fmt"
	_ "fmt"
	"testing"
)

// run TestInitDb
func TestInitDB(t *testing.T) {

	fmt.Println("开始测试~")
	newDB := NewDBinfo("192.168.0.87", "3306", "mis_go_new", "root", "p@ssw0rd")
	newDB.InitDB()
	fmt.Println("连接数据库成功！")

	newDB.SetConnPool(50, 10)
	fmt.Printf("查询数据库连接状态[%v]", Db.Stats())
	newDB.CloseDB()
}
