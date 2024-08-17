package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"r25ta.com/web-service-gin/internal/utility"
	"r25ta.com/web-service-gin/model"
)

func GetAllAlbums() ([]model.Album, error) {
	connStr := utility.ConnectionString()
	connDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Não foi possivel estabelecer conexão com BD %s", err)
	}

	rows, err := connDB.Query("SELECT * FROM album")

	if err != nil {
		log.Fatal("Error! please, try again!")
		return nil, fmt.Errorf("albums: %v", err)
	}

	defer rows.Close()
	defer connDB.Close()

	var albums []model.Album

	for rows.Next() {
		var alb model.Album

		err := rows.Scan(&alb.ID,
			&alb.Title,
			&alb.Artist,
			&alb.Price)

		if err != nil {
			log.Fatal("Error, in Find!")
			return nil, fmt.Errorf("albuns: %v", err)
		}
		albums = append(albums, alb)

	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtists: %v", err)
	}
	return albums, nil

}
func GetAlbumByArtist(artist string) ([]model.Album, error) {
	connStr := utility.ConnectionString()
	//connStr := "postgres://postgres:admin@localhost:5432/recordings?sslmode=disable"
	connDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Não foi possivel estabelecer conexão com BD %s", err)
	}

	rows, err := connDB.Query("SELECT * FROM album WHERE artist LIKE $1", "%"+artist+"%")

	if err != nil {
		log.Fatal("Error! please, try again!")
		return nil, fmt.Errorf("albums: %v", err)
	}

	defer rows.Close()
	defer connDB.Close()

	var albums []model.Album

	for rows.Next() {
		var alb model.Album
		err = rows.Scan(
			&alb.ID,
			&alb.Title,
			&alb.Artist,
			&alb.Price,
		)
		if err != nil {
			log.Fatal("Error, in Find!")
			return nil, fmt.Errorf("albuns: %v", err)
		}
		albums = append(albums, alb)

	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtists: %v", err)
	}
	return albums, nil

}
func GetAlbumById(id int64) (model.Album, error) {

	connStr := utility.ConnectionString()
	connDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Não foi possivel estabelecer conexão com BD %s", err)
	}

	row := connDB.QueryRow("SELECT *FROM album WHERE id = $1", id)
	var alb model.Album

	err = row.Scan(
		&alb.ID,
		&alb.Title,
		&alb.Artist,
		&alb.Price,
	)
	defer connDB.Close()

	if err == sql.ErrNoRows {
		return alb, fmt.Errorf("albumsById %d: no such album", id)

	} else if err != nil {
		return alb, fmt.Errorf("albumsById %d: %v", id, err)

	} else {
		return alb, nil

	}
}
