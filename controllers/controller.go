package controllers

import (
	"go-gin/database"
	model "go-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AlbumController struct {
	DB *gorm.DB
}

func NewAlbumController(db *gorm.DB) *AlbumController {
	return &AlbumController{DB: db}
}

func GetAlbum(ctx *gin.Context) {
	db := database.Init()

	var albums []model.Album
	if err := db.Find(&albums).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "failed to get products"})
		return
	}
	ctx.JSON(200, albums)
}

func GetAlbumById(ctx *gin.Context) {
	db := database.Init()

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid album ID"})
		return
	}
	var album model.Album
	if err := db.First(&album, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "album not found"})
		return
	}
	ctx.JSON(200, album)

}

func CreateAlbum(ctx *gin.Context) {
	db := database.Init()

	var newAlbum model.Album
	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request album"})
		return
	}
	if err := db.Create(&newAlbum).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "failed create album"})
		return
	}
	ctx.JSON(200, gin.H{"success": "create album", "album": newAlbum})
}

func UpdateAlbum(ctx *gin.Context) {
	db := database.Init()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid album ID"})
		return
	}
	var album model.Album
	if err := db.First(&album, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "album not found"})
		return
	}
	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request album"})
		return
	}
	if err := db.Save(&album).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "failed update album"})
		return
	}
	ctx.JSON(200, gin.H{"success": "album updated", "album": album})

}

func DeleteAlbum(ctx *gin.Context) {
	db := database.Init()

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid album ID"})
		return
	}

	var album model.Album
	if err := db.First(&album, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Album not found"})
		return
	}
	if err := db.Delete(&album).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "failed to delete album"})
		return
	}
	ctx.JSON(200, gin.H{"success": "removed album"})
}
