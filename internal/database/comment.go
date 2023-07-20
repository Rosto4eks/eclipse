package database

import (
	"fmt"
	"github.com/Rosto4eks/eclipse/internal/models"
)

func (d *database) GetComments(articleId int) ([]models.CommentResponse, error) {
	query := "SELECT id, (SELECT name FROM users WHERE id = user_id) AS author, article_id, text, to_char(date, 'HH24:MI YYYY-MM-DD') as date FROM comments WHERE article_id = $1 ORDER BY date DESC"
	var response []models.CommentResponse
	err := d.db.Select(&response, query, articleId)
	fmt.Println(response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (d *database) GetCommentById(commentId int) (models.CommentResponse, error) {
	query := "SELECT id, (SELECT name FROM users WHERE id = user_id) AS author, article_id, text, to_char(date, 'HH24:MI YYYY-MM-DD') as date FROM comments WHERE id = $1"
	var response models.CommentResponse
	err := d.db.Get(&response, query, commentId)
	fmt.Println(response)
	if err != nil {
		return models.CommentResponse{}, err
	}
	return response, nil
}

func (d *database) AddComment(comment models.Comment) error {
	query := "INSERT INTO comments (user_id, article_id, text, date) VALUES($1,$2,$3,$4)"
	_, err := d.db.Exec(query, comment.UserId, comment.ArticleID, comment.Text, comment.Date)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (d *database) ChangeComment(comemntId int, newComment string) error {
	query := "UPDATE comments SET text = $2, date = NOW() WHERE comment_id = $1"
	result, err := d.db.Exec(query, comemntId, newComment)
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
