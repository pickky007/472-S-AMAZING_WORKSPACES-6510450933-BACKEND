// controllers/userController.go
package controllers

import (
	"github.com/gofiber/fiber/v2"
	"onez19/services"
)

func GetUsers(c *fiber.Ctx) error {
	// ดึงข้อมูลผู้ใช้ทั้งหมดจาก userService
	users, err := services.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	// ส่งข้อมูลผู้ใช้กลับไปยัง client
	return c.JSON(users)
}

// ฟังก์ชัน Register และ Login ยังคงเหมือนเดิม
