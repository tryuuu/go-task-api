package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/tryuuu/go-task-api/internal/infrastructure"
	"github.com/tryuuu/go-task-api/internal/interface/handler"
	"github.com/tryuuu/go-task-api/internal/usecase"
)

func main() {
	// DB初期化
	db := infrastructure.InitDB()

	// 内側の層から順に構造体を生成
	userRepo := infrastructure.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// Ginルーター設定
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// handler(プレゼンテーション層)がhttpリクエストを受け取り、内部でusecase(アプリケーション層)を呼び出す
	r.POST("/signup", userHandler.SignupHandler)
	r.POST("/login", userHandler.LoginHandler)

	// サーバー起動
	r.Run(":8080")
}
