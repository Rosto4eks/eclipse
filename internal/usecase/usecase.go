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
	GetAllAlbums() ([]models.AlbumResponse, error)
	DeleteAlbum(int) error
	GetAllArticles() ([]models.ArticleResponse, error)
	GetArticleByTheme(string) ([]models.ArticleResponse, error)
	NewArticle(models.Article) error
	GetThemes() ([]string, error)
	DeleteArticle(int) error
}

type usecase struct {
	database database.Idatabase
}

func New(database database.Idatabase) *usecase {
	return &usecase{
		database,
	}
}
