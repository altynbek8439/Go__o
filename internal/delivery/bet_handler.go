package delivery

import (
	"betting-site/internal/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BetHandler struct {
	service *services.BetService
}

func NewBetHandler(service *services.BetService) *BetHandler {
	return &BetHandler{service: service}
}

func (h *BetHandler) CreateBet(c *gin.Context) {
	// Получаем userID из контекста (установлен middleware AuthRequired)
	userIDFloat, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID := int(userIDFloat.(uint))

	// Парсим тело запроса
	type BetRequest struct {
		EventID int     `json:"event_id"`
		Amount  float32 `json:"amount"`
		Outcome string  `json:"outcome"`
	}
	var req BetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	log.Printf("Создание ставки: user_id=%d, event_id=%d, amount=%.2f, outcome=%s", userID, req.EventID, req.Amount, req.Outcome)

	newBet, err := h.service.Create(userID, req.EventID, req.Amount, req.Outcome)
	if err != nil {
		if err == gorm.ErrInvalidData {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
			return
		}
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

func (h *BetHandler) DeleteBet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bet ID"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bet not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bet deleted successfully"})
}
