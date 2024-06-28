package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	constant "example.com/web-service-gin/constant"
	model "example.com/web-service-gin/model"
)

func getConnection() (con *sql.DB) {

	dbCredencials := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", constant.USER, constant.PWD, constant.SERVER, constant.PORT, constant.DATABASE)

	conDb, conErr := sql.Open("postgres", dbCredencials)

	if conErr != nil {
		log.Fatal("Error connecting to the database", conErr)
		return nil
	}

	//	defer conDb.Close()
	return conDb
}

func main() {

	conDb := getConnection()

	pingErr := conDb.Ping()

	if pingErr != nil {
		log.Fatal("Error ")

	}
	fmt.Println("Connected in database!")

	fmt.Println(getAllAbums())
}

func getAllAbums() ([]model.Album, error) {
	var albums []model.Album

	conDb := getConnection()

	rows, err := conDb.Query("SELECT * FROM album")

	if err != nil {
		log.Fatal("Error! please, try again!")
		return nil, fmt.Errorf("albums: %v", err)
	}

	defer rows.Close()
	defer conDb.Close()

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

func getAbumById(id int64) (model.Album, error) {
	var alb model.Album
	conDb := getConnection()

	row := conDb.QueryRow("SELECT *FROM album WHERE id = $1", id)

	err := row.Scan(
		&alb.ID,
		&alb.Title,
		&alb.Artist,
		&alb.Price,
	)
	defer conDb.Close()

	if err == sql.ErrNoRows {
		return alb, fmt.Errorf("albumsById %d: no such album", id)

	} else if err != nil {
		return alb, fmt.Errorf("albumsById %d: %v", id, err)

	} else {
		return alb, nil

	}
}
