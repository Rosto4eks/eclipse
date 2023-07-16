package models

type Comment struct {
	ID        int
	UserId    int
	ArticleID int
	Text      string
	Date      string
}

type CommentResponse struct {
	ID         int
	AuthorName string
	ArticleID  int
	Text       string
	Date       string
}
