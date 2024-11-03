package routes

import (
	"github.com/gofiber/fiber/v2"
	"onez19/controllers"
)

func SetupRoutes(app *fiber.App) {
	// auth := app.Group("/auth")
	// auth.Post("/register", controllers.Register)
	// auth.Post("/login", controllers.Login)
	app.Get("/users", controllers.GetUsers)
	app.Get("/workspaces/:workspaceId/sections", controllers.GetAllSectionsByWorkspaceID)
	app.Get("/users/:username/workspaces", controllers.GetAllWorkspacesByUsername)
}
