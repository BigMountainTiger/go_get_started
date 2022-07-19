package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"song.com/go_get_started/web_service_gin/aaa"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Song Li", Artist: "Song li-1", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	load_auth_key()
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/OK", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, aaa.Geta())
	})

	router.GET("/auth", func(c *gin.Context) {

		result, err := authenticate(c)

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
		} else {
			c.IndentedJSON(http.StatusOK, result)
		}
	})

	router.GET("/auth-key", func(c *gin.Context) {

		result, err := auth_key()

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
		} else {
			c.IndentedJSON(http.StatusOK, result)
		}
	})

	router.Run("localhost:8080")
	fmt.Println("Server started")
}
