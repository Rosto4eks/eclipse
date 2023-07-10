package usecase

import (
	"mime/multipart"

	"github.com/Rosto4eks/eclipse/internal/database"
	"github.com/Rosto4eks/eclipse/internal/models"
)

type Iusecase interface {
	NewAlbum([]*multipart.FileHeader, models.Album) error
	GetUserByName(string) (models.User, error)
	NewUser(models.User) error
	SignIn(name, password string) error
	GetAlbumById(int) (models.Album, error)
}

type usecase struct {
	database database.Idatabase
}

func New(database database.Idatabase) *usecase {
	return &usecase{
		database,
	}
}
