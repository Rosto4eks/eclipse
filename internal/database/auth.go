package database

import (
	"github.com/Rosto4eks/eclipse/internal/models"
)

func (d *database) AddUser(user models.User) (int, error) {
	query := "INSERT INTO users (name,password, role) VALUES($1,$2,$3) RETURNING id;"
	var id int
	if err := d.db.Get(&id, query, user.Name, user.Password, user.Role); err != nil {
		return 0, err
	}

	return id, nil
}

func (d *database) DelUser(ID int) error {
	query := "DELETE FROM users WHERE id = $1;"
	_, err := d.db.Exec(query, ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *database) GetUsersByName(name string) (models.User, error) {
	query := "SELECT id, name, password, role FROM users WHERE name = $1;"
	var response models.User
	if err := d.db.Get(&response, query, name); err != nil {
		return models.User{}, err
	}

	return response, nil
}
