package Articles

import (
	"articles-consumer/Controller/Dto"
	"articles-consumer/Controller/Dto/Request"
	"articles-consumer/Helper"
	"encoding/json"
	"net/http"
)

type ControllerAPI interface {
	GetArticlesList(w http.ResponseWriter, r *http.Request)
}

func NewControllerAPI(u Dto.Utilities) ControllerAPI {
	return ArticleAPI{u}
}

func (a ArticleAPI) GetArticlesList(w http.ResponseWriter, r *http.Request) {
	var dataRequest Request.GetArticles
	err := json.NewDecoder(r.Body).Decode(&dataRequest)
	if err != nil {
		Helper.HttpResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	articlesList, err := a.Modules.Articles.GetArticlesList(dataRequest)
	if err != nil {
		Helper.HttpResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Helper.HttpResponseSuccess(w, articlesList)
}
