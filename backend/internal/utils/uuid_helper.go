package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetUserIDFromContext extracts and parses user_id from gin context
func GetUserIDFromContext(c *gin.Context) (uuid.UUID, error) {
	userIDStr, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, gorm.ErrRecordNotFound
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

// MustGetUserID extracts user_id from context or returns error response
func MustGetUserID(c *gin.Context) (uuid.UUID, bool) {
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return uuid.Nil, false
	}
	return userID, true
}

