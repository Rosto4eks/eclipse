package database

import (
	"github.com/Rosto4eks/eclipse/internal/models"
)

func (d *database) GetComments(articleId int) ([]models.CommentResponse, error) {
	query := "SELECT id, (SELECT name FROM users WHERE id = user_id) AS author, article_id, text, to_char(date, 'HH24:MI YYYY-MM-DD') as date FROM comments WHERE article_id = $1 ORDER BY date DESC"
	var response []models.CommentResponse
	err := d.db.Select(&response, query, articleId)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (d *database) GetCommentById(commentId int) (models.CommentResponse, error) {
	query := "SELECT id, (SELECT name FROM users WHERE id = user_id) AS author, article_id, text, to_char(date, 'HH24:MI YYYY-MM-DD') as date FROM comments WHERE id = $1"
	var response models.CommentResponse
	err := d.db.Get(&response, query, commentId)
	if err != nil {
		return models.CommentResponse{}, err
	}
	return response, nil
}

func (d *database) AddComment(comment models.Comment) (int64, error) {
	query := "INSERT INTO comments (user_id, article_id, text, date) VALUES($1,$2,$3,$4) RETURNING id"
	var id int64
	_ = d.db.QueryRow(query, comment.UserId, comment.ArticleID, comment.Text, comment.Date).Scan(&id)
	return id, nil
}

func (d *database) ChangeComment(comemntId int, newComment string) error {
	query := "UPDATE comments SET text = $2, date = NOW() WHERE id = $1"
	_, err := d.db.Exec(query, comemntId, newComment)
	if err != nil {
		return err
	}
	return nil
}

func (d *database) DeleteCommentById(commentId int) error {
	query := "DELETE FROM comments WHERE id = $1"
	_, err := d.db.Exec(query, commentId)
	if err != nil {
		return err
	}
	return nil
}
