package models

type Comment struct {
	ID        int
	UserId    int
	ArticleID int
	Text      string
	Date      string
}

type CommentResponse struct {
	ID         int    `db:"id"`
	AuthorName string `db:"author"`
	ArticleID  int    `db:"article_id"`
	Text       string `db:"text"`
	Date       string `db:"date"`
}
