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
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Theme       string `db:"theme"`
	NameAuthor  string `db:"name_author"`
	ImagesCount int    `db:"images_count"`
	Date        string `db:"date"`
	Text        string `db:"text"`
}
