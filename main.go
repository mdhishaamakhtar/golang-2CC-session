package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mdhishaamakhtar/learnGo/api/endpoints"
	"github.com/mdhishaamakhtar/learnGo/pkg/database"
	"github.com/mdhishaamakhtar/learnGo/pkg/models"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error: .env not found. Relying on environment variables")
	}
	// Connecting to Database
	db, err := database.ConnectDB()
	if err != nil {
		log.Panicln(fmt.Errorf("fatal error in connecting to db: %w", err))
	}
	log.Println("Connected to Database")

	// Running Migrations
	err = db.AutoMigrate(&models.Posts{})
	if err != nil {
		log.Panicln(fmt.Errorf("error migrating models: %w", err))
	}
	log.Println("Database Migration Done")

	app := fiber.New()

	// Initializing any middlewares
	app.Use(cors.New())
	app.Use(logger.New())

	// Mounting routes
	endpoints.MountRoutes(app)

	// Running the server
	log.Fatal(app.Listen(":3000"))
}
