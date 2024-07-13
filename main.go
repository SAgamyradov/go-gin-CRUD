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

	router.Run(":8080")
}
