package main

import (
	"database/sql"
	"log"

	"go-backend-task/internal/db"
	"go-backend-task/internal/handler"
	"go-backend-task/internal/routes"
	"go-backend-task/internal/service"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	app := fiber.New()

	// PostgreSQL connection
	dbConn, err := sql.Open("pgx", "postgres://postgres:9514@localhost:5432/backenddb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(dbConn)
	userService := service.NewUserService(queries)
	userHandler := handler.NewUserHandler(userService)

	// Setup routes
	routes.SetupRoutes(app, userHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
