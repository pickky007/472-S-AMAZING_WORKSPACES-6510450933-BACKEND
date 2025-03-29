package controllers

import (
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"onez19/models"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func CreateMessageMockUp(serviceFunc func(models.Message) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var message models.Message
		if err := c.BodyParser(&message); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		if err := serviceFunc(message); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create message"})
		}

		return c.JSON(fiber.Map{
			"message":      message.Message,
			"username":     message.Username,
			"workspace_id": message.WorkspaceID,
		})
	}
}

func TestCreateMessage(t *testing.T) {
	tests := []struct {
		name       string
		inputBody  string
		mockFunc   func(models.Message) error
		wantStatus int
		wantBody   string
	}{
		{
			name:      "Valid message",
			inputBody: `{"message":"Hello","username":"user1","workspace_id":"123"}`,
			mockFunc: func(message models.Message) error {
				return nil
			},
			wantStatus: fiber.StatusOK,
			wantBody:   `{"message":"Hello","username":"user1","workspace_id":"123"}`,
		},
		{
			name:      "Service error",
			inputBody: `{"message":"Hello","username":"user1","workspace_id":"123"}`,
			mockFunc: func(message models.Message) error {
				return errors.New("Service error")
			},
			wantStatus: fiber.StatusInternalServerError,
			wantBody:   `{"error":"Failed to create message"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Post("/message", CreateMessageMockUp(tt.mockFunc))

			req := httptest.NewRequest("POST", "/message", strings.NewReader(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("Status = %v, want %v", resp.StatusCode, tt.wantStatus)
			}

			body, _ := ioutil.ReadAll(resp.Body)
			if strings.TrimSpace(string(body)) != tt.wantBody {
				t.Errorf("Body = %v, want %v", string(body), tt.wantBody)
			}
		})
	}
}

// Mock DeleteMessage function
func DeleteMessageMockUp(serviceFunc func(messageID string) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var deleteInput struct {
			ID string `json:"id"`
		}

		if err := c.BodyParser(&deleteInput); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		err := serviceFunc(deleteInput.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete message"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Message deleted successfully"})
	}
}

func TestDeleteMessage(t *testing.T) {
	tests := []struct {
		name       string
		inputBody  string
		mockFunc   func(messageID string) error
		wantStatus int
		wantBody   string
	}{
		{
			name:      "Valid Delete Request",
			inputBody: `{"id":"12345"}`,
			mockFunc: func(messageID string) error {
				return nil
			},
			wantStatus: fiber.StatusOK,
			wantBody:   `{"message":"Message deleted successfully"}`,
		},
		{
			name:      "Invalid JSON Request",
			inputBody: `{"invalid":}`,
			mockFunc: func(messageID string) error {
				return nil
			},
			wantStatus: fiber.StatusBadRequest,
		},
		{
			name:      "Service Error",
			inputBody: `{"id":"12345"}`,
			mockFunc: func(messageID string) error {
				return errors.New("database error")
			},
			wantStatus: fiber.StatusInternalServerError,
			wantBody:   `{"error":"Failed to delete message"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Delete("/message", DeleteMessageMockUp(tt.mockFunc))

			req := httptest.NewRequest("DELETE", "/message", strings.NewReader(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, tt.wantStatus, resp.StatusCode)

			body, _ := ioutil.ReadAll(resp.Body)
			if tt.wantBody != "" {
				assert.JSONEq(t, tt.wantBody, string(body))
			}
		})
	}
}

// Mock GetAllMessagesByWorkspaceID function
func GetAllMessagesByWorkspaceIDMockUp(serviceFunc func(workspaceID string) ([]models.Message, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		workspaceID := c.Params("workspaceId")

		messages, err := serviceFunc(workspaceID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch message"})
		}

		return c.JSON(messages)
	}
}

func TestGetAllMessagesByWorkspaceID(t *testing.T) {
	tests := []struct {
		name        string
		workspaceID string
		mockFunc    func(workspaceID string) ([]models.Message, error)
		wantStatus  int
		wantBody    string
	}{
		{
			name:        "Valid Workspace ID",
			workspaceID: "workspace123",
			mockFunc: func(workspaceID string) ([]models.Message, error) {
				return []models.Message{
					{ID: 1, Message: "Hello", WorkspaceID: "workspace123", Username: "user1"},
				}, nil
			},
			wantStatus: fiber.StatusOK,
			wantBody:   `[{"date":"0001-01-01T00:00:00Z", "id":1, "message":"Hello", "username":"user1", "workspace_id":"workspace123"}]`,
		},
		{
			name:        "Service Error",
			workspaceID: "workspace123",
			mockFunc: func(workspaceID string) ([]models.Message, error) {
				return nil, errors.New("database error")
			},
			wantStatus: fiber.StatusInternalServerError,
			wantBody:   `{"error":"Failed to fetch message"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/messages/:workspaceId", GetAllMessagesByWorkspaceIDMockUp(tt.mockFunc))

			req := httptest.NewRequest("GET", "/messages/"+tt.workspaceID, nil)
			resp, _ := app.Test(req)

			assert.Equal(t, tt.wantStatus, resp.StatusCode)

			body, _ := ioutil.ReadAll(resp.Body)
			if tt.wantBody != "" {
				assert.JSONEq(t, tt.wantBody, string(body))
			}
		})
	}
}
