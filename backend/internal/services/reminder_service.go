package services

import (
	"medical-records-app/internal/database"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReminderService struct {
	db *gorm.DB
}

func NewReminderService(db *gorm.DB) *ReminderService {
	return &ReminderService{db: db}
}

func (s *ReminderService) CreateReminder(userID uuid.UUID, reminder *database.Reminder) error {
	reminder.UserID = userID
	reminder.ID = uuid.New()
	reminder.CreatedAt = time.Now()
	reminder.UpdatedAt = time.Now()
	return s.db.Create(reminder).Error
}

func (s *ReminderService) GetReminders(userID uuid.UUID, upcomingOnly bool) ([]database.Reminder, error) {
	var reminders []database.Reminder
	query := s.db.Where("user_id = ?", userID)
	
	if upcomingOnly {
		query = query.Where("reminder_date >= ? AND is_completed = ?", time.Now(), false)
	}
	
	if err := query.Order("reminder_date ASC").Find(&reminders).Error; err != nil {
		return nil, err
	}
	return reminders, nil
}

func (s *ReminderService) GetReminderByID(userID, reminderID uuid.UUID) (*database.Reminder, error) {
	var reminder database.Reminder
	if err := s.db.Where("id = ? AND user_id = ?", reminderID, userID).First(&reminder).Error; err != nil {
		return nil, err
	}
	return &reminder, nil
}

func (s *ReminderService) UpdateReminder(userID, reminderID uuid.UUID, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return s.db.Model(&database.Reminder{}).
		Where("id = ? AND user_id = ?", reminderID, userID).
		Updates(updates).Error
}

func (s *ReminderService) DeleteReminder(userID, reminderID uuid.UUID) error {
	return s.db.Where("id = ? AND user_id = ?", reminderID, userID).
		Delete(&database.Reminder{}).Error
}

func (s *ReminderService) GetUpcomingReminders(userID uuid.UUID, daysAhead int) ([]database.Reminder, error) {
	var reminders []database.Reminder
	cutoffDate := time.Now().AddDate(0, 0, daysAhead)
	
	if err := s.db.Where("user_id = ? AND is_completed = ?", userID, false).
		Where("reminder_date >= ? AND reminder_date <= ?", time.Now(), cutoffDate).
		Order("reminder_date ASC").
		Find(&reminders).Error; err != nil {
		return nil, err
	}
	
	return reminders, nil
}

