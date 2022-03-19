package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Playlist struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

var db *sql.DB

func Connect() {
	fmt.Println("TRY CONNECT")
	connStr := "dbname=spoti sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func GetPlaylists() ([]Playlist, error) {
	var playlists []Playlist

	fmt.Println("TRY GET PLAYLISTS", db)
	rows, err := db.Query("SELECT * FROM playlists")

	if err != nil {
		return nil, fmt.Errorf("getPlaylists %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var playlist Playlist
		if err := rows.Scan(&playlist.ID, &playlist.Title, &playlist.Artist, &playlist.Price); err != nil {
			return nil, fmt.Errorf("getPlaylists %v", err)
		}
		playlists = append(playlists, playlist)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getPlaylists %v", err)
	}

	return playlists, nil
}

func GetPlaylistByID(id int64) (Playlist, error) {
	var playlist Playlist

	row := db.QueryRow("SELECT * FROM playlists WHERE id = $1", id)
	if err := row.Scan(&playlist.ID, &playlist.Title, &playlist.Artist, &playlist.Price); err != nil {
		if err == sql.ErrNoRows {
			return playlist, fmt.Errorf("getPlaylistByID %d: no such playlist : %v", id, err)
		}
		return playlist, fmt.Errorf("getPlaylistByID %d: %v", id, err)
	}

	return playlist, nil
}

func AddPlaylist(playlist Playlist) (int64, error) {
	var playlistId int64
	err := db.QueryRow("INSERT INTO playlists (title, artist, price) VALUES ($1, $2, $3) RETURNING id", playlist.Title, playlist.Artist, playlist.Price).Scan(&playlistId)
	if err != nil {
		return 0, fmt.Errorf("addPlaylist %v", err)
	}

	return playlistId, nil
}
