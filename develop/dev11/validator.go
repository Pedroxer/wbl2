package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Validator предоставляет интерфейс валидации запросов
type Validator struct{}

// NewValidator - конструктор Validator'a
func NewValidator() *Validator {
	return &Validator{}
}

// ValidateEventCreate возвращает ошибку, если данные нельзя десериализовать в нужный объект
func (v *Validator) ValidateEventCreate(data []byte) error {
	var e Event
	err := json.Unmarshal(data, &e)
	if err != nil {
		return err
	}
	return nil
}

// ValidateEventID возвращает ошибку, если данные нельзя десериализовать в нужный объект, либо не указан id события
func (v *Validator) ValidateEventID(data []byte) error {
	var e Event
	err := json.Unmarshal(data, &e)
	if err != nil {
		return err
	}
	if e.ID == "" {
		return fmt.Errorf("event id is not provided")
	}
	return nil
}

// ValidateDate возвращает ошибку, если дата не указана в формате "2006-01-02"
func (v *Validator) ValidateDate(params url.Values) error {
	if !params.Has("date") {
		return fmt.Errorf("no date provided")
	}

	_, err := time.Parse("2006-01-02", params.Get("date"))
	if err != nil {
		return err
	}
	return nil
}

// ValidateMonth возвращает ошибку, если дата не указана в формате "2006-01"
func (v *Validator) ValidateMonth(params url.Values) error {
	if !params.Has("date") {
		return fmt.Errorf("no date provided")
	}

	_, err := time.Parse("2006-01", params.Get("date"))
	if err != nil {
		return err
	}
	return nil
}
