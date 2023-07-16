package usecase

import "github.com/Rosto4eks/eclipse/internal/models"

func (u *usecase) GetAllArticles() ([]models.ArticleResponse, error) {
	return u.database.GetAllArticles()
}

func (u *usecase) GetArticleByTheme(theme string) ([]models.ArticleResponse, error) {
	return u.database.GetArticlesByTheme(theme)
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
