package main

import (
	//"net/http"
	"unit3_IM/ginchat/router"
	//"github.com/gin-gonic/gin"
)
// 将gin例子中的代码拆成三块
func main() {
	r := router.Router()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
