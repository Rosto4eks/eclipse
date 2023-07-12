package database

import (
	"fmt"

	"github.com/Rosto4eks/eclipse/internal/models"
)

func (d *database) GetAllAlbums() ([]models.Album, error) {
	query := "SELECT * FROM albums"
	var response []models.Album
	if err := d.db.Select(&response, query); err != nil {
		return nil, err
	}

	return response, nil
}

func (d *database) GetAlbumByID(ID int) (models.Album, error) {
	query := "SELECT id, TRIM(name) as name, author_id, images_count, date, description FROM albums WHERE id = $1"
	var response models.Album
	if err := d.db.Get(&response, query, ID); err != nil {
		fmt.Println(err)
		return models.Album{}, err
	}
	fmt.Println(response)
	return response, nil
}

func (d *database) AddAlbum(album models.Album) error {
	query := "INSERT INTO albums (name, author_id, images_count, date, description) VALUES($1,$2,$3,$4,$5)"
	if _, err := d.db.Exec(query, album.Name, album.Author_id, album.Images_count, album.Date, album.Description); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (d *database) DelAlbumByID(ID int) error {
	query := "DELETE FROM albums WHERE id = $1"
	if _, err := d.db.Exec(query, ID); err != nil {
		return err
	}

	return nil
}
