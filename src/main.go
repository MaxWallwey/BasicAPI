package main

import (
	"basic-api/cassandra"
	"basic-api/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	Session := cassandra.SetupCassandra()
	defer Session.Close()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	log.Println("get albums")
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	log.Println(fmt.Sprintf("get album with id: %s", id))
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum types.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	log.Println(fmt.Sprintf("post album"))
}
