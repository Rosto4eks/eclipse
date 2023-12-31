package usecase

import (
	"mime/multipart"

	"github.com/Rosto4eks/eclipse/internal/database"
	"github.com/Rosto4eks/eclipse/internal/logger"
	"github.com/Rosto4eks/eclipse/internal/models"
)

type Iusecase interface {
	NewAlbum([]*multipart.FileHeader, models.Album) error
	GetUserByName(string) (models.User, error)
	SignUp(models.User) (string, error)
	SignIn(name, password string) (string, error)
	GetAlbumById(int) (models.AlbumResponse, error)
	Auth(string, string) error
	AuthHeader(string) string
	GetAlbums(offset, count int) ([]models.AlbumResponse, error)
	DeleteAlbum(int) error
	GetArticles(int, int) ([]models.ArticleResponse, error)
	GetArticleById(int) (models.ArticleResponse, error)
	NewArticle([]*multipart.FileHeader, models.Article) error
	GetThemes() ([]string, error)
	DeleteArticle(int) error
	GetArticleComments(articleId int) ([]models.CommentResponse, error)
	ChangeArticle(articleId int, newText string) error
	SearchArticle(string) ([]models.ArticleResponse, error)
	GetCommentById(commentId int) (models.CommentResponse, error)
	AddNewComment(comment models.Comment) (int, error)
	DeleteComment(commentId int) error
	ChangeComment(commentId int, newText string) error
}

type usecase struct {
	database database.Idatabase
	logger   logger.Ilogger
}

func New(database database.Idatabase, logger logger.Ilogger) *usecase {
	return &usecase{
		database,
		logger,
	}
}
