package delivery

import (
	"betting-site/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service *services.EventService
}

func NewEventHandler(service *services.EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	type EventRequest struct {
		Name     string  `json:"name"`
		Date     string  `json:"date"`
		OddsWin1 float32 `json:"odds_win1"`
		OddsDraw float32 `json:"odds_draw"`
		OddsWin2 float32 `json:"odds_win2"`
	}

	var req EventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Логируем входящие данные
	log.Printf("Создание события: name=%s, date=%s, odds_win1=%.2f, odds_draw=%.2f, odds_win2=%.2f", req.Name, req.Date, req.OddsWin1, req.OddsDraw, req.OddsWin2)

	newEvent, err := h.service.Create(req.Name, req.Date, req.OddsWin1, req.OddsDraw, req.OddsWin2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, newEvent)
}

func (h *EventHandler) GetEvents(c *gin.Context) {
	events, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}

	c.JSON(http.StatusOK, events)
}
