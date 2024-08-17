package albumService

import (
	"log"

	"r25ta.com/web-service-gin/model"
	"r25ta.com/web-service-gin/repository"
)

func GetAllAlbums() ([]model.Album, error) {
	albums, err := repository.GetAllAlbums()

	if err != nil {
		log.Fatal("Error! please, try again!")
		return nil, err
	}

	return albums, nil
}

func GetAlbumById(id int64) (model.Album, error) {
	var album model.Album
	album, err := repository.GetAlbumById(id)

	if err != nil {
		log.Fatal("Error! please, try again!")
		return album, err
	}

	return album, nil
}

func GetAlbumByArtist(artist string) ([]model.Album, error) {

	albums, err := repository.GetAlbumByArtist(artist)

	if err != nil {
		log.Fatal("Error! please, try again!")
		return nil, err
	}

	return albums, nil
}
