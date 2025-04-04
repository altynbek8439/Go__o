package delivery

import (
	"betting-site/internal/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BetHandler struct {
	service *services.BetService
}

func NewBetHandler(service *services.BetService) *BetHandler {
	return &BetHandler{service: service}
}

func (h *BetHandler) CreateBet(c *gin.Context) {
	type BetRequest struct {
		UserID  int     `json:"user_id"`
		EventID int     `json:"event_id"`
		Amount  float32 `json:"amount"`
		Outcome string  `json:"outcome"`
	}

	var req BetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	log.Printf("Создание ставки: user_id=%d, event_id=%d, amount=%.2f, outcome=%s", req.UserID, req.EventID, req.Amount, req.Outcome)

	newBet, err := h.service.Create(req.UserID, req.EventID, req.Amount, req.Outcome)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bet"})
		return
	}

	c.JSON(http.StatusCreated, newBet)
}

func (h *BetHandler) GetBetsByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	bets, err := h.service.GetBetsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bets"})
		return
	}

	c.JSON(http.StatusOK, bets)
}
