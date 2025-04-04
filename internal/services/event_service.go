package services

import (
	"betting-site/internal/models"
	"betting-site/internal/repository"
)

type EventService struct {
	repo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) Create(name, date string, oddsWin1, oddsDraw, oddsWin2 float32) (*models.Event, error) {
	event := &models.Event{
		Name:     name,
		Date:     date,
		OddsWin1: oddsWin1,
		OddsDraw: oddsDraw,
		OddsWin2: oddsWin2,
	}
	err := s.repo.Create(event)
	return event, err
}

func (s *EventService) GetAll() ([]models.Event, error) {
	return s.repo.GetAll()
}
