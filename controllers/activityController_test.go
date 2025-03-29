package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"onez19/models"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func GetActivitiesByWorkspaceMockUp(serviceFunc func(workspaceID string) ([]models.Activity, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		workspaceID := c.Params("workspaceId")

		// Debugging
		fmt.Printf("Mock function called with workspaceID: %s\n", workspaceID)

		activities, err := serviceFunc(workspaceID)
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch activities"})
		}

		return c.JSON(activities)
	}
}

func TestGetActivitiesByWorkspace(t *testing.T) {

	tests := []struct {
		name        string
		workspaceID string
		mockFunc    func(workspaceID string) ([]models.Activity, error)
		wantStatus  int
		wantBody    string
	}{
		{
			name:        "Valid workspaceID",
			workspaceID: "valid-id",
			mockFunc: func(workspaceID string) ([]models.Activity, error) {
				return []models.Activity{
					{ID: 1, Name: "Activity1", Description: "Description1"},
				}, nil
			},
			wantStatus: fiber.StatusOK,
			wantBody:   `[{"id":1,"name":"Activity1","description":"Description1","start_date":"","end_date":"","section_id":0,"owner":"","workspace_id":""}]`,
		},
		{
			name:        "Invalid workspaceID",
			workspaceID: "invalid-id",
			mockFunc: func(workspaceID string) ([]models.Activity, error) {
				return nil, errors.New("Service error")
			},
			wantStatus: fiber.StatusInternalServerError,
			wantBody:   `{"error":"Failed to fetch activities"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/activities/:workspaceId", GetActivitiesByWorkspaceMockUp(tt.mockFunc))

			req := httptest.NewRequest("GET", "/activities/"+tt.workspaceID, nil)
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
