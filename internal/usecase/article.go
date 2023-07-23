package usecase

import (
	"fmt"
	"mime/multipart"
	"os"
	"regexp"

	"github.com/Rosto4eks/eclipse/internal/models"
)

func (u *usecase) GetAllArticles() ([]models.ArticleResponse, error) {
	return u.database.GetAllArticles()
}

func (u *usecase) GetArticleById(articleId int) (models.ArticleResponse, error) {
	return u.database.GetArticlesById(articleId)
}

func (u *usecase) GetThemes() ([]string, error) {
	return u.database.GetThemes()
}

func (u *usecase) NewArticle(files []*multipart.FileHeader, article models.Article) error {
	u.changeSrcs(&article, len(files)-1)
	if err := u.database.AddArticle(article); err != nil {
		return err
	}
	return u.saveArticleImages(files, article)
}

func (u *usecase) DeleteArticle(articleId int) error {
	article, err := u.database.GetArticlesById(articleId)
	if err != nil {
		return err
	}
	path := "public/articles/" + article.Date + "-" + article.Name
	err = os.RemoveAll(path)
	if err != nil {
		return err
	}
	return u.database.DeleteArticle(articleId)
}

func (u *usecase) changeSrcs(article *models.Article, count int) {
	for i := 0; i <= count; i++ {
		regex := regexp.MustCompile(fmt.Sprintf("img-src-%d", i))
		article.Text = regex.ReplaceAllString(article.Text, fmt.Sprintf("/public/articles/%s-%s/%d.jpeg", article.Date, article.Name, i))
	}
}

func (u *usecase) saveArticleImages(files []*multipart.FileHeader, article models.Article) error {
	path := fmt.Sprintf("./public/articles/%s-%s", article.Date, article.Name)
	// create destination folder
	os.Mkdir(path, os.ModePerm)
	errChan := make(chan error)

	for i, file := range files {
		go func(i, max int, file *multipart.FileHeader, errChan chan<- error) {
			defer func() {
				if i == max {
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
			if err := saveImage(src, path, i, i == max); err != nil {
				errChan <- err
				return
			}

		}(i, len(files)-1, file, errChan)
	}
	if err := <-errChan; err != nil {
		return err
	}
	close(errChan)
	return nil
}

func (u *usecase) ChangeArticle(articleId int, newText string) error {
	return u.database.ChangeArticle(articleId, newText)
}
