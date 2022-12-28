package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func router() *gin.Engine{
	r := gin.Default()
	r.GET("/index", service.GetIndex) 

	return r
}

// 将gin例子中的代码拆成三块
// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }