package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"r25ta.com/web-service-gin/model"
	albumService "r25ta.com/web-service-gin/service"
)

// albums slice to seed record album data.
var albums = []model.Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)

	router.GET("/albums/:id", getAlbumById)

	router.GET("/albums/artist/:artist", getAlbumByArtist)

	//Associate the POST method at the /albums path with the postAlbums function.
	router.POST("/albums", postAlbums)

	//Use the Run function to attach the router to an http.Server and start the server.
	router.Run("localhost:8180")

	//TEST => $ curl http://localhost:8080/albums
	//fmt.Print(albums)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(ctx *gin.Context) {
	albums, err := albumService.GetAllAlbums()

	if err != nil {
		log.Fatal("Error! please, try again!")
		ctx.IndentedJSON(http.StatusBadRequest, err)
	}

	ctx.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(ctx *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	err := ctx.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	//Add a 201 status code to the response, along with JSON representing the album you added.
	ctx.IndentedJSON(http.StatusCreated, albums)

}

func getAlbumById(ctx *gin.Context) {
	strId := ctx.Param("id")

	id, _ := strconv.ParseInt(strId, 10, 64)

	var album model.Album
	album, err := albumService.GetAlbumById(id)

	if err != nil {
		log.Fatal("Error! please, try again!")
		ctx.IndentedJSON(http.StatusBadRequest, err)
	}

	if album.ID > 0 {
		ctx.IndentedJSON(http.StatusOK, album)
		return
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found!"})
}

func getAlbumByArtist(ctx *gin.Context) {
	artist := ctx.Param("artist")

	albums, err := albumService.GetAlbumByArtist(artist)

	if err != nil {
		log.Fatal("Error! please, try again!")
		ctx.IndentedJSON(http.StatusBadRequest, err)
	}

	if len(albums) > 0 {
		ctx.IndentedJSON(http.StatusOK, albums)
		return
	}
	//Return an HTTP 404 error with http.StatusNotFound if the album isn’t found.
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found!"})
}
