package services

import (
	"errors"
	"medical-records-app/internal/auth"
	"medical-records-app/internal/database"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Register(email, password, firstName, lastName, phone string) (*database.User, error) {
	// Check if user already exists
	var existingUser database.User
	if err := s.db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &database.User{
		ID:           uuid.New(),
		Email:        email,
		Phone:        phone,
		PasswordHash: hashedPassword,
		FirstName:    firstName,
		LastName:     lastName,
		Role:         "patient",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(email, password string) (*database.User, error) {
	var user database.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if !auth.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func (s *UserService) GetUserByID(userID uuid.UUID) (*database.User, error) {
	var user database.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) UpdateUser(userID uuid.UUID, updates map[string]interface{}) error {
	return s.db.Model(&database.User{}).Where("id = ?", userID).Updates(updates).Error
}

