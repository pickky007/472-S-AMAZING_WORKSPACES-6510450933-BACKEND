package routes

import (
	"onez19/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// auth := app.Group("/auth")
	// auth.Post("/register", controllers.Register)
	// auth.Post("/login", controllers.Login)
	app.Get("/users", controllers.GetUsers)
	app.Get("/workspaces/:workspaceId/sections", controllers.GetAllSectionsByWorkspaceID)
	app.Get("/users/:username/workspaces", controllers.GetAllWorkspacesByUsername)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/users/:username/workspaces/create", controllers.CreateWorkspace)
}
