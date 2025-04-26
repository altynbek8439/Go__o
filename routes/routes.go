package routes

import (
	"betting-site/internal/auth"
	"betting-site/internal/delivery"
	"betting-site/internal/middleware"
	"betting-site/internal/repository"
	"betting-site/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	betRepo := repository.NewBetRepository(db)
	eventRepo := repository.NewEventRepository(db)

	betService := services.NewBetService(betRepo, db) // Передаём db
	eventService := services.NewEventService(eventRepo)

	betHandler := delivery.NewBetHandler(betService)
	eventHandler := delivery.NewEventHandler(eventService)

	// Роуты для аутентификации
	authRoutes := r.Group("/api/v1/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}

	// Открытые роуты
	api := r.Group("/api/v1")
	{
		api.GET("/events", eventHandler.GetEvents)
	}

	// Защищенные роуты
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/me", auth.Me)
		protected.POST("/events", eventHandler.CreateEvent)
		protected.PUT("/events/:id", eventHandler.UpdateEvent)
		protected.DELETE("/events/:id", eventHandler.DeleteEvent)

		protected.POST("/bets", betHandler.CreateBet)
		protected.GET("/bets/user/:user_id", betHandler.GetBetsByUser)
		protected.DELETE("/bets/:id", betHandler.DeleteBet)
	}
}
