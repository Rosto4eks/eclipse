package usecase

import (
	"fmt"
	"image"
	"io"
	"mime/multipart"
	"os"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"

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

func (u *usecase) GetAlbumById(id int) (models.AlbumResponse, error) {
	return u.database.GetAlbumByID(id)
}

func (u *usecase) GetAllAlbums() ([]models.AlbumResponse, error) {
	return u.database.GetAllAlbums()
}

func (u *usecase) DeleteAlbum(id int) error {
	album, err := u.database.GetAlbumByID(id)
	if err != nil {
		return err
	}
	path := "public/albums/" + album.Date + "-" + album.Name
	fmt.Println(path)
	err = os.RemoveAll(path)
	if err != nil {
		return err
	}
	return u.database.DelAlbumByID(id)
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
			// save original image
			if err := saveImage(src, path, i); err != nil {
				errChan <- err
				return
			}

			src.Seek(0, io.SeekStart)
			// comress image and save it
			if err := compressAndSaveImage(src, path, i); err != nil {
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

func saveImage(src multipart.File, path string, i int) error {
	dst, err := os.Create(fmt.Sprintf("%s/%d.jpeg", path, i))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func compressAndSaveImage(src multipart.File, path string, i int) error {
	img, err := imaging.Decode(src)
	if err != nil {
		return err
	}

	src.Seek(0, 0)
	ex, err := exif.Decode(src)
	if err != nil {
		return err
	}

	if orient, err := ex.Get(exif.Orientation); orient != nil && err == nil {
		img = reverse(img, orient.String())
	} else if err != nil {
		return err
	}

	newimg := imaging.Resize(img, img.Bounds().Dx()/8, img.Bounds().Dy()/8, imaging.Lanczos)

	dst, err := os.Create(fmt.Sprintf("%s/%d-compressed.jpeg", path, i))
	if err != nil {
		return err
	}
	defer dst.Close()

	return imaging.Encode(dst, newimg, imaging.JPEG)
}

func reverse(img image.Image, orient string) image.Image {
	switch orient {
	case "1":
		return img
	case "2":
		return imaging.FlipV(img)
	case "3":
		return imaging.Rotate180(img)
	case "4":
		return imaging.Rotate180(imaging.FlipV(img))
	case "5":
		return imaging.Rotate270(imaging.FlipV(img))
	case "6":
		return imaging.Rotate270(img)
	case "7":
		return imaging.Rotate90(imaging.FlipV(img))
	case "8":
		return imaging.Rotate90(img)
	}
	return img
}
