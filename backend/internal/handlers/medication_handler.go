package handlers

import (
	"medical-records-app/internal/database"
	"medical-records-app/internal/services"
	"medical-records-app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MedicationHandler struct {
	medicationService *services.MedicationService
}

func NewMedicationHandler(medicationService *services.MedicationService) *MedicationHandler {
	return &MedicationHandler{medicationService: medicationService}
}

// CreateMedication creates a new medication
// @Summary Create medication
// @Description Add a regular medication with pharmacy information
// @Tags medications
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param medication body database.Medication true "Medication details"
// @Success 201 {object} database.Medication
// @Router /medications [post]
func (h *MedicationHandler) CreateMedication(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	var medication database.Medication
	if err := c.ShouldBindJSON(&medication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.medicationService.CreateMedication(userID, &medication); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, medication)
}

// GetMedications retrieves all medications
// @Summary Get medications
// @Description Get all medications for the authenticated user
// @Tags medications
// @Security BearerAuth
// @Produce json
// @Param active query bool false "Active only" default(false)
// @Success 200 {object} map[string]interface{}
// @Router /medications [get]
func (h *MedicationHandler) GetMedications(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}
	activeOnly := c.Query("active") == "true"

	medications, err := h.medicationService.GetMedications(userID, activeOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medications})
}

// GetMedicationsNeedingRefill retrieves medications that need refill
// @Summary Get medications needing refill
// @Description Get medications that need to be refilled soon
// @Tags medications
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /medications/refill-needed [get]
func (h *MedicationHandler) GetMedicationsNeedingRefill(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	medications, err := h.medicationService.GetMedicationsNeedingRefill(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medications})
}

