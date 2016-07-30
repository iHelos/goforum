package common

import (
	"github.com/iHelos/goforum/handlers"
	"net/http"
)

const (
	UriStatus = "status"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	body, status := status()
	handlers.WriteAnswer(w, body, status)
}

func status() (interface{}, int){
	return nil, 3
}