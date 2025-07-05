package main

import (
	"articles-publisher/Controller"
	"articles-publisher/Controller/Dto/Request"
	"articles-publisher/Controller/Dto/Response"
	"articles-publisher/Service"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// POSITIVE CASE
func TestGetArticlesListDefault(t *testing.T) {
	utilities := Service.GenerateUtilities()
	controller := Controller.InitController(utilities)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/article/list?page=1", nil)
	controller.Articles.GetArticlesList(recorder, request)

	result := recorder.Result()
	if result.StatusCode != http.StatusOK {
		log.Fatal("ERROR")
	}

	// check length of page
	var resData Response.Articles
	json.NewDecoder(result.Body).Decode(&resData)
	if len(resData.ArticlesList) >= 5 {
		log.Fatal("ERROR Data Length")
	}
}

func TestGetArticlesListFilterAndKeyword(t *testing.T) {
	utilities := Service.GenerateUtilities()
	controller := Controller.InitController(utilities)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/article/list?page=1&author=Ripang2&keyword=Niat", nil)
	controller.Articles.GetArticlesList(recorder, request)

	result := recorder.Result()
	if result.StatusCode != http.StatusOK {
		log.Fatal("ERROR")
	}
}

func TestCreateArticles(t *testing.T) {
	utilities := Service.GenerateUtilities()
	controller := Controller.InitController(utilities)

	bodyReq := Request.CreateArticles{
		AuthorName: "Asta",
		Title:      "Beruang Kutub Datang",
		Body:       "Berugan Ditemukan Di rumah Warga",
	}

	content, err := json.Marshal(bodyReq)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/article/create", bytes.NewBuffer(content))
	controller.Articles.CreateArticles(recorder, request)

	result := recorder.Result()
	var resData map[string]interface{}
	json.NewDecoder(result.Body).Decode(&resData)

	if result.StatusCode != http.StatusOK {
		log.Fatal("ERROR")
	}
}

// NEGATIVE CASE
func TestGetArticlesListPageIsZero(t *testing.T) {
	utilities := Service.GenerateUtilities()
	controller := Controller.InitController(utilities)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/article/list?page=0", nil)
	controller.Articles.GetArticlesList(recorder, request)

	result := recorder.Result()
	if result.StatusCode != http.StatusBadRequest {
		log.Fatal("This Scenario Should Fail")
	}
}
