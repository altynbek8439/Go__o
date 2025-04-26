package repository

import (
	"betting-site/internal/models"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Create(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *EventRepository) GetAll() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *EventRepository) GetByID(id int) (*models.Event, error) {
	var event models.Event
	err := r.db.First(&event, id).Error
	return &event, err
}

func (r *EventRepository) Update(id int, eventEdit *models.EventEdit) error {
	return r.db.Model(&models.Event{}).Where("id = ?", id).Omit("id, CreatedAt").Updates(eventEdit).Error
}

func (r *EventRepository) Delete(eventID int) error {
	return r.db.Delete(&models.Event{}, eventID).Error
}
