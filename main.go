package main

import (
	"go-gin/model"
	"go-gin/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, model.Albums)
}

func postAlbums(ctx *gin.Context) {
	var newAlbum model.Album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}
	model.Albums = append(model.Albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range model.Albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "not found",
	})
}

func deleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, a := range model.Albums {
		if a.ID == id {
			model.Albums = append(model.Albums[:1], model.Albums[i+1:]...)
			ctx.IndentedJSON(http.StatusNoContent, a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "could not delete",
	})

}

func updateById(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, a := range model.Albums {
		if a.ID == id {
			ctx.BindJSON(&model.Albums[i])
			ctx.IndentedJSON(http.StatusOK, model.Albums[i])
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "could not update",
	})
}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteById)
	router.PUT("/albums/:id", updateById)

	router.Run(":8080")

	storage.Database()
}
