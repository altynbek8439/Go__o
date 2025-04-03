package routes

import (
	"betting-site/internal/delivery"
	"betting-site/internal/repository"
	"betting-site/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Репозитории
	betRepo := repository.NewBetRepository(db)
	eventRepo := repository.NewEventRepository(db)

	// Сервисы
	betService := services.NewBetService(betRepo)
	eventService := services.NewEventService(eventRepo)

	// Хендлеры
	betHandler := delivery.NewBetHandler(betService)
	eventHandler := delivery.NewEventHandler(eventService)

	// API роуты
	api := r.Group("/api/v1")
	{
		// События
		api.GET("/events", eventHandler.GetEvents)
		api.POST("/events", eventHandler.CreateEvent)

		// Ставки
		api.POST("/bets", betHandler.CreateBet)
		api.GET("/bets/user/:user_id", betHandler.GetBetsByUser)
	}
}
