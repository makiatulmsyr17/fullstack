package main

import (
	"log"
	"makiatul/database"
	"makiatul/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	database.ConnectDB()
	if err := godotenv.Load(); err != nil {
		log.Fatal("gagal memuat file .env")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
	}))

	port := os.Getenv("PORT")

	routes.SetupRoute(app)

	app.Listen(":" + port)
}
