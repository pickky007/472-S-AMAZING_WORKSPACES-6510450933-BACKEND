package controllers

import (
	"onez19/models"
	"onez19/services"

	"github.com/gofiber/fiber/v2"
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

func CreateActivity(ctx *fiber.Ctx) error {
	workspaceID := ctx.Params("workspaceId")
	sectionID, err := ctx.ParamsInt("sectionId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid section ID"})
	}

	var activityInput struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		StartDate   string `json:"start_date"` // use time.Time for actual date handling
		EndDate     string `json:"end_date"`   // use time.Time for actual date handling
	}

	if err := ctx.BodyParser(&activityInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	activity := models.Activity{
		Name:        activityInput.Name,
		Description: activityInput.Description,
		StartDate:   activityInput.StartDate,
		EndDate:     activityInput.EndDate,
		SectionID:   sectionID,
		WorkspaceID: workspaceID,
	}

	if err := services.CreateActivity(activity); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(activity)
}

func MoveActivity(ctx *fiber.Ctx) error {

	var moveInput struct {
		NewSectionID int `json:"new_section_id"`
		ActivityID   int `json:"activity_id"`
	}

	if err := ctx.BodyParser(&moveInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if moveInput.ActivityID == 0 || moveInput.NewSectionID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid activityId or newSectionId"})
	}

	if err := services.MoveActivity(moveInput.ActivityID, moveInput.NewSectionID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Activity moved successfully"})
}

func EditActivity(ctx *fiber.Ctx) error {
	workspaceID := ctx.Params("workspaceId")
	sectionID, err := ctx.ParamsInt("sectionId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid sectionId"})
	}

	activityID, err := ctx.ParamsInt("activityId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid activityId"})
	}

	if sectionID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid sectionId"})
	}

	if activityID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid activityId"})
	}

	var activityInput struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		StartDate   string `json:"start_date"` // use time.Time for actual date handling
		EndDate     string `json:"end_date"`   // use time.Time for actual date handling
	}

	if err := ctx.BodyParser(&activityInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	activity := models.Activity{
		ID:          activityID,
		Name:        activityInput.Name,
		Description: activityInput.Description,
		StartDate:   activityInput.StartDate,
		EndDate:     activityInput.EndDate,
		SectionID:   sectionID,
		WorkspaceID: workspaceID,
	}

	if err := services.EditActivity(activity); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(activity)
}
