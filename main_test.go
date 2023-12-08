package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func setupRouter() *gin.Engine {

	router := gin.Default()
	validApiKey := os.Getenv("APIKey")
	router.Use(ApiKeyMiddleware(validApiKey))
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	return router
}

func TestGetListOfAlbums(t *testing.T) {

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetAlbumByID(t *testing.T) {

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
