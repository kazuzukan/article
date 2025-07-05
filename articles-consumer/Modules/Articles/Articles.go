package Articles

import (
	"articles-consumer/Controller/Dto/Request"
	"articles-consumer/Controller/Dto/Response"
)

func (a article) CreateArticle(dataRequest Request.CreateArticle) (err error) {
	err = a.repository.Articles.CreateArticle(dataRequest)
	return
}

func (a article) GetArticlesList(dataRequest Request.GetArticles) (articles Response.Articles, err error) {
	articles, err = a.repository.Articles.GetArticles(dataRequest)
	return
}
