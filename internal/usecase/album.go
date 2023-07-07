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
	if err := u.database.AddAlbum(album); err != nil {
		return err
	}
	return nil
}

func (u *usecase) saveAlbumImages(files []*multipart.FileHeader, album models.Album) error {
	path := fmt.Sprintf("./public/albums/%s-%s", album.Date, album.Name)
	// create destination folder
	os.Mkdir(path, os.ModePerm)
	errChan := make(chan error)
	counter := 0
	for i, file := range files {
		go func(i, max int, counter *int, file *multipart.FileHeader, errChan chan<- error) {
			defer func() {
				// count all perfomed gorutines
				*counter++
				if *counter == max+1 {
					errChan <- nil
				}
			}()
			src, err := file.Open()
			if err != nil {
				errChan <- err
				return
			}
			defer src.Close()

			// destination file
			dst, err := os.Create(fmt.Sprintf("%s/%d.jpg", path, i))
			if err != nil {
				errChan <- err
				return
			}
			defer dst.Close()

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				errChan <- err
				return
			}
		}(i, len(files)-1, &counter, file, errChan)
	}
	if err := <-errChan; err != nil {
		return err
	}
	close(errChan)
	return nil
}
