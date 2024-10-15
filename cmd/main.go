package main

import (
	"go-gin/internal/app/repository"
	"go-gin/internal/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db := repository.Init()

	router := gin.Default()

	routes.InitRoutes(router, db)
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})

	router.Run(":8080")
}
