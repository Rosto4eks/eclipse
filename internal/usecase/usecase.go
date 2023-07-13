package usecase

import (
	"mime/multipart"

	"github.com/Rosto4eks/eclipse/internal/database"
	"github.com/Rosto4eks/eclipse/internal/models"
)

type Iusecase interface {
	NewAlbum([]*multipart.FileHeader, models.Album) error
	GetUserByName(string) (models.User, error)
	SignUp(models.User) (string, error)
	SignIn(name, password string) (string, error)
	GetAlbumById(int) (models.AlbumResponse, error)
	Auth(string, string) error
}

type usecase struct {
	database database.Idatabase
}

func New(database database.Idatabase) *usecase {
	return &usecase{
		database,
	}
}
