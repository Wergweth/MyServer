package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func main() {
	sql.Open("mysql", "root:1q2w3e@(127.0.0.1:3306)/test1")
	engine := gin.New()
	engine.GET("", MyHandler)
	engine.GET("/ebalo34", MyHandler2)
	engine.Run("localhost:8080")
}

func MyHandler(c *gin.Context) {
	c.String(200, "ebalo")
}

func MyHandler2(c *gin.Context) {
	c.String(200, "Urrra")
}
