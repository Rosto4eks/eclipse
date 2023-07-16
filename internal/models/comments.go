package models

type Comment struct {
	ID        int
	UserId    int
	ArticleID int
	Text      string
	Date      string
}
