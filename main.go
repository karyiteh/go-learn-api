package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
// In a real use case this will be from a database.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "folklore", Artist: "Taylor Swift", Price: 13.99},
}

// Main handler code for web API
func main() {
	// Initialize the router
	router := gin.Default()

	// Providing the handler functions
	router.GET("/albums", getAlbums)

	// Attach router to http.Server and start server
	router.Run("localhost:8080")
}

// gin.Context carries request details, validates and serializes JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
