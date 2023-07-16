package database

import (
	"fmt"
	"github.com/Rosto4eks/eclipse/internal/models"
)

// добавить статью
func (d *database) AddArticle(articles models.Article) error {
	query := "INSERT INTO articles (name, theme, author_id, images_count, date, text) VALUES($1,$2,$3,$4,$5,$6)"
	_, err := d.db.Exec(query, articles.Name, articles.Theme, articles.AuthorID, articles.ImagesCount, articles.Date, articles.Text)
	if err != nil {
		return err
	}
	return nil
}

// прсмотр статей по автору
func (d *database) GetArticlesByAuthorId(authorId int) ([]models.ArticleResponse, error) {
	query := "SELECT name, theme, (SELECT name FROM users WHERE id = author_id) as name_author, images_count, to_char(date,'YYYY-MM-DD') AS date, text FROM articles WHERE author_id = $1"
	var response []models.ArticleResponse
	err := d.db.Select(&response, query, authorId)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// просмотр всех статей
func (d *database) GetAllArticles() ([]models.ArticleResponse, error) {
	query := "SELECT name, theme, (SELECT name FROM users) as name_author, images_count, to_char(date,'YYYY-MM-DD') AS date, text FROM articles"
	var response []models.ArticleResponse
	err := d.db.Select(&response, query)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// просмотр списка тем статей
func (d *database) GetThemes() ([]string, error) {
	query := "SELECT DISTINCT theme FROM articles"
	var response []string
	err := d.db.Select(&response, query)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// просмотр списка статей по выбранной теме
func (d *database) GetArticlesByTheme(theme string) ([]models.ArticleResponse, error) {
	query := "SELECT name, theme, (SELECT name FROM users where id = author_id) as name_author, images_count, to_char(date,'YYYY-MM-DD') AS date, text FROM articles WHERE theme = $1"
	var response []models.ArticleResponse
	err := d.db.Select(&response, query, theme)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// удаление статьи по ее id
func (d *database) DeleteArticle(articleId int) error {
	query := "DELETE FROM articles WHERE article_id = $1"
	result, err := d.db.Exec(query, articleId)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
