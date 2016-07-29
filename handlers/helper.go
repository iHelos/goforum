package handlers

import (
	"encoding/json"
)

type Answer struct {
	code     int
	response Body
}

type Body interface{}

func getAnswer(code int, response Body) ([]byte, error) {
	ans := Answer{code, response}
	result, err := json.Marshal(ans)
	return result, err
}