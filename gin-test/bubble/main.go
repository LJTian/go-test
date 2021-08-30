package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var (
	G_db *gorm.DB
)

// model
type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func addTodo(c *gin.Context) {
	var todo Todo
	c.ShouldBind(&todo)
	if todo.Title == "" {
		fmt.Println("db is NULL")
		return
	}
	// 创建数据到数据库
	G_db.Create(&todo)
	// 跳转到首页
	return
}

func updateTodo(c *gin.Context) {
	var todo Todo
	c.ShouldBind(&todo)
	ID := c.Param("id")
	fmt.Println("ID is ", ID)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	todo.ID = id
	G_db.Model(&Todo{}).Debug().Where("ID = ?", todo.ID).Update("Status", todo.Status)
	return
}

func findTodoAll(c *gin.Context) {

	var todo []Todo

	G_db.Debug().Find(&todo)

	c.JSON(http.StatusOK, &todo)
}

func findTodo(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func deleteTodo(c *gin.Context) {
	var todo Todo
	c.ShouldBind(&todo)
	ID := c.Param("id")
	fmt.Println("ID is ", ID)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	todo.ID = id
	G_db.Debug().Delete(&Todo{}, todo.ID)
	return
}

func main() {
	// 初始化数据库
	sTmpBuf := "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(sTmpBuf), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	// 迁移
	db.AutoMigrate(&Todo{})
	G_db = db

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	r.GET("/", index)

	V1Group := r.Group("/v1")
	{
		V1Group.POST("/todo", addTodo)
		V1Group.PUT("/todo/:id", updateTodo)
		V1Group.GET("/todo", findTodoAll)
		V1Group.GET("/todo/:id", findTodo)
		V1Group.DELETE("/todo/:id", deleteTodo)
	}

	r.Run(":8005")
}
