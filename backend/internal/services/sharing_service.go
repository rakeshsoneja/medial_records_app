package services

import (
	"encoding/json"
	"errors"
	"medical-records-app/internal/database"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SharingService struct {
	db *gorm.DB
}

func NewSharingService(db *gorm.DB) *SharingService {
	return &SharingService{db: db}
}

func (s *SharingService) CreateShareLink(userID uuid.UUID, recordType string, recordIDs []uuid.UUID, expiresInHours int, maxAccessCount int, allowDownload bool, recipientEmail, recipientPhone, shareMethod string) (*database.SharedRecord, error) {
	// Generate unique share token
	shareToken := uuid.New().String()

	// Calculate expiration
	expiresAt := time.Now().Add(time.Duration(expiresInHours) * time.Hour)

	// Convert record IDs to JSON
	recordIDsJSON, err := json.Marshal(recordIDs)
	if err != nil {
		return nil, err
	}

	sharedRecord := &database.SharedRecord{
		ID:                uuid.New(),
		UserID:            userID,
		ShareToken:        shareToken,
		RecordType:        recordType,
		RecordIDs:         string(recordIDsJSON),
		ExpiresAt:         expiresAt,
		MaxAccessCount:    maxAccessCount,
		CurrentAccessCount: 0,
		AllowDownload:     allowDownload,
		RecipientEmail:    recipientEmail,
		RecipientPhone:    recipientPhone,
		ShareMethod:       shareMethod,
		IsActive:          true,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := s.db.Create(sharedRecord).Error; err != nil {
		return nil, err
	}

	return sharedRecord, nil
}

func (s *SharingService) GetSharedRecordByToken(token string) (*database.SharedRecord, error) {
	var sharedRecord database.SharedRecord
	if err := s.db.Where("share_token = ? AND is_active = ?", token, true).First(&sharedRecord).Error; err != nil {
		return nil, err
	}

	// Check if expired
	if time.Now().After(sharedRecord.ExpiresAt) {
		return nil, errors.New("share link has expired")
	}

	// Check access count limit
	if sharedRecord.MaxAccessCount > 0 && sharedRecord.CurrentAccessCount >= sharedRecord.MaxAccessCount {
		return nil, errors.New("share link has reached maximum access count")
	}

	return &sharedRecord, nil
}

func (s *SharingService) RecordAccess(sharedRecordID uuid.UUID, ipAddress, userAgent, action string) error {
	auditLog := &database.AuditLog{
		ID:             uuid.New(),
		SharedRecordID: sharedRecordID,
		IPAddress:      ipAddress,
		UserAgent:      userAgent,
		AccessedAt:     time.Now(),
		Action:         action,
	}

	if err := s.db.Create(auditLog).Error; err != nil {
		return err
	}

	// Increment access count
	return s.db.Model(&database.SharedRecord{}).
		Where("id = ?", sharedRecordID).
		Update("current_access_count", gorm.Expr("current_access_count + 1")).Error
}

func (s *SharingService) GetSharedRecordsByUser(userID uuid.UUID) ([]database.SharedRecord, error) {
	var sharedRecords []database.SharedRecord
	if err := s.db.Where("user_id = ?", userID).
		Preload("AccessLogs").
		Order("created_at DESC").
		Find(&sharedRecords).Error; err != nil {
		return nil, err
	}
	return sharedRecords, nil
}

func (s *SharingService) RevokeShareLink(userID, shareID uuid.UUID) error {
	return s.db.Model(&database.SharedRecord{}).
		Where("id = ? AND user_id = ?", shareID, userID).
		Update("is_active", false).Error
}

func (s *SharingService) GetRecordsByIDs(recordType string, recordIDs []uuid.UUID) (interface{}, error) {
	switch recordType {
	case "prescription":
		var prescriptions []database.Prescription
		if err := s.db.Where("id IN ?", recordIDs).Find(&prescriptions).Error; err != nil {
			return nil, err
		}
		return prescriptions, nil
	case "appointment":
		var appointments []database.Appointment
		if err := s.db.Where("id IN ?", recordIDs).Find(&appointments).Error; err != nil {
			return nil, err
		}
		return appointments, nil
	case "lab_report":
		var labReports []database.LabReport
		if err := s.db.Where("id IN ?", recordIDs).Find(&labReports).Error; err != nil {
			return nil, err
		}
		return labReports, nil
	case "bundle":
		// Return all types
		var prescriptions []database.Prescription
		var appointments []database.Appointment
		var labReports []database.LabReport
		
		s.db.Where("id IN ?", recordIDs).Find(&prescriptions)
		s.db.Where("id IN ?", recordIDs).Find(&appointments)
		s.db.Where("id IN ?", recordIDs).Find(&labReports)
		
		return map[string]interface{}{
			"prescriptions": prescriptions,
			"appointments":  appointments,
			"lab_reports":   labReports,
		}, nil
	default:
		return nil, errors.New("invalid record type")
	}
}

