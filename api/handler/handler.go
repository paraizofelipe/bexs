package handler

import (
	"encoding/json"
	"log"
)

type Handler struct {
	Trip Trip
}

type ErrorResponse struct {
	Status int    `json:"-"`
	Error  string `json:"errors"`
}

type SuccessResponse struct {
	Status int    `json:"-"`
	Msg    string `json:"msg"`
}

func (r ErrorResponse) Json() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}

func (r SuccessResponse) Json() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}

func New(filePath string, logger *log.Logger) Handler {
	return Handler{
		Trip: NewTripHandler(filePath, logger),
	}
}
