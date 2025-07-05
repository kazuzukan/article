package Articles

import (
	"articles-consumer/Config"
	"articles-consumer/Controller/Dto/Request"
	"articles-consumer/Controller/Dto/Response"
)

type Repository interface {
	CreateArticle(dataRequest Request.CreateArticle) (err error)
	GetArticles(params Request.GetArticles) (articles Response.Articles, err error)
}

type article struct {
	dbCon Config.ConnectionInterface
}

func NewRepository(dbCon Config.ConnectionInterface) Repository {
	return article{
		dbCon: dbCon,
	}
}
