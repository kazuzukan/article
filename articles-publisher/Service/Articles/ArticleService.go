package Articles

import (
	"articles-publisher/Controller/Dto/Request"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	EndpointGetArticles = "%s/articles/list"
)

type httpResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
}

type ArticleServiceInterface interface {
	getHost() (host string)
	GetArticleList(dataRequest Request.GetArticles) (res httpResponse, err error)
}
type ArticleService struct {
}

func NewArticleService() ArticleService {
	return ArticleService{}
}

func (a ArticleService) getHost() (host string) {
	host = os.Getenv("CONSUMER_HOST")
	if host == "" {
		host = "http://localhost:9000"
	}

	return
}

func (a ArticleService) GetArticleList(dataRequest Request.GetArticles) (res httpResponse, err error) {
	header := make(map[string]string)
	header["Content-type"] = "application/json"
	jsonByte, _ := json.Marshal(dataRequest)

	url := fmt.Sprintf(EndpointGetArticles, a.getHost())
	responseData, err := a.newRequest(http.MethodPost, url, header, bytes.NewReader(jsonByte))
	if err != nil {
		return
	}

	err = json.Unmarshal(responseData, &res)
	if err != nil {
		return
	}

	if res.Code != http.StatusOK {
		err = errors.New(res.Data.(string))
		return
	}

	return
}

func (a ArticleService) newRequest(method, url string, key map[string]string, body io.Reader) (responseByte json.RawMessage, err error) {
	transportConfig := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	timeout := time.Second * 30
	client := http.Client{Timeout: timeout, Transport: transportConfig}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Close = true

	if key != nil {
		for key, value := range key {
			request.Header.Set(key, value)
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseByte)
	if err != nil {
		return nil, err
	}

	return
}
