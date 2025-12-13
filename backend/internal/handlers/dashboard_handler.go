package handlers

import (
	"medical-records-app/internal/services"
	"medical-records-app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	recordService    *services.RecordService
	medicationService *services.MedicationService
	reminderService   *services.ReminderService
}

func NewDashboardHandler(recordService *services.RecordService, medicationService *services.MedicationService, reminderService *services.ReminderService) *DashboardHandler {
	return &DashboardHandler{
		recordService:     recordService,
		medicationService: medicationService,
		reminderService:   reminderService,
	}
}

// GetDashboard returns dashboard summary
// @Summary Get dashboard
// @Description Get dashboard summary with active prescriptions, upcoming appointments, recent lab reports, and reminders
// @Tags dashboard
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /dashboard [get]
func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	// Get active prescriptions
	prescriptions, _, _ := h.recordService.GetPrescriptions(userID, 5, 0)
	
	// Get upcoming appointments
	appointments, _, _ := h.recordService.GetAppointments(userID, 5, 0, true)
	
	// Get recent lab reports
	labReports, _, _ := h.recordService.GetLabReports(userID, 5, 0)
	
	// Get active medications
	medications, _ := h.medicationService.GetMedications(userID, true)
	
	// Get upcoming reminders
	reminders, _ := h.reminderService.GetUpcomingReminders(userID, 30)

	c.JSON(http.StatusOK, gin.H{
		"prescriptions": prescriptions,
		"appointments":  appointments,
		"lab_reports":   labReports,
		"medications":   medications,
		"reminders":     reminders,
	})
}

