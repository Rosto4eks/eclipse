package database

import (
	"fmt"
	"github.com/Rosto4eks/eclipse/internal/models"
)

func (d *database) GetComments(articleId int) ([]models.CommentResponse, error) {
	query := "SELECT (SELECT name FROM users WHERE id = user_id) as author, text, to_char(date, 'YYYY-MM-DD') as date FROM comments WHERE article_id = $1"
	var response []models.CommentResponse
	err := d.db.Select(response, query, articleId)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (d *database) AddComment(comment models.Comment) error {
	query := "INSERT INTO comments (userId, ArticleId, Text, Date) VALUES($1,$2,$3,$4)"
	_, err := d.db.Exec(query, comment.UserId, comment.ArticleID, comment.Text, comment.Date)
	if err != nil {
		return err
	}
	return nil
}

func (d *database) ChangeComment(userId, articleId int, newComment string) error {
	query := "UPDATE comments SET text = $3 WHERE user_id = $1 AND id = $2"
	result, err := d.db.Exec(query, userId, articleId, newComment)
	if err != nil {
		return err
	}
	fmt.Println(result)
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
