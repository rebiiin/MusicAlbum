package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Singer string  `json:"singer"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "The record", Singer: "boygenius", Price: 20.00},
	{ID: "2", Title: "Local Natives - Time Will Wait For No One", Singer: "Indie Rock", Price: 23.99},
	{ID: "3", Title: "Clutch - Sunrise On Slaughter Beach", Singer: "Sarah Vaughan", Price: 9.09},
	{ID: "4", Title: "Tears For Fears - The Tipping Point", Singer: "Sarah Vaughan", Price: 14.17},
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	validApiKey := os.Getenv("APIKey")
	router := gin.Default()
	router.Use(ApiKeyMiddleware(validApiKey))
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:5000")

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func ApiKeyMiddleware(apiKey string) gin.HandlerFunc {

	return func(c *gin.Context) {

		clientApiKey := c.GetHeader("API-Key")

		if clientApiKey != apiKey {

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Invalid API key"})
			return
		}

		c.Next()
	}
}
