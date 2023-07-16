package usecase

import "github.com/Rosto4eks/eclipse/internal/models"

func (u *usecase) GetAllArticles() ([]models.Articles, error) {
	return u.database.GetAllArticles()
}

func (u *usecase) GetArticleByTheme(theme string) ([]models.Articles, error) {
	return u.database.GetArticlesByTheme(theme)
}

func (u *usecase) GetArticlesThemes(articleId int) ([]string, error) {
	return u.database.GetThemesByArticle(articleId)
}

func (u *usecase) NewArticle(article models.Articles) error {
	return u.database.AddArticle(article)
}

func (u *usecase) DeleteArticle(articleId int) error {
	return u.database.DeleteArticle(articleId)
}

/*просмотр статьи по теме (/articles/:theme/:id)
func (u *usecase)*/
