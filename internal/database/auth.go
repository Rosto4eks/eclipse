package database

import (
	"errors"

	"github.com/Rosto4eks/eclipse/internal/models"
)

func (d *database) AddUser(user models.User) error {
	query := "INSERT INTO users (name,password) VALUES($1,$2)"
	if _, err := d.db.Exec(query, user.Name, user.Password); err != nil {
		return errors.New("user with this name is already exists")
	}
	return nil
}

func (d *database) DelUser(ID int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := d.db.Exec(query, ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *database) GetUserByName(name string) (models.User, error) {
	query := "SELECT * FROM users WHERE name = $1"
	var response models.User
	if err := d.db.Get(&response, query, name); err != nil {
		return models.User{}, errors.New("there are no user in db with such credits.")
	}

	return response, nil
}
