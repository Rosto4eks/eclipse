package models

import "time"

type Articles struct {
	ID          int
	Name        string
	Theme       string
	AuthorID    int
	ImagesCount int
	Date        time.Time
	Text        string
}
