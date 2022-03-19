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

	r.GET("/login", spotifyLogin)
	r.GET("/callback", loginCallback)

	return r
}

func loginCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login Callback"})
}

func spotifyLogin(c *gin.Context) {
	var clientId = "f62d9fb9fc5c4d2192d49a897312ff69"
	var redirectCallback = "http://localhost:8008/callback"
	var scope = "user-read-private user-read-email"

	req, err := http.NewRequest("GET", "https://accounts.spotify.com/authorize", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("response_type", "code")
	q.Add("client_id", clientId)
	q.Add("scope", scope)
	q.Add("redirect_uri", redirectCallback)

	req.URL.RawQuery = q.Encode()

	c.Redirect(http.StatusMovedPermanently, req.URL.String())
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
