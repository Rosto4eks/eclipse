package database

import (
	"fmt"
	"github.com/Rosto4eks/eclipse/internal/models"
)

func (d *database) GetAllAlbums() ([]models.Album, error) {
	query := fmt.Sprintf("SELECT * FROM albums;")
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}

	var response []models.Album
	if err = rows.Scan(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func (d *database) GetAlbumByID(ID int) (models.Album, error) {
	query := fmt.Sprintf("SELECT name, author_id, images_count, date, description FROM albums WHERE id = $1;", ID)
	row := d.db.QueryRow(query)

	var response models.Album
	if err := row.Scan(&response); err != nil {
		return models.Album{}, err
	}
	return response, nil
}

func (d *database) AddAlbum(album models.Album) error {
	query := fmt.Sprintf("INSERT INTO albums (name, author_id, images_count, date, description) VALUES($1,$2,$3,$4,$5);",
		album.Name, album.Author, album.Count, album.Date, album.Description)
	result, err := d.db.Exec(query)
	if err != nil {
		return nil
	}
	fmt.Print(result)
	return nil
}

func (d *database) DelAlbumByID(ID int) error {
	query := fmt.Sprintf("DELETE FROM albums WHERE id = $1;", ID)
	result, err := d.db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Print(result)
	return nil
}
