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
	GetArticlesById(int) (models.ArticleResponse, error)
	GetThemes() ([]string, error)
	DeleteArticle(int) error
	ChangeArticle(articleId int, newText string) error
	GetComments(articleId int) ([]models.CommentResponse, error)
	GetCommentById(commentId int) (models.CommentResponse, error)
	ChangeComment(comemntId int, newComment string) error
	AddComment(models.Comment) (int, error)
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
