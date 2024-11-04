package controllers

import (
	"onez19/models"
	"onez19/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllWorkspacesByUsername(c *fiber.Ctx) error {
	username := c.Params("username") // รับ username จากพารามิเตอร์ใน URL

	workspaces, err := services.GetWorkspacesByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch workspaces"})
	}

	return c.JSON(workspaces) // ส่งข้อมูล workspaces กลับไปยัง client
}

func CreateWorkspace(ctx *fiber.Ctx) error {
	owner := ctx.Params("username")
	var workspaceInput struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := ctx.BodyParser(&workspaceInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	workspace := models.Workspace{
		Name:        workspaceInput.Name,
		Description: workspaceInput.Description,
		Owner:       owner,
	}

	if err := services.CreateWorkspace(workspace); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(workspace)
}
