package models

type Album struct {
	Id           int
	Author_id    int
	Images_count int
	Date         string
	Name         string
	Description  string
}

type AlbumResponse struct {
	Id           int
	Author       string
	Images_count int
	Date         string
	Name         string
	Description  string
}
