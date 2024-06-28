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

	fmt.Println(getAbumById(99))
}

func getAbumById(id int64) (model.Album, error) {
	var alb model.Album
	conDb := getConnection()

	row := conDb.QueryRow("SELECT *FROM album WHERE id = $1", id)

	err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)

	if err == sql.ErrNoRows {
		return alb, fmt.Errorf("albumsById %d: no such album", id)

	} else if err != nil {
		return alb, fmt.Errorf("albumsById %d: %v", id, err)

	} else {
		return alb, nil

	}
}
