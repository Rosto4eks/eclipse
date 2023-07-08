package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/Rosto4eks/eclipse/internal/models"
)

func (u *usecase) NewUser(usr models.User) error {
	usr.Password = hash(usr.Password)
	return u.database.AddUser(usr)
}

func (u *usecase) SignIn(name, password string) error {
	usr, err := u.database.GetUserByName(name)
	if err != nil {
		return err
	} else if usr.Password == hash(password) {
		return nil
	}
	return errors.New("there are no user with such credits")
}

func hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
