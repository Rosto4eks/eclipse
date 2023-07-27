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
		d.logger.Error("database", "AddArticle", err)
		return err
	}
	d.logger.Info("database", "AddArticle", fmt.Sprintf("added %s article", articles.Name))
	return nil
}

// прсмотр статей по автору
func (d *database) GetArticlesByAuthorId(authorId int) ([]models.ArticleResponse, error) {
	query := "SELECT id, name, theme, (SELECT name FROM users WHERE id = author_id) AS name_author, images_count, to_char(date,'YYYY-MM-DD') AS date, text FROM articles WHERE author_id = $1"
	var response []models.ArticleResponse
	err := d.db.Select(&response, query, authorId)
	if err != nil {
		d.logger.Error("database", "GetArticlesByAuthorId", err)
		return nil, err
	}
	return response, nil
}

// просмотр всех статей
func (d *database) GetArticles(offset, limit int) ([]models.ArticleResponse, error) {
	query := "SELECT id, name, theme, (SELECT name FROM users WHERE id = author_id) AS name_author, images_count, to_char(date,'YYYY-MM-DD') AS date, text FROM articles ORDER BY date DESC OFFSET $1 LIMIT $2"
	var response []models.ArticleResponse
	err := d.db.Select(&response, query, offset, limit)
	if err != nil {
		d.logger.Error("database", "GetArticles", err)
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
		d.logger.Error("database", "GetThemes", err)
		return nil, err
	}
	return response, nil
}

// просмотр списка статей по выбранной теме
func (d *database) GetArticlesById(articleId int) (models.ArticleResponse, error) {
	query := "SELECT id, name, theme, (SELECT name FROM users where id = author_id) as name_author, images_count, to_char(date,'YYYY-MM-DD') AS date, text FROM articles WHERE id = $1"
	var response models.ArticleResponse
	err := d.db.Get(&response, query, articleId)
	if err != nil {
		d.logger.Error("database", "GetArticlesById", err)
		return models.ArticleResponse{}, err
	}
	return response, nil
}

func (d *database) ChangeArticle(articleId int, newText string) error {
	query := "UPDATE articles SET text = $2 WHERE id = $1"
	_, err := d.db.Exec(query, articleId, newText)
	if err != nil {
		d.logger.Error("database", "ChangeArticle", err)
		return err
	}
	d.logger.Info("database", "ChangeArticle", fmt.Sprintf("article with id = %d changed", articleId))
	return nil
}

// удаление статьи по ее id
func (d *database) DeleteArticle(articleId int) error {
	query := "DELETE FROM articles WHERE id = $1"
	_, err := d.db.Exec(query, articleId)
	if err != nil {
		d.logger.Error("database", "DeleteArticle", err)
		return err
	}
	d.logger.Info("database", "DeleteArticle", fmt.Sprintf("article with id = %d deleted", articleId))
	return nil
}

func (d *database) SearchArticles(value string) ([]models.ArticleResponse, error) {
	query := fmt.Sprintf("SELECT id, name, theme, (SELECT name FROM users where id = author_id) as name_author, images_count, to_char(date,'YYYY-MM-DD') AS date, text FROM articles WHERE LOWER(name) LIKE LOWER('%%%s%%')", value)
	var response []models.ArticleResponse
	err := d.db.Select(&response, query)
	if err != nil {
		d.logger.Error("database", "SearchArticles", err)
		return nil, err
	}
	return response, nil
}
