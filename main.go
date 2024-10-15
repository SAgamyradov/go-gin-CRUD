package main

import (
	"go-gin/database"
	"go-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.Init()

	router := gin.Default()

	routes.InitRoutes(router, db)
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})

	router.Run(":8080")
}
