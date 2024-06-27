package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	constant "example.com/web-service-gin/constant"
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

}
