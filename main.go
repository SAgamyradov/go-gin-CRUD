package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Actor string  `json:"actor"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Film", Actor: "Johny Dep", Price: 34.23},
	{ID: "2", Title: "Jeru", Actor: "Gerry Mulligan", Price: 17.99},
}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(ctx *gin.Context) {
	var newAlbum album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range albums {
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

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:1], albums[i+1:]...)
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

	for i, a := range albums {
		if a.ID == id {
			ctx.BindJSON(&albums[i])
			ctx.IndentedJSON(http.StatusOK, albums[i])
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
}
