package usecase

import "github.com/Rosto4eks/eclipse/internal/models"

func (u *usecase) GetAllArticles() ([]models.ArticleResponse, error) {
	return u.database.GetAllArticles()
}

func (u *usecase) GetArticleById(articleId int) (models.ArticleResponse, error) {
	return u.database.GetArticlesById(articleId)
}

func (u *usecase) GetThemes() ([]string, error) {
	return u.database.GetThemes()
}

func (u *usecase) NewArticle(article models.Article) error {
	return u.database.AddArticle(article)
}

func (u *usecase) DeleteArticle(articleId int) error {
	return u.database.DeleteArticle(articleId)
}

func (u *usecase) ChangeArticle(articleId int, newText string) error {
	return u.database.ChangeArticle(articleId, newText)
}
