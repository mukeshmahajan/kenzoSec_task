package tests

import (
	"kenzoSec_task/routes"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllTasks(t *testing.T) {
	router := gin.Default()
	routes.PublicRoutes(router)

	token := os.Getenv("AUTH_TOKEN") // Replace with your actual test token

	// Prepare the request with the Bearer token in the Authorization header
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks?page=1&limit=5", nil)
	req.Header.Add("Authorization", "Bearer "+token) // Add Bearer token here

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
