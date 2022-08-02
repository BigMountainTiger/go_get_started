package main

import (
	"log"
	"net/http"
	"os"

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
	ok := init_auth()
	if !ok {
		log.Fatal("Unable to init auth capability")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/OK", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, aaa.Geta())
	})

	router.GET("/profile/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	router.GET("/auth", func(c *gin.Context) {

		result, err := authenticate(c)

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
		} else {
			c.IndentedJSON(http.StatusOK, result)
		}
	})

	router.GET("/get-item", func(c *gin.Context) {

		_, err := authenticate(c)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
			return
		}

		result, err := get_dynamo_entry()

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
		} else {
			c.IndentedJSON(http.StatusOK, result)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	router.Run("0.0.0.0:" + port)
	log.Println("Server started")
}
