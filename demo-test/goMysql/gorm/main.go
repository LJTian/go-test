package main

import (
	"fmt"
	"go_test/goMysql/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func FuncInSet(db *gorm.DB) {
	// 单条插入
	//User := model.User{
	//	Name:   "ljtian",
	//	Age:    27,
	//	IcCard: "11032132131231231231",
	//}
	//dbResp := db.Create(&User)
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	// 指定字段插入
	//user2 := model.User{Name: "zhangsan", Age: 18}
	//dbResp = db.Select("Name", "Age", "CreatedAt").Create(&user2)
	//// insert into user (name, age) values ("zhangsan",18);
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	// 批量插入
	//var userOnce model.User
	//users := make([]model.User, 0)

	//for i := 0; i < 10; i++ {
	//	userOnce.Name = fmt.Sprintf("ljtian%d", i)
	//	userOnce.Age = uint8(i + 25)
	//	userOnce.IcCard = fmt.Sprintf("11032132131231231231%d", i)
	//	users = append(users, userOnce)
	//}
	//dbResp = db.Create(&users)
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	// 批量定量插入
	//for i := 0; i < 1000; i++ {
	//	userOnce.Name = fmt.Sprintf("ljtian%d", i)
	//	userOnce.Age = uint8(i + 25)
	//	userOnce.IcCard = fmt.Sprintf("11032132131231231231%d", i)
	//	users = append(users, userOnce)
	//}
	//dbResp := db.CreateInBatches(&users, 100)
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	// 钩子测试
	//User := model.User{
	//	Name:   "123",
	//	Age:    27,
	//	IcCard: "11032132131231231231",
	//}
	//dbResp := db.Create(&User)
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	// map插入 map插入没有 调用hook函数
	//dbResp := db.Model(&model.User{}).Create(map[string]interface{}{
	//	"Name": "name321", "Age": 21,
	//})
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	// map + sql自定义
	//dbResp := db.Model(&model.User{}).Create(map[string]interface{}{
	//	"Name": "自定义sqlMap",
	//	"Age": clause.Expr{
	//		SQL:  "round(RAND()*?%?/1)",
	//		Vars: []interface{}{100, 100},
	//	},
	//})
	//
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	// 自定义类型创建记录 + sql
	//dbResp := db.Create(&model.Commodity{
	//	Name:       "商品1",
	//	CommodCode: model.Code{X: 1000, Y: 1000},
	//})
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	// 高级功能 关联创建 重点【设置好 外键】
	//dbResp := db.Create(&model.Order{
	//	UserId: "0001",
	//	Name:   "商品2",
	//
	//	Commodity: model.Commodity{
	//		CommodCode: "sp2",
	//	},
	//})
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	// 跳过关联创建
	//dbResp := db.Omit("Commodity").Create(&model.Order{
	//	UserId: "0001",
	//	Name:   "商品3",
	//
	//	Commodity: model.Commodity{
	//		CommodCode: "sp2",
	//	},
	//})
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	//data := db.Find(&model.User{}, "name = ?", "ljtian")
	//if data.Error != nil {
	//	fmt.Println(data.Error.Error())
	//}
	//data.Row().Scan(&user1)
	//fmt.Println(user1.Name)
}

func FuncSelect(db *gorm.DB) {

	// var user1 model.User

	// 获取第一条记录（主键升序）
	//dbResp := db.First(&user1)
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	// 随机获取一条记录
	//dbResp := db.Take(&user1)
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}
	//// SELECT * FROM users LIMIT 1;

	// 倒叙去一条记录
	//dbResp := db.Last(&user1)
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	// 获取数据
	//dbResp := db.First(&user1)
	//if errors.Is(dbResp.Error, gorm.ErrRecordNotFound) {
	//	fmt.Println("没有找到数据")
	//} else if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}
	//fmt.Printf("获取的数据条数为[%d]\n", dbResp.RowsAffected)

	// model
	// result := map[string]interface{}{}
	// db.Model(&model.User{}).First(&result)
	// OR
	// table
	//result := map[string]interface{}{}
	//db.Table("users").Take(&result)

	//db.First(&user1, 35)
	//fmt.Println(user1)
	//fmt.Println(user1)

	// var users []model.User

	// 获取全部
	// result := db.Find(&users)

	// 根据条件查询
	//result := db.Where("Age < ? ", 50).Find(&users)
	//fmt.Println(result.RowsAffected)
	//if result.Error != nil {
	//	fmt.Println(result.Error.Error())
	//}

	var user model.User
	// Struct
	db.Where(&model.User{Name: "zhangsan", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
	fmt.Println(user)

	//// Map
	var users []model.User
	// db.Where(map[string]interface{}{"name": "zhangsan", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 主键切片条件
	db.Where("age", []int64{30, 31, 32}).Find(&users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);
	//
	for k, v := range users {
		fmt.Printf("[%d]:[%d:%s:%d:%s]\n", k, v.ID, v.Name, v.Age, v.IcCard)
	}
}

func FuncDelete(db *gorm.DB) {

	// 软删除，就是将deleted_at字段添加上值
	//dbResp := db.Where("name = ?", "zhangsan").Delete(&model.User{})
	//if dbResp.Error != nil {
	//	fmt.Println(dbResp.Error.Error())
	//}

	// 硬删除
	dbResp := db.Unscoped().Where("name = ?", "lisi").Delete(&model.User{})
	if dbResp.Error != nil {
		fmt.Println(dbResp.Error.Error())
	}

}

func FuncUpdate(db *gorm.DB) {
	dbResp := db.Model(&model.User{}).Where("age > ?", 80).Update("name", "lisi")
	if dbResp.Error != nil {
		fmt.Println(dbResp.Error.Error())
	}
}

func FuncNativeSql(db *gorm.DB) {

	type Result struct {
		ID   int
		Name string
		Age  int
	}

	// 原生查询
	var result []Result
	dbResp := db.Raw("select id, name, age from users where id > ?", 1000).Scan(&result)
	if dbResp.Error != nil {
		fmt.Println(dbResp.Error.Error())
	}

	for k, v := range result {
		fmt.Printf("[%d] [%v]\n", k, v)
	}

	// 原生执行
	dbResp = db.Exec("delete from users where id > ?", 100)
	if dbResp.Error != nil {
		fmt.Println(dbResp.Error.Error())
	}
	fmt.Printf("删除的条数为[%d]\n", dbResp.RowsAffected)
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:23306)/gromTest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("链接成功")

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Commodity{})
	db.AutoMigrate(&model.Order{})
	// db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&model.User{})
	// db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&model.Commodity{})

	//FuncInSet(db) // 增
	//FuncSelect(db) // 查
	// FuncDelete(db) // 删
	// FuncUpdate(db) // 改
	FuncNativeSql(db) // 原生sql

	fmt.Println("退出")
}
