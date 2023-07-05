package database

import (
	"fmt"
	"github.com/Rosto4eks/eclipse/internal/models"
)

func (d *database) AddUser(user models.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO users (name,password) VALUES($1,$2) RETURNING id;")
	row := d.db.QueryRow(query, user.Name, user.Password)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (d *database) DelUser(ID int) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id = $1;", ID)
	result, err := d.db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Print(result)
	return nil
}

func (d *database) GetUsersByName(name string) (models.User, error) {
	query := fmt.Sprintf("SELECT id, name, password, role FROM users WHERE name = $1;", name)
	row, err := d.db.Query(query)
	if err != nil {
		return models.User{}, err
	}

	var response models.User
	if err = row.Scan(&response); err != nil {
		return models.User{}, err
	}

	return response, nil
}
