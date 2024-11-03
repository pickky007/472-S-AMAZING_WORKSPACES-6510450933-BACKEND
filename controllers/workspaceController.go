package controllers

import (
	"github.com/gofiber/fiber/v2"
	"onez19/services"
)

func GetAllWorkspacesByUsername(c *fiber.Ctx) error {
	username := c.Params("username") // รับ username จากพารามิเตอร์ใน URL

	workspaces, err := services.GetWorkspacesByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch workspaces"})
	}

	return c.JSON(workspaces) // ส่งข้อมูล workspaces กลับไปยัง client
}
