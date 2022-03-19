package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjohnep/spoti/db"
)

func main() {
	db.Connect()

	r := setupRouter()
	r.Run(":8008")
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/playlists", postPlaylist)

	r.GET("/playlists", getPlaylists)

	r.GET("/playlists/:id", getPlaylistByID)

	return r
}

func postPlaylist(c *gin.Context) {
	var newPlaylist db.Playlist
	c.BindJSON(&newPlaylist)

	playlistId, err := db.AddPlaylist(newPlaylist)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, playlistId)
	}
}

func getPlaylists(c *gin.Context) {
	playlists, err := db.GetPlaylists()

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, playlists)
}

func getPlaylistByID(c *gin.Context) {
	var idString = c.Param("id")
	idInt, err := strconv.ParseInt(idString, 10, 64)

	if err != nil {
		panic(err)
	}

	playlist, errQuery := db.GetPlaylistByID(idInt)
	fmt.Print(errQuery)
	if errQuery != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": errQuery.Error()})
	} else {
		c.JSON(http.StatusOK, playlist)
	}
}
