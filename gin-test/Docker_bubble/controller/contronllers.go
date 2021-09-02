package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 添加一个事件
func AddTodo(c *gin.Context) {
	var todo models.Todo
	var err error
	c.BindJSON(&todo)
	if todo.Title == "" {
		c.JSON(http.StatusOK, "Title is NULL")
	}
	if err = models.AddTodo(&todo); err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		// ok
	}
}

// 更改一个事件
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if Dbtodo, err := models.FindTodo(c.Param("id")); err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.BindJSON(&todo)
		Dbtodo.Status = todo.Status
		if err = models.UpdateTodo(c.Param("id"), &Dbtodo); err != nil {
			c.JSON(http.StatusOK, err.Error())
		} else {
			// OK
		}
	}
}

// 查询全部数据
func FindTodoAll(c *gin.Context) {
	if todoList, err := models.FindTodoAll(); err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// 删除指定数据
func DeleteTodo(c *gin.Context) {
	if err := models.DeleteTodo(c.Param("id")); err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		// 目前不处理
	}
}
