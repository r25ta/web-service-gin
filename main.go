package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	model "example.com/web-service-gin/model"
)

// albums slice to seed record album data.
var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	//Initialize a Gin router using Default.
	//gin.SetMode("release")
	router := gin.Default()

	/*Use the GET function to associate the GET HTTP method and /albums path with a handler function.
	Note that you’re passing the name of the getAlbums function.
	This is different from passing the result of the function,
	which you would do by passing getAlbums() (note the parenthesis).
	*/
	router.GET("/albums", getAlbums)

	router.GET("/albums/:id", getAlbumById)

	//Associate the POST method at the /albums path with the postAlbums function.
	router.POST("/albums", postAlbums)

	//Use the Run function to attach the router to an http.Server and start the server.
	router.Run("localhost:8180")

	//TEST => $ curl http://localhost:8080/albums
	//fmt.Print(albums)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(ctx *gin.Context) {
	/*Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	The function’s first argument is the HTTP status code you want to send to the client.
	Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.
	*/
	ctx.IndentedJSON(http.StatusOK, albums)

	/*Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON.
	In practice, the indented form is much easier to work with
	when debugging and the size difference is usually small.
	*/
	//ctx.JSON(http.StatusOK, albums)
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

/*
getAlbumByID locates the album whose ID value matches the id

	parameter sent by the client, then returns that album as a response.
*/
func getAlbumById(ctx *gin.Context) {
	/*Use Context.Param to retrieve the id path parameter from the URL.
	When you map this handler to a path, you’ll include a placeholder for the parameter in the path.
	*/
	id := ctx.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, idAlb := range albums {
		if idAlb.ID == id {
			ctx.IndentedJSON(http.StatusOK, idAlb)
			return
		}
	}
	//Return an HTTP 404 error with http.StatusNotFound if the album isn’t found.
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found!"})
}
