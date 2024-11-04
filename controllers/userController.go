// controllers/userController.go
package controllers

import (
	"onez19/models"
	"onez19/services"

	"github.com/gofiber/fiber/v2"
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
func Register(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if user.Username == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "all fields are required"})
	}

	if err := services.RegisterUser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}
