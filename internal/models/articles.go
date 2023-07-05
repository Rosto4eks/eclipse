package models

type Articles struct {
	ID          int
	Name        string
	Theme       string
	AuthorID    int
	ImagesCount int
	Date        string
	Text        string
}
