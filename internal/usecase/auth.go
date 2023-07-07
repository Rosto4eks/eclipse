package usecase

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/Rosto4eks/eclipse/internal/models"
)

func (u *usecase) NewUser(usr models.User) error {
	usr.Password = hash(usr.Password)
	return u.database.AddUser(usr)
}

func hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
