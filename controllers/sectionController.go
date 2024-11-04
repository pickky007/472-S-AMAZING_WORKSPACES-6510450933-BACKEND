package controllers

import (
	"github.com/gofiber/fiber/v2"
	"onez19/services"
)

func GetAllSectionsByWorkspaceID(c *fiber.Ctx) error {
	workspaceID := c.Params("workspaceId") // รับ workspace ID จากพารามิเตอร์ใน URL เป็น string
	if len(workspaceID) != 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid workspace ID length"})
	}

	// เรียกใช้งานบริการเพื่อนำข้อมูล sections ตาม workspace ID
	sections, err := services.GetAllSectionsByWorkspaceID(workspaceID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch sections"})
	}

	return c.JSON(sections) // ส่งข้อมูล sections กลับไปยัง client
}
