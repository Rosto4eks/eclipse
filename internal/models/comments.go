package models

import "time"

type Comments struct {
	ID        int
	UserId    int
	ArticleID int
	Text      string
	Date      time.Time
}
