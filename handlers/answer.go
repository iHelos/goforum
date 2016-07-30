package handlers

import (
	"encoding/json"
	"net/http"
	"log"
)

type Answer struct {
	Code     int `json:"code"`
	Response Body `json:"response"`
}

type Body interface{}

func MakeAnswer(code int, response Body) ([]byte, error) {
	ans := Answer{code, response}
	result, err := json.Marshal(ans)
	return result, err
}

func WriteAnswer(w http.ResponseWriter, response Body, status int) {
	w.Header().Set("GoForum", "A Go WebServer for TP database")
	w.Header().Set("Content-Type", "application/json")
	answer, err := MakeAnswer(status, response)
	if err != nil{
		log.Print(err)
	}
	w.Write(answer)
}