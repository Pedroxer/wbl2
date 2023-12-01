package main

import (
	"encoding/json"
)

// ErrorResponse описывает структуру ответа в случае ошибок
type ErrorResponse struct {
	Message string `json:"error"`
}

// NewErrorResponse служит конструктором ErrorResponse
func NewErrorResponse(err error) []byte {
	resp := ErrorResponse{err.Error()}
	jsonResp, _ := json.Marshal(resp)
	return jsonResp
}

// EditResponse описывает структуру ответа для POST запросов
type EditResponse struct {
	Message string `json:"result"`
}

// NewEditResponse служит конструктором EditResponse
func NewEditResponse(msg string) []byte {
	resp := EditResponse{msg}
	jsonResp, _ := json.Marshal(resp)
	return jsonResp
}

// EventsResponse описывает структуру ответа для GET запросов
type EventsResponse struct {
	Data []Event `json:"result"`
}

// NewEventsResponse служит конструктором EventsResponse
func NewEventsResponse(data []Event) []byte {
	resp := EventsResponse{data}
	jsonResp, _ := json.Marshal(resp)
	return jsonResp
}
