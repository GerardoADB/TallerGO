package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("albums.sqlite"), &gorm.Config{})
	if err != nil {
		return
	}
	var albums []album
	db.Find(&albums)
	c.IndentedJSON(200, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	db, err := gorm.Open(sqlite.Open("albums.sqlite"), &gorm.Config{})

	if err != nil {
		return
	}

	var album album

	err = db.First(&album, id).Error

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, &album)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	db, err := gorm.Open(sqlite.Open("albums.sqlite"), &gorm.Config{})

	if err != nil {
		return
	}
	err = db.Delete(&album{}, id).Error
	if err != nil {
		return
	}

}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("albums.sqlite"), &gorm.Config{})
	if err != nil {
		return
	}
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	db.Create(&newAlbum)
	c.IndentedJSON(200, newAlbum)
}

func main() {
	db, err := gorm.Open(sqlite.Open("albums.sqlite"), &gorm.Config{})
	if err != nil {
		return
	}

	db.AutoMigrate(&album{})

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("localhost:8080") //jamas poner rutas despues de iniciar el servidor

}
