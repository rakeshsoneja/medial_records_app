package handlers

import (
	"medical-records-app/internal/database"
	"medical-records-app/internal/services"
	"medical-records-app/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RecordHandler struct {
	recordService *services.RecordService
}

func NewRecordHandler(recordService *services.RecordService) *RecordHandler {
	return &RecordHandler{recordService: recordService}
}

// CreatePrescription creates a new prescription
// @Summary Create prescription
// @Description Add a new prescription record
// @Tags prescriptions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param prescription body database.Prescription true "Prescription details"
// @Success 201 {object} database.Prescription
// @Failure 400 {object} map[string]string
// @Router /prescriptions [post]
func (h *RecordHandler) CreatePrescription(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	var prescription database.Prescription
	if err := c.ShouldBindJSON(&prescription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.recordService.CreatePrescription(userID, &prescription); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, prescription)
}

// GetPrescriptions retrieves all prescriptions for the user
// @Summary Get prescriptions
// @Description Get all prescriptions for the authenticated user
// @Tags prescriptions
// @Security BearerAuth
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} map[string]interface{}
// @Router /prescriptions [get]
func (h *RecordHandler) GetPrescriptions(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	prescriptions, total, err := h.recordService.GetPrescriptions(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  prescriptions,
		"total": total,
		"limit": limit,
		"offset": offset,
	})
}

// GetPrescription retrieves a single prescription
// @Summary Get prescription
// @Description Get a specific prescription by ID
// @Tags prescriptions
// @Security BearerAuth
// @Produce json
// @Param id path string true "Prescription ID"
// @Success 200 {object} database.Prescription
// @Failure 404 {object} map[string]string
// @Router /prescriptions/{id} [get]
func (h *RecordHandler) GetPrescription(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}
	prescriptionID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prescription ID"})
		return
	}

	prescription, err := h.recordService.GetPrescriptionByID(userID, prescriptionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Prescription not found"})
		return
	}

	c.JSON(http.StatusOK, prescription)
}

// UpdatePrescription updates a prescription
// @Summary Update prescription
// @Description Update an existing prescription
// @Tags prescriptions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Prescription ID"
// @Param prescription body map[string]interface{} true "Update fields"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /prescriptions/{id} [put]
func (h *RecordHandler) UpdatePrescription(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}
	prescriptionID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prescription ID"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.recordService.UpdatePrescription(userID, prescriptionID, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Prescription updated successfully"})
}

// DeletePrescription deletes a prescription
// @Summary Delete prescription
// @Description Delete a prescription record
// @Tags prescriptions
// @Security BearerAuth
// @Produce json
// @Param id path string true "Prescription ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /prescriptions/{id} [delete]
func (h *RecordHandler) DeletePrescription(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}
	prescriptionID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prescription ID"})
		return
	}

	if err := h.recordService.DeletePrescription(userID, prescriptionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Prescription deleted successfully"})
}

// CreateAppointment creates a new appointment
// @Summary Create appointment
// @Description Add a new appointment
// @Tags appointments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param appointment body database.Appointment true "Appointment details"
// @Success 201 {object} database.Appointment
// @Failure 400 {object} map[string]string
// @Router /appointments [post]
func (h *RecordHandler) CreateAppointment(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	var appointment database.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.recordService.CreateAppointment(userID, &appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

// GetAppointments retrieves all appointments
// @Summary Get appointments
// @Description Get all appointments for the authenticated user
// @Tags appointments
// @Security BearerAuth
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Param upcoming query bool false "Upcoming only" default(false)
// @Success 200 {object} map[string]interface{}
// @Router /appointments [get]
func (h *RecordHandler) GetAppointments(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	upcomingOnly := c.Query("upcoming") == "true"

	appointments, total, err := h.recordService.GetAppointments(userID, limit, offset, upcomingOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  appointments,
		"total": total,
		"limit": limit,
		"offset": offset,
	})
}

// CreateLabReport creates a new lab report
// @Summary Create lab report
// @Description Add a new lab report
// @Tags lab-reports
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param labReport body database.LabReport true "Lab report details"
// @Success 201 {object} database.LabReport
// @Failure 400 {object} map[string]string
// @Router /lab-reports [post]
func (h *RecordHandler) CreateLabReport(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	var labReport database.LabReport
	if err := c.ShouldBindJSON(&labReport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.recordService.CreateLabReport(userID, &labReport); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, labReport)
}

// GetLabReports retrieves all lab reports
// @Summary Get lab reports
// @Description Get all lab reports for the authenticated user
// @Tags lab-reports
// @Security BearerAuth
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} map[string]interface{}
// @Router /lab-reports [get]
func (h *RecordHandler) GetLabReports(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	labReports, total, err := h.recordService.GetLabReports(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  labReports,
		"total": total,
		"limit": limit,
		"offset": offset,
	})
}

// CreateHealthInsurance creates health insurance record
// @Summary Create health insurance
// @Description Add health insurance information
// @Tags insurance
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param insurance body database.HealthInsurance true "Insurance details"
// @Success 201 {object} database.HealthInsurance
// @Router /insurance [post]
func (h *RecordHandler) CreateHealthInsurance(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	var insurance database.HealthInsurance
	if err := c.ShouldBindJSON(&insurance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.recordService.CreateHealthInsurance(userID, &insurance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, insurance)
}

// GetHealthInsurances retrieves all health insurance records
// @Summary Get health insurances
// @Description Get all health insurance records
// @Tags insurance
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /insurance [get]
func (h *RecordHandler) GetHealthInsurances(c *gin.Context) {
	userID, ok := utils.MustGetUserID(c)
	if !ok {
		return
	}

	insurances, err := h.recordService.GetHealthInsurances(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": insurances})
}

