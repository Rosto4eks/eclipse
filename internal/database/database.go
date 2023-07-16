package database

import (
	"fmt"

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
	AddArticle(articles models.Article) error
	GetArticlesByAuthorId(int) ([]models.ArticleResponse, error)
	GetAllArticles() ([]models.ArticleResponse, error)
	GetArticlesByTheme(string) ([]models.ArticleResponse, error)
	GetThemes() ([]string, error)
	DeleteArticle(int) error
	AddComment(models.Comment) error
	DeleteCommentById(int) error
}

type database struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *database {
	return &database{db: db}
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
