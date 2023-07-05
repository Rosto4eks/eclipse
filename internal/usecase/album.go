package usecase

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/Rosto4eks/eclipse/internal/models"
)

func (u *usecase) NewAlbum(files []*multipart.FileHeader, album models.Album) error {
	if err := u.saveAlbumImages(files, album); err != nil {
		return err
	}
	// write to db
	return nil
}

func (u *usecase) saveAlbumImages(files []*multipart.FileHeader, album models.Album) error {
	path := fmt.Sprintf("./public/albums/%s-%s", album.Date, album.Name)
	// create destination folder
	os.Mkdir(path, os.ModePerm)

	for i, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// destination file
		dst, err := os.Create(fmt.Sprintf("%s/%d.jpg", path, i))
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}
	return nil
}
