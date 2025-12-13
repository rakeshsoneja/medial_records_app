package handlers

import (
	"medical-records-app/internal/database"
	"medical-records-app/internal/services"
	"medical-records-app/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReminderHandler struct {
	reminderService *services.ReminderService
}

func NewReminderHandler(reminderService *services.ReminderService) *ReminderHandler {
	return &ReminderHandler{reminderService: reminderService}
}

// CreateReminder creates a new reminder
// @Summary Create reminder
// @Description Add a health check-up reminder
// @Tags reminders
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param reminder body database.Reminder true "Reminder details"
// @Success 201 {object} database.Reminder
// @Router /reminders [post]
func (h *ReminderHandler) CreateReminder(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	var reminder database.Reminder
	if err := c.ShouldBindJSON(&reminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.reminderService.CreateReminder(userID, &reminder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reminder)
}

// GetReminders retrieves all reminders
// @Summary Get reminders
// @Description Get all reminders for the authenticated user
// @Tags reminders
// @Security BearerAuth
// @Produce json
// @Param upcoming query bool false "Upcoming only" default(false)
// @Success 200 {object} map[string]interface{}
// @Router /reminders [get]
func (h *ReminderHandler) GetReminders(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}
	upcomingOnly := c.Query("upcoming") == "true"

	reminders, err := h.reminderService.GetReminders(userID, upcomingOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reminders})
}

// GetUpcomingReminders retrieves upcoming reminders
// @Summary Get upcoming reminders
// @Description Get reminders within the specified days ahead
// @Tags reminders
// @Security BearerAuth
// @Produce json
// @Param days query int false "Days ahead" default(30)
// @Success 200 {object} map[string]interface{}
// @Router /reminders/upcoming [get]
func (h *ReminderHandler) GetUpcomingReminders(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}
	days, _ := strconv.Atoi(c.DefaultQuery("days", "30"))

	reminders, err := h.reminderService.GetUpcomingReminders(userID, days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reminders})
}

