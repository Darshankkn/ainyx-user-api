package main

import (
	"log"

	"ainyx-user-api/config"
	"ainyx-user-api/internal/handler"
	"ainyx-user-api/internal/repository"
	"ainyx-user-api/internal/routes"
	"ainyx-user-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"ainyx-user-api/internal/logger"
)

func main() {

	conn, err := config.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(nil)

	userRepo := repository.NewUserRepository(conn)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)
	logger.InitLogger()
    defer logger.Log.Sync()

	app := fiber.New()

	routes.SetupRoutes(app, userHandler)

	log.Fatal(app.Listen(":3000"))
}