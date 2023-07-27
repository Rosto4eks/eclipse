package database

import (
	"fmt"

	"github.com/Rosto4eks/eclipse/internal/logger"
	"github.com/Rosto4eks/eclipse/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Idatabase interface {
	GetAllAlbums() ([]models.AlbumResponse, error)
	GetAlbumByID(int) (models.AlbumResponse, error)
	AddAlbum(models.Album) error
	DelAlbumByID(int) error
	AddUser(models.User) error
	DelUser(int) error
	GetUserByName(string) (models.User, error)
	AddArticle(models.Article) error
	GetArticlesByAuthorId(int) ([]models.ArticleResponse, error)
	GetAllArticles() ([]models.ArticleResponse, error)
	GetArticlesById(int) (models.ArticleResponse, error)
	GetThemes() ([]string, error)
	DeleteArticle(int) error
	ChangeArticle(int, string) error
	SearchArticles(string) ([]models.ArticleResponse, error)
	GetComments(int) ([]models.CommentResponse, error)
	GetCommentById(int) (models.CommentResponse, error)
	ChangeComment(int, string) error
	AddComment(models.Comment) (int, error)
	DeleteCommentById(int) error
}

type database struct {
	db     *sqlx.DB
	logger logger.Ilogger
}

func New(db *sqlx.DB, logger logger.Ilogger) *database {
	return &database{
		db:     db,
		logger: logger,
	}
}

func Connect(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.Username, cfg.Database, cfg.Password))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to DB :)")
	return db, nil
}
