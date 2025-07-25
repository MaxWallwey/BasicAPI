package main

import (
	"basic-api/db"
	"basic-api/types"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	Session := db.SetupCassandra()
	defer Session.Close()

	var album types.Album

	if err := Session.Query("SELECT id, title, artist, price, stockedSince FROM albums WHERE id = ? LIMIT 1", id).WithContext(context.Background()).Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.StockedSince); err != nil {
		fmt.Println("select error")
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func setupCassandra() *gocql.Session {
	cluster := gocql.NewCluster(cassandraHost)
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{
		NumRetries: 3,
	}
	cluster.Keyspace = "roster"
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cassandraUser,
		Password: cassandraPassword,
	}
	var session *gocql.Session
	var err error
	for {
		session, err = cluster.CreateSession()
		if err == nil {
			break
		}
		log.Printf("CreateSession: %v", err)
		time.Sleep(time.Second)
	}
	log.Printf("Connected OK\n")
	return session
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum types.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
