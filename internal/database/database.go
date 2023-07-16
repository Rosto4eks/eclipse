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
	AddArticle(articles models.Articles) error
	GetArticlesById(int) ([]models.Articles, error)
	GetAllArticles() ([]models.Articles, error)
	GetThemesByArticle(int) ([]string, error)
	GetArticlesByTheme(string) ([]models.Articles, error)
	GetThemes() ([]string, error)
	DeleteArticle(int) error
	AddComment(userId, articleId int, comment string) error
	DeleteCommentByUser(userId, articleId int, comment string) error
	DeleteCommentsByEditor(articleId int, comment string) error
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
