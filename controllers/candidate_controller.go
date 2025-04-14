package controllers

import (
	"net/http"

	"github.com/DAT-CANDIDATE/db"
	"github.com/DAT-CANDIDATE/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

// CreateCandidate handles the creation of a candidate
func CreateCandidate(c *gin.Context) {
	var candidate models.Candidate

	if err := c.ShouldBindJSON(&candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validate.Struct(candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if Unique_Id is already taken by another record
	count, err := db.CandidateCollection.CountDocuments(c, bson.M{"unique_id": candidate.Unique_Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Unique_Id already exists"})
		return
	}

	candidate.ID = primitive.NewObjectID()
	_, err = db.CandidateCollection.InsertOne(c, candidate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create candidate"})
		return
	}

	c.JSON(http.StatusCreated, candidate)
}

// GetCandidate retrieves a candidate by unique id
func GetCandidateByUniqueId(c *gin.Context) {
	uniqueID := c.Param("unique_id")

	var candidate models.Candidate

	err := db.CandidateCollection.FindOne(c, bson.M{"unique_id": uniqueID}).Decode(&candidate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		return
	}

	c.JSON(http.StatusOK, candidate)
}

// GetCandidates retrieves all candidates
func GetCandidates(c *gin.Context) {

	var candidates []models.Candidate

	cursor, err := db.CandidateCollection.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var candidate models.Candidate
		if err := cursor.Decode(&candidate); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding candidate"})
			return
		}
		candidates = append(candidates, candidate)
	}

	c.JSON(http.StatusOK, candidates)
}

// UpdateCandidate by unique id updates a candidate's details
func UpdateCandidate(c *gin.Context) {
	uniqueID := c.Param("unique_id")

	var candidate models.Candidate
	if err := c.ShouldBindJSON(&candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validate.Struct(candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if Unique_Id is already taken by another record
	count, err := db.CandidateCollection.CountDocuments(c, bson.M{"unique_id": uniqueID, "_id": bson.M{"id": candidate.ID}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Unique_Id already exists"})
		return
	}

	// Update filter based on unique_id
	filter := bson.M{"unique_id": uniqueID}

	update := bson.M{"$set": candidate}
	_, err = db.CandidateCollection.UpdateOne(c, filter, update)
	//_, err = db.CandidateCollection.UpdateByID(ctx, uniqueId, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update candidate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Candidate updated successfully"})
}
