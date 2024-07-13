package routes

import (
	"go-gin/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {

	router.POST("/product", controllers.CreateAlbum)
	router.GET("/product", controllers.GetAlbum)
	router.GET("/product/:id", controllers.GetAlbumById)
	router.PUT("/product/:id", controllers.UpdateAlbum)
	router.DELETE("/product/:id", controllers.DeleteAlbum)
}
