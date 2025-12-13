package handlers

import (
	"encoding/json"
	"medical-records-app/internal/services"
	"medical-records-app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SharingHandler struct {
	sharingService *services.SharingService
}

func NewSharingHandler(sharingService *services.SharingService) *SharingHandler {
	return &SharingHandler{sharingService: sharingService}
}

type CreateShareRequest struct {
	RecordType      string      `json:"record_type" binding:"required"`
	RecordIDs       []uuid.UUID `json:"record_ids" binding:"required"`
	ExpiresInHours  int         `json:"expires_in_hours" binding:"required"`
	MaxAccessCount  int         `json:"max_access_count"`
	AllowDownload   bool        `json:"allow_download"`
	RecipientEmail  string      `json:"recipient_email"`
	RecipientPhone  string      `json:"recipient_phone"`
	ShareMethod     string      `json:"share_method" binding:"required"` // email, sms, link
}

// CreateShareLink creates a shareable link
// @Summary Create share link
// @Description Create a time-limited shareable link for medical records
// @Tags sharing
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body CreateShareRequest true "Share link details"
// @Success 201 {object} database.SharedRecord
// @Failure 400 {object} map[string]string
// @Router /sharing/create [post]
func (h *SharingHandler) CreateShareLink(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	var req CreateShareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sharedRecord, err := h.sharingService.CreateShareLink(
		userID,
		req.RecordType,
		req.RecordIDs,
		req.ExpiresInHours,
		req.MaxAccessCount,
		req.AllowDownload,
		req.RecipientEmail,
		req.RecipientPhone,
		req.ShareMethod,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: Send email/SMS if shareMethod is email or sms

	c.JSON(http.StatusCreated, gin.H{
		"shared_record": sharedRecord,
		"share_url":     "/share/" + sharedRecord.ShareToken,
	})
}

// GetSharedRecord retrieves a shared record by token
// @Summary Get shared record
// @Description Access a shared medical record using the share token
// @Tags sharing
// @Produce json
// @Param token path string true "Share Token"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /share/{token} [get]
func (h *SharingHandler) GetSharedRecord(c *gin.Context) {
	token := c.Param("token")

	sharedRecord, err := h.sharingService.GetSharedRecordByToken(token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Share link not found or expired"})
		return
	}

	// Record access
	ipAddress := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")
	h.sharingService.RecordAccess(sharedRecord.ID, ipAddress, userAgent, "viewed")

	// Parse record IDs from JSON
	var recordIDs []uuid.UUID
	if err := json.Unmarshal([]byte(sharedRecord.RecordIDs), &recordIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse record IDs"})
		return
	}

	// Get the actual records
	records, err := h.sharingService.GetRecordsByIDs(sharedRecord.RecordType, recordIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shared_record": sharedRecord,
		"records":       records,
		"allow_download": sharedRecord.AllowDownload,
	})
}

// GetMySharedRecords gets all shared records created by the user
// @Summary Get my shared records
// @Description Get all share links created by the authenticated user
// @Tags sharing
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /sharing/my-shares [get]
func (h *SharingHandler) GetMySharedRecords(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	sharedRecords, err := h.sharingService.GetSharedRecordsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sharedRecords})
}

// RevokeShareLink revokes a share link
// @Summary Revoke share link
// @Description Revoke an active share link
// @Tags sharing
// @Security BearerAuth
// @Produce json
// @Param id path string true "Share Record ID"
// @Success 200 {object} map[string]string
// @Router /sharing/{id}/revoke [post]
func (h *SharingHandler) RevokeShareLink(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}
	shareID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid share ID"})
		return
	}

	if err := h.sharingService.RevokeShareLink(userID, shareID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Share link revoked successfully"})
}

