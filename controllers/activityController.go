package controllers

import (
	"github.com/gofiber/fiber/v2"
	"onez19/services"
)

func GetActivitiesBySectionAndWorkspace(c *fiber.Ctx) error {
	sectionID, err := c.ParamsInt("sectionId") // รับ section ID จากพารามิเตอร์ใน URL
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid section ID"})
	}

	workspaceID := c.Params("workspaceId") // รับ workspace ID จากพารามิเตอร์ใน URL (เป็น string)

	activities, err := services.GetActivitiesBySectionAndWorkspace(sectionID, workspaceID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch activities"})
	}

	return c.JSON(activities) // ส่งข้อมูล activities กลับไปยัง client
}
