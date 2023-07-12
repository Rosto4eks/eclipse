package usecase

import (
	"github.com/labstack/echo/v4"
	"mime/multipart"

	"github.com/Rosto4eks/eclipse/internal/database"
	"github.com/Rosto4eks/eclipse/internal/models"
)

type Iusecase interface {
	NewAlbum([]*multipart.FileHeader, models.Album) error
	GetUserByName(string) (models.User, error)
	NewUser(models.User) error
	SignIn(name, password string) error
	GetAlbumById(int) (models.AlbumResponse, error)
	GenerateToken(name, password, role string) (string, error)
	ParseToken(token string, signingKey []byte) (string, error)
	WriteCookie(token string, ctx echo.Context) error
	ReadCookie(ctx echo.Context) (string, error)
}

type usecase struct {
	database database.Idatabase
}

func New(database database.Idatabase) *usecase {
	return &usecase{
		database,
	}
}
