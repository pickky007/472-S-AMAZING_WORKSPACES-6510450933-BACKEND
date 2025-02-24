package controllers
import (
    "time"
	"onez19/models"
	"onez19/services"
	"github.com/gofiber/fiber/v2"
)

func CreateMessage(ctx *fiber.Ctx) error {
	var messageInput struct {
		Message string `json:"message"`
		Username string `json:"username"`
		WorkspaceID string `json:"workspace_id"`
	}

	if err := ctx.BodyParser(&messageInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	message := models.Message{
		Message: messageInput.Message,
		Date: time.Now().Format(time.RFC3339),
		WorkspaceID: messageInput.WorkspaceID,
		Username: messageInput.Username,
	
	}

	if err := services.CreateMessage(message); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(message)
}

func GetAllMessagesByWorkspaceID(c *fiber.Ctx) error {
	workspaceID := c.Params("workspaceId")
	
	messages, err := services.GetAllMessagesByWorkspaceID(workspaceID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch message"})
	}

	return c.JSON(messages)
}
