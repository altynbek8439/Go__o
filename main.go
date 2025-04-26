package main

import (
	"betting-site/internal/db"
	"betting-site/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db.InitDB()

	// Применение миграций
	m, err := migrate.New(
		"file://C:/Users/temir/betting-site/internal/db/migrations",
		"postgres://postgres:9801042@localhost:5432/betting_site1?sslmode=disable&search_path=public",
	)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("Applying migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("Migrations applied successfully")

	// Настройка Gin
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		log.Printf("Входящий запрос: %s %s", c.Request.Method, c.Request.URL.Path)
		log.Printf("Заголовки: %v", c.Request.Header)
		c.Next()
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"}, // Добавляем Authorization
	}))

	routes.SetupRoutes(r, db.DB)

	r.Run(":8080")
}
