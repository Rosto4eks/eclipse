package models

type Comments struct {
	ID        int
	UserId    int
	ArticleID int
	Text      string
	Date      string
}
