package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tryuuu/go-task-api/internal/infrastructure"
	"github.com/tryuuu/go-task-api/internal/interface/handler"
)

func main() {
	infrastructure.InitDB()
	r := gin.Default()

	// CORS ミドルウェアを追加
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.POST("/signup", handler.SignupHandler)

	r.Run(":8080")
}
