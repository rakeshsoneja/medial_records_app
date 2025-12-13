package services

import (
	"medical-records-app/internal/database"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RecordService struct {
	db *gorm.DB
}

func NewRecordService(db *gorm.DB) *RecordService {
	return &RecordService{db: db}
}

// Prescription methods
func (s *RecordService) CreatePrescription(userID uuid.UUID, prescription *database.Prescription) error {
	prescription.UserID = userID
	prescription.ID = uuid.New()
	prescription.CreatedAt = time.Now()
	prescription.UpdatedAt = time.Now()
	return s.db.Create(prescription).Error
}

func (s *RecordService) GetPrescriptions(userID uuid.UUID, limit, offset int) ([]database.Prescription, int64, error) {
	var prescriptions []database.Prescription
	var total int64

	query := s.db.Where("user_id = ?", userID)
	s.db.Model(&database.Prescription{}).Where("user_id = ?", userID).Count(&total)

	if err := query.Order("prescription_date DESC").Limit(limit).Offset(offset).Find(&prescriptions).Error; err != nil {
		return nil, 0, err
	}

	return prescriptions, total, nil
}

func (s *RecordService) GetPrescriptionByID(userID, prescriptionID uuid.UUID) (*database.Prescription, error) {
	var prescription database.Prescription
	if err := s.db.Where("id = ? AND user_id = ?", prescriptionID, userID).First(&prescription).Error; err != nil {
		return nil, err
	}
	return &prescription, nil
}

func (s *RecordService) UpdatePrescription(userID, prescriptionID uuid.UUID, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return s.db.Model(&database.Prescription{}).
		Where("id = ? AND user_id = ?", prescriptionID, userID).
		Updates(updates).Error
}

func (s *RecordService) DeletePrescription(userID, prescriptionID uuid.UUID) error {
	return s.db.Where("id = ? AND user_id = ?", prescriptionID, userID).
		Delete(&database.Prescription{}).Error
}

// Appointment methods
func (s *RecordService) CreateAppointment(userID uuid.UUID, appointment *database.Appointment) error {
	appointment.UserID = userID
	appointment.ID = uuid.New()
	appointment.CreatedAt = time.Now()
	appointment.UpdatedAt = time.Now()
	return s.db.Create(appointment).Error
}

func (s *RecordService) GetAppointments(userID uuid.UUID, limit, offset int, upcomingOnly bool) ([]database.Appointment, int64, error) {
	var appointments []database.Appointment
	var total int64

	query := s.db.Where("user_id = ?", userID)
	if upcomingOnly {
		query = query.Where("appointment_date >= ? AND is_completed = ?", time.Now(), false)
	}
	s.db.Model(&database.Appointment{}).Where("user_id = ?", userID).Count(&total)

	if err := query.Order("appointment_date ASC").Limit(limit).Offset(offset).Find(&appointments).Error; err != nil {
		return nil, 0, err
	}

	return appointments, total, nil
}

func (s *RecordService) GetAppointmentByID(userID, appointmentID uuid.UUID) (*database.Appointment, error) {
	var appointment database.Appointment
	if err := s.db.Where("id = ? AND user_id = ?", appointmentID, userID).First(&appointment).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (s *RecordService) UpdateAppointment(userID, appointmentID uuid.UUID, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return s.db.Model(&database.Appointment{}).
		Where("id = ? AND user_id = ?", appointmentID, userID).
		Updates(updates).Error
}

func (s *RecordService) DeleteAppointment(userID, appointmentID uuid.UUID) error {
	return s.db.Where("id = ? AND user_id = ?", appointmentID, userID).
		Delete(&database.Appointment{}).Error
}

// Lab Report methods
func (s *RecordService) CreateLabReport(userID uuid.UUID, labReport *database.LabReport) error {
	labReport.UserID = userID
	labReport.ID = uuid.New()
	labReport.CreatedAt = time.Now()
	labReport.UpdatedAt = time.Now()
	return s.db.Create(labReport).Error
}

func (s *RecordService) GetLabReports(userID uuid.UUID, limit, offset int) ([]database.LabReport, int64, error) {
	var labReports []database.LabReport
	var total int64

	query := s.db.Where("user_id = ?", userID)
	s.db.Model(&database.LabReport{}).Where("user_id = ?", userID).Count(&total)

	if err := query.Order("test_date DESC").Limit(limit).Offset(offset).Find(&labReports).Error; err != nil {
		return nil, 0, err
	}

	return labReports, total, nil
}

func (s *RecordService) GetLabReportByID(userID, labReportID uuid.UUID) (*database.LabReport, error) {
	var labReport database.LabReport
	if err := s.db.Where("id = ? AND user_id = ?", labReportID, userID).First(&labReport).Error; err != nil {
		return nil, err
	}
	return &labReport, nil
}

func (s *RecordService) UpdateLabReport(userID, labReportID uuid.UUID, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return s.db.Model(&database.LabReport{}).
		Where("id = ? AND user_id = ?", labReportID, userID).
		Updates(updates).Error
}

func (s *RecordService) DeleteLabReport(userID, labReportID uuid.UUID) error {
	return s.db.Where("id = ? AND user_id = ?", labReportID, userID).
		Delete(&database.LabReport{}).Error
}

// Health Insurance methods
func (s *RecordService) CreateHealthInsurance(userID uuid.UUID, insurance *database.HealthInsurance) error {
	insurance.UserID = userID
	insurance.ID = uuid.New()
	insurance.CreatedAt = time.Now()
	insurance.UpdatedAt = time.Now()
	return s.db.Create(insurance).Error
}

func (s *RecordService) GetHealthInsurances(userID uuid.UUID) ([]database.HealthInsurance, error) {
	var insurances []database.HealthInsurance
	if err := s.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&insurances).Error; err != nil {
		return nil, err
	}
	return insurances, nil
}

func (s *RecordService) UpdateHealthInsurance(userID, insuranceID uuid.UUID, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return s.db.Model(&database.HealthInsurance{}).
		Where("id = ? AND user_id = ?", insuranceID, userID).
		Updates(updates).Error
}

func (s *RecordService) DeleteHealthInsurance(userID, insuranceID uuid.UUID) error {
	return s.db.Where("id = ? AND user_id = ?", insuranceID, userID).
		Delete(&database.HealthInsurance{}).Error
}

