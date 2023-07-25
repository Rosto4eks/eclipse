package database

import (
	"errors"
	"fmt"

	"github.com/Rosto4eks/eclipse/internal/models"
)

func (d *database) AddUser(user models.User) error {
	query := "INSERT INTO users (name,password) VALUES($1,$2)"
	if _, err := d.db.Exec(query, user.Name, user.Password); err != nil {
		d.logger.Error("database", "AddUser", err)
		return errors.New("user with this name is already exists")
	}
	d.logger.Info("database", "AddUser", fmt.Sprintf("%s registered", user.Name))
	return nil
}

func (d *database) DelUser(ID int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := d.db.Exec(query, ID)
	if err != nil {
		d.logger.Error("database", "DelUser", err)
		return err
	}
	d.logger.Info("databse", "DelUser", fmt.Sprintf("user with id = %d deleted", ID))
	return nil
}

func (d *database) GetUserByName(name string) (models.User, error) {
	query := "SELECT * FROM users WHERE name = $1"
	var response models.User
	if err := d.db.Get(&response, query, name); err != nil {
		d.logger.Error("database", "GetUserByName", err)
		return models.User{}, errors.New("there are no user in db with such credits.")
	}

	return response, nil
}
