package main

import (
	"approvez-backend/database"
	"approvez-backend/routes"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	if err := database.MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	app.Listen(":8000")
}
