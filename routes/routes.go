package routes

import (
	"approvez-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)

	//Auth
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)

	//Campaign
	app.Post("/api/campaign/create", controllers.CreateCampaign)
	app.Get("/api/campaign/get", controllers.GetCampaigns)
}
