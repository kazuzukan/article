package Articles

import (
	"articles-publisher/Config"
	"articles-publisher/Controller/Dto"
	"articles-publisher/Controller/Dto/Request"
	"articles-publisher/Helper"
	"encoding/json"
	"net/http"
	"strconv"
)

type ArticlesControllerInterface interface {
	CreateArticles(w http.ResponseWriter, r *http.Request)
	GetArticlesList(w http.ResponseWriter, r *http.Request)
}

func NewController(u Dto.Utilities) Atricles {
	return Atricles{u}
}

func (a Atricles) CreateArticles(w http.ResponseWriter, r *http.Request) {
	var dataRequest Request.CreateArticles
	err := json.NewDecoder(r.Body).Decode(&dataRequest)
	if err != nil {
		Helper.HttpResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg, err := a.RabbitMQ.PublishMessageRPC(Config.ArticlesBindingKey, Config.ArticlesExchange, dataRequest)
	if err != nil {
		Helper.HttpResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respRPC := Helper.GetResponseFromRPCReply(msg)
	if respRPC.Code != http.StatusOK {
		Helper.HttpResponseError(w, respRPC.Status, respRPC.Code)
		return
	}

	Helper.HttpResponseSuccess(w, respRPC.Data)
}

func (a Atricles) GetArticlesList(w http.ResponseWriter, r *http.Request) {
	var dataRequest Request.GetArticles
	var err error

	dataRequest.Page, err = strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		Helper.HttpResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if dataRequest.Page <= 0 {
		Helper.HttpResponseError(w, nil, http.StatusBadRequest)
		return
	}

	dataRequest.Keyword = r.URL.Query().Get("keyword")
	dataRequest.AuthorName = r.URL.Query().Get("author")
	articleList, err := a.ArticlesService.GetArticleList(dataRequest)
	if err != nil {
		Helper.HttpResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Helper.HttpResponseSuccess(w, articleList.Data)
}
