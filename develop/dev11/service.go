package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Service предоставляет интерфейс для слоя бизнес-логики приложения
type Service struct {
	Cache *Cache
}

// NewService служит конструктором Service
func NewService() *Service {
	return &Service{NewCache()}
}

// CreateEvent создает новое событие
func (s *Service) CreateEvent(data []byte) error {
	var e Event
	err := json.Unmarshal(data, &e)
	if err != nil {
		return err
	}

	s.Cache.Add(e)
	return nil
}

// UpdateEvent обновляет событие
func (s *Service) UpdateEvent(data []byte) error {
	var e Event
	err := json.Unmarshal(data, &e)
	if err != nil {
		return err
	}

	if !s.Cache.Contains(e) {
		return fmt.Errorf("event id %s not found", e.ID)
	}

	s.Cache.Update(e)
	return nil
}

// DeleteEvent удаляет событие
func (s *Service) DeleteEvent(data []byte) error {
	var e Event
	err := json.Unmarshal(data, &e)
	if err != nil {
		return err
	}

	if !s.Cache.Contains(e) {
		return fmt.Errorf("event id %s not found", e.ID)
	}

	s.Cache.Delete(e)
	return nil
}

// GetEventsDay возвращает события дня из кэша
func (s *Service) GetEventsDay(params url.Values) ([]byte, error) {
	date, err := time.Parse("2006-01-02", params.Get("date"))
	if err != nil {
		return nil, err
	}

	events := s.Cache.GetByDate(date)
	jsonRes := NewEventsResponse(events)

	return jsonRes, nil
}

// GetEventsWeek возвращает события недели из кэша
func (s *Service) GetEventsWeek(params url.Values) ([]byte, error) {
	date, err := time.Parse("2006-01-02", params.Get("date"))
	if err != nil {
		return nil, err
	}

	events := s.Cache.GetByWeek(date)
	jsonRes := NewEventsResponse(events)

	return jsonRes, nil
}

// GetEventsMonth возвращает события месяца из кэша
func (s *Service) GetEventsMonth(params url.Values) ([]byte, error) {
	date, err := time.Parse("2006-01", params.Get("date"))
	if err != nil {
		return nil, err
	}

	events := s.Cache.GetByMonth(date)
	jsonRes := NewEventsResponse(events)

	return jsonRes, nil
}
