package controllers

import (
	"onez19/models"
	"onez19/services"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateMessage(ctx *fiber.Ctx) error {
	var messageInput struct {
		Message     string `json:"message"`
		Username    string `json:"username"`
		WorkspaceID string `json:"workspace_id"`
	}

	if err := ctx.BodyParser(&messageInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	message := models.Message{
		Message:     messageInput.Message,
		Date:        time.Now().Format(time.RFC3339),
		WorkspaceID: messageInput.WorkspaceID,
		Username:    messageInput.Username,
	}

	if err := services.CreateMessage(message); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(message)
}

func DeleteMessage(ctx *fiber.Ctx) error {
	var deleteInput struct {
		ID string `json:"id"`
	}

	if err := ctx.BodyParser(&deleteInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := services.DeleteMessage(deleteInput.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete message"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Message deleted successfully"})
}

func GetAllMessagesByWorkspaceID(c *fiber.Ctx) error {
	workspaceID := c.Params("workspaceId")

	messages, err := services.GetAllMessagesByWorkspaceID(workspaceID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch message"})
	}

	return c.JSON(messages)
}

func SearchMessages(c *fiber.Ctx) error {
	workspaceID := c.Params("workspaceId")
	query := c.Query("query")
	regex := c.Query("regex")

	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Query parameter cannot be empty",
		})
	}

	var messages []models.Message
	var err error

	if regex == "true" {
		_, err = regexp.Compile(query)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid regular expression",
			})
		}
		messages, err = services.SearchMessagesByRegex(query, workspaceID)
	} else {
		messages, err = services.SearchMessagesByText(query, workspaceID)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(messages)
}
