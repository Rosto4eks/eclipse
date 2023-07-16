package models

type Article struct {
	ID          int
	Name        string
	Theme       string
	AuthorID    int
	ImagesCount int
	Date        string
	Text        string
}

type ArticleResponse struct {
	ID          int
	Name        string
	Theme       string
	NameAuthor  string
	ImagesCount int
	Date        string
	Text        string
}
