package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateCandidate(t *testing.T) {
	// Setup the Gin router
	router := gin.New()
	mockCreateCandidate := func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"unique_id": "abcdef123456",
		})
	}
	router.POST("/candidates", mockCreateCandidate)

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
	assert.Equal(t, http.StatusCreated, recorder.Code, "Expected HTTP Response code should be 200 Created")
	assert.Contains(t, recorder.Body.String(), `"unique_id":"abcdef123456"`, "Response body should contain the unique_id 'abcdef123456'")
}

func TestGetCandidateByUniqueId(t *testing.T) {
	// Setup the Gin router
	router := gin.New()
	mockGetCandidateByUniqueId := func(c *gin.Context) {
		uniqueId := c.Param("unique_id")
		if uniqueId == "abcdef123456" {
			c.JSON(http.StatusOK, gin.H{
				"unique_id":  "abcdef123456",
				"name":       "Abhi Vijay",
				"address":    "123 Abc Street",
				"contact_no": "+1234567890",
				"email":      "abhi@gmail.com",
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		}
	}
	router.GET("/candidates/:unique_id", mockGetCandidateByUniqueId)

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/candidates/abcdef123456", nil)

	// Create a ResponseRecorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(recorder, req)

	// Check the response
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected HTTP Response code should be 200 OK")
	assert.Contains(t, recorder.Body.String(), `"unique_id":"abcdef123456"`, "Response body should contain the unique_id 'abcdef123456'")
	assert.Contains(t, recorder.Body.String(), `"name":"Abhi Vijay"`, "Response body should contain the name 'Abhi Vijay'")
}

func TestGetCandidateByUniqueIdNotFound(t *testing.T) {
	// Setup the Gin router
	router := gin.New()
	mockGetCandidateByUniqueId := func(c *gin.Context) {
		uniqueId := c.Param("unique_id")
		if uniqueId == "abcdef123456" {
			c.JSON(http.StatusOK, gin.H{
				"unique_id":  "abcdef123456",
				"name":       "Abhi Vijay",
				"address":    "123 Abc Street",
				"contact_no": "+1234567890",
				"email":      "abhi@gmail.com",
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		}
	}
	router.GET("/candidates/:unique_id", mockGetCandidateByUniqueId)

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/candidates/wrongId12345", nil)

	// Create a ResponseRecorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(recorder, req)

	// Check the response
	assert.Equal(t, http.StatusNotFound, recorder.Code, "Expected HTTP Response code should be 404 Not Found")
	assert.Contains(t, recorder.Body.String(), `"error":"Candidate not found"`, "Response body should contain the error message 'Candidate not found'")
}

func TestUpdateCandidate(t *testing.T) {

	// Setup the Gin router
	router := gin.New()
	mockUpdateCandidate := func(c *gin.Context) {
		uniqueId := c.Param("unique_id")

		if uniqueId == "abcdef123456" {
			c.JSON(http.StatusOK, gin.H{
				"unique_id": uniqueId,
				"name":      "Updated Abhi Vijay",
				"email":     "Updatedabhi@gmail.com",
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		}
	}
	router.PUT("/candidates/:unique_id", mockUpdateCandidate)
	// Create a new HTTP request
	req, _ := http.NewRequest("PUT", "/candidates/abcdef123456", nil)
	req.Header.Set("Content-Type", "application/json")
	// Create a ResponseRecorder to record the response
	recorder := httptest.NewRecorder()
	// Serve the request
	router.ServeHTTP(recorder, req)
	// Check the response
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected HTTP Response code should be 200 OK")
	assert.Contains(t, recorder.Body.String(), `"unique_id":"abcdef123456"`, "Response body should contain the unique_id 'abcdef123456'")
	assert.Contains(t, recorder.Body.String(), `"name":"Updated Abhi Vijay"`, "Response body should contain the updated name 'Updated Abhi Vijay'")
	assert.Contains(t, recorder.Body.String(), `"email":"Updatedabhi@gmail.com"`, "Response body should contain the updated email 'Updatedabhi@gmail.com'")

}
func TestUpdateCandidateNotFound(t *testing.T) {
	router := gin.New()
	mockUpdateCandidate := func(c *gin.Context) {
		uniqueId := c.Param("unique_id")
		if uniqueId == "abcdef123456" {
			c.JSON(http.StatusOK, gin.H{
				"unique_id": uniqueId,
				"name":      "Updated Abhi Vijay",
				"email":     "updateabhi@gmail.com",
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		}
	}
	router.PUT("/candidates/:unique_id", mockUpdateCandidate)
	// Create a new HTTP request
	req, _ := http.NewRequest("PUT", "/candidates/wrongId12345", nil)
	req.Header.Set("Content-Type", "application/json")
	// Create a ResponseRecorder to record the response
	recorder := httptest.NewRecorder()
	// Serve the request
	router.ServeHTTP(recorder, req)
	// Check the response
	assert.Equal(t, http.StatusNotFound, recorder.Code, "Expected HTTP Response code should be 404 Not Found")
	assert.Contains(t, recorder.Body.String(), `"error":"Candidate not found"`, "Response body should contain the error message 'Candidate not found'")
}
func TestGetCandidates(t *testing.T) {
	router := gin.New()
	mockGetCandidates := func(c *gin.Context) {
		c.JSON(http.StatusOK, []gin.H{
			{
				"unique_id":  "abcdef123456",
				"name":       "Abhi Vijay",
				"address":    "123 Abc Street",
				"contact_no": "+1234567890",
				"email":      "abhivijay@gmail.com",
			},
		})
	}
	router.GET("/candidates", mockGetCandidates)
	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/candidates", nil)
	// Create a ResponseRecorder to record the response
	recorder := httptest.NewRecorder()
	// Serve the request
	router.ServeHTTP(recorder, req)
	// Check the response
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected HTTP Response code should be 200 OK")
	assert.Contains(t, recorder.Body.String(), `"unique_id":"abcdef123456"`, "Response body should contain the unique_id 'abcdef123456'")
	assert.Contains(t, recorder.Body.String(), `"name":"Abhi Vijay"`, "Response body should contain the name 'Abhi Vijay'")
	assert.Contains(t, recorder.Body.String(), `"address":"123 Abc Street"`, "Response body should contain the address '123 Abc Street'")
	assert.Contains(t, recorder.Body.String(), `"contact_no":"+1234567890"`, "Response body should contain the contact_no '+1234567890'")
}
