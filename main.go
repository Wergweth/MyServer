package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type user struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	connection, err := sqlx.Open("mysql", "root:1q2w3e@(127.0.0.1:3306)/test1")
	if err != nil {
		return
	}
	var result []user
	err = connection.Select(&result, "select * from users where age = ?", 37)
	println(result[0])
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
