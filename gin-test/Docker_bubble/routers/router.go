package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func StartRouter() (r *gin.Engine) {
	r = gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	r.GET("/", controller.Index)

	V1Group := r.Group("/v1")
	{
		V1Group.POST("/todo", controller.AddTodo)
		V1Group.PUT("/todo/:id", controller.UpdateTodo)
		V1Group.GET("/todo", controller.FindTodoAll)
		V1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return
}
