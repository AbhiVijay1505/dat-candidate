package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DAT-CANDIDATE/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateCandidate(t *testing.T) {
	// Setup the Gin router
	router := gin.Default()
	router.POST("/candidates", controllers.CreateCandidate)

	// Mock request body
	requestBody := []byte(`{
		"unique_id": "abcdef123456",
		"name": "Abhi Vijay",
		"address": "123 Abc Street",
		"contact_no": "+1234567890",
		"email": "abhi@gmail.com"
	}`)

	// Create a new HTTP request
	req, _ := http.NewRequest("POST", "/candidates", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(recorder, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, recorder.Code, "Response code should be 201 Created")
	assert.Contains(t, recorder.Body.String(), `"unique_id":"abcdef123456"`, "Response should contain the unique_id")
}
