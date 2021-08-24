package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Code  string
	Price uint
}

func main() {

	fmt.Println("练习 gorm库")

	dsn := "root:123456@tcp(81.70.17.60:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("链接数据库失败err:[%s]\n", err)
	}
	fmt.Printf("链接数据库成功~\n")

	//db.Create(&Product{Code: "D43", Price: 100})

	var product Product

	db.Where("Price", 100).First(&product, "Code", "D42")
	fmt.Println(product)

	//db.Model(&product).Update("Price", 200)
	//db.Commit()
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	db.Delete(&product, 1)
	return
}
