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

	//Posts
	app.Post("/api/posts/create", controllers.CreatePost)
	app.Post("/api/posts/delete", controllers.DeletePost)
	app.Get("/api/posts/getall", controllers.GetPosts)
	app.Get("/api/posts/get", controllers.GetPost)
	app.Post("/api/posts/update", controllers.UpdatePost)
}
