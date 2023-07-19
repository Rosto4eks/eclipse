package usecase

import "github.com/Rosto4eks/eclipse/internal/models"

func (u *usecase) GetArticleComments(articleId int) ([]models.CommentResponse, error) {
	return u.database.GetComments(articleId)
}

func (u *usecase) GetCommentById(commentId int) (models.CommentResponse, error) {
	return u.database.GetCommentById(commentId)
}

func (u *usecase) DeleteComment(commentId int) error {
	return u.database.DeleteCommentById(commentId)
}
