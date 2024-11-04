package controllers

import (
	"onez19/models"
	"onez19/services"

	"github.com/gofiber/fiber/v2"
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

func CreateSection(ctx *fiber.Ctx) error {
	workspaceID := ctx.Params("workspaceId")

	var sectionInput struct {
		Name string `json:"name"`
	}

	if err := ctx.BodyParser(&sectionInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	section := models.Section{
		WorkspaceID: workspaceID,
		Name:        sectionInput.Name,
	}

	if err := services.CreateSection(section); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(section)
}

func EditSectionName(ctx *fiber.Ctx) error {
	sectionID, err := ctx.ParamsInt("sectionId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid sectionId"})
	}

	if sectionID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid sectionId"})
	}

	var sectionInput struct {
		NewName string `json:"new_name"`
	}

	if err := ctx.BodyParser(&sectionInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := services.EditSectionName(sectionID, sectionInput.NewName); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Section name updated successfully"})
}
