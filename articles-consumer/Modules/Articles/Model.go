package Articles

import (
	"articles-consumer/Controller/Dto/Request"
	"articles-consumer/Controller/Dto/Response"
	"articles-consumer/Repository"
)

type ArticleModule interface {
	CreateArticle(dataRequest Request.CreateArticle) (err error)
	GetArticlesList(dataRequest Request.GetArticles) (articles Response.Articles, err error)
}

type article struct {
	repository Repository.Repository
}

func NewModule(repo Repository.Repository) ArticleModule {
	return &article{
		repository: repo,
	}
}
