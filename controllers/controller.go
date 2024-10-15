package controllers

import (
	"go-gin/database"
	model "go-gin/models"
	"go-gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAlbum(ctx *gin.Context) {
	db := database.Init()

	var albums []model.Album
	if err := db.Find(&albums).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "failed to get products"})
		return
	}
	ctx.JSON(http.StatusOK, albums)
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
	ctx.JSON(http.StatusOK, album)

}

func CreateAlbum(ctx *gin.Context) {
	db := database.Init()
	ctx.Set("db", db)
	var newAlbum model.Album
	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request album"})
		return
	}
	if err := utils.ValidateStruct(ctx, newAlbum); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
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
	ctx.JSON(200, gin.H{"removed id:": id})
}
