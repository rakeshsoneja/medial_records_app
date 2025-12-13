package services

import (
	"medical-records-app/internal/database"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicationService struct {
	db *gorm.DB
}

func NewMedicationService(db *gorm.DB) *MedicationService {
	return &MedicationService{db: db}
}

func (s *MedicationService) CreateMedication(userID uuid.UUID, medication *database.Medication) error {
	medication.UserID = userID
	medication.ID = uuid.New()
	medication.CreatedAt = time.Now()
	medication.UpdatedAt = time.Now()
	return s.db.Create(medication).Error
}

func (s *MedicationService) GetMedications(userID uuid.UUID, activeOnly bool) ([]database.Medication, error) {
	var medications []database.Medication
	query := s.db.Where("user_id = ?", userID)
	
	if activeOnly {
		query = query.Where("is_active = ?", true)
	}
	
	if err := query.Order("created_at DESC").Find(&medications).Error; err != nil {
		return nil, err
	}
	return medications, nil
}

func (s *MedicationService) GetMedicationByID(userID, medicationID uuid.UUID) (*database.Medication, error) {
	var medication database.Medication
	if err := s.db.Where("id = ? AND user_id = ?", medicationID, userID).First(&medication).Error; err != nil {
		return nil, err
	}
	return &medication, nil
}

func (s *MedicationService) UpdateMedication(userID, medicationID uuid.UUID, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return s.db.Model(&database.Medication{}).
		Where("id = ? AND user_id = ?", medicationID, userID).
		Updates(updates).Error
}

func (s *MedicationService) DeleteMedication(userID, medicationID uuid.UUID) error {
	return s.db.Where("id = ? AND user_id = ?", medicationID, userID).
		Delete(&database.Medication{}).Error
}

func (s *MedicationService) GetMedicationsNeedingRefill(userID uuid.UUID) ([]database.Medication, error) {
	var medications []database.Medication
	now := time.Now()
	
	// Find medications where next refill date is within the reminder days
	if err := s.db.Where("user_id = ? AND is_active = ?", userID, true).
		Where("next_refill_date IS NOT NULL AND next_refill_date <= ?", 
			now.AddDate(0, 0, 7)). // Default 7 days ahead
		Find(&medications).Error; err != nil {
		return nil, err
	}
	
	return medications, nil
}

