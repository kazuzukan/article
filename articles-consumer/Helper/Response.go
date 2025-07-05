package Helper

import (
	"encoding/json"
	"net/http"
	"time"
)

type ResponsePublisher struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Code       int         `json:"code"`
	AccessTime string      `json:"accessTime"`
}

func SetResponsePublisher(status string, data interface{}, code int) ResponsePublisher {
	responseData := ResponsePublisher{
		Status: status,
		Data:   data,
		Code:   code,
	}

	return responseData
}

func HttpResponseSuccess(w http.ResponseWriter, data interface{}) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	responseData := ResponsePublisher{
		Code:       http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Data:       data,
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05"),
	}

	response, _ := json.Marshal(responseData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func HttpResponseError(w http.ResponseWriter, data interface{}, code int) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	setResponse := ResponsePublisher{
		Status:     http.StatusText(code),
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05"),
		Data:       data,
		Code:       code,
	}

	response, _ := json.Marshal(setResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetResponseFromRPCReply(message []byte) (response ResponsePublisher) {
	json.Unmarshal(message, &response)
	return
}
