package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {

	var nodeName string
	if os.Getenv("kafka_listener_security_protocol_map") != "" {
		nodeName = os.Getenv("NodeName")
	} else {
		nodeName = "NULL"
	}
	fmt.Println("node Name is", nodeName)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"nodeName": nodeName,
		})
	})

	r.Run(":" + "8080")

	return
}
