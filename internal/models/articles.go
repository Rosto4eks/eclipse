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

type ArticlesResponse struct {
	ID          int
	Name        string
	Theme       string
	NameAuthor  string
	ImagesCount int
	Date        string
	Text        string
}
