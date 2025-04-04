package main

import (
	"betting-site/internal/models"
	"betting-site/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=9801042 dbname=betting_site port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.AutoMigrate(&models.Event{}, &models.Bet{}, &models.User{})
	if err != nil {
		log.Fatal("Error on migrating to the DB", err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		log.Printf("Входящий запрос: %s %s", c.Request.Method, c.Request.URL.Path)
		log.Printf("Заголовки: %v", c.Request.Header)
		c.Next()
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(r, db)

	r.Run(":8080")
}
