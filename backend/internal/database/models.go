package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a patient or user of the system
type User struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Email             string    `gorm:"uniqueIndex;not null" json:"email"`
	Phone             string    `gorm:"index" json:"phone"`
	PasswordHash      string    `gorm:"not null" json:"-"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	DateOfBirth       *time.Time `json:"date_of_birth"`
	IsEmailVerified   bool      `gorm:"default:false" json:"is_email_verified"`
	IsPhoneVerified   bool      `gorm:"default:false" json:"is_phone_verified"`
	Role              string    `gorm:"default:patient" json:"role"` // patient, doctor, caregiver
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	HealthInsurances  []HealthInsurance `gorm:"foreignKey:UserID" json:"health_insurances,omitempty"`
	Prescriptions     []Prescription    `gorm:"foreignKey:UserID" json:"prescriptions,omitempty"`
	Appointments      []Appointment     `gorm:"foreignKey:UserID" json:"appointments,omitempty"`
	LabReports        []LabReport       `gorm:"foreignKey:UserID" json:"lab_reports,omitempty"`
	Medications       []Medication      `gorm:"foreignKey:UserID" json:"medications,omitempty"`
	Reminders         []Reminder        `gorm:"foreignKey:UserID" json:"reminders,omitempty"`
	SharedRecords     []SharedRecord    `gorm:"foreignKey:UserID" json:"shared_records,omitempty"`
}

// HealthInsurance stores insurance information
type HealthInsurance struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID            uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	InsuranceProvider string    `gorm:"not null" json:"insurance_provider"`
	PolicyNumber      string    `gorm:"not null" json:"policy_number"`
	GroupNumber       string    `json:"group_number"`
	MemberID          string    `json:"member_id"`
	EffectiveDate     *Date     `gorm:"type:date" json:"effective_date"`
	ExpirationDate    *Date     `gorm:"type:date" json:"expiration_date"`
	Notes             string    `gorm:"type:text" json:"notes"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	User              User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// Prescription represents a prescription record
type Prescription struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID            uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	MedicineName      string    `gorm:"not null" json:"medicine_name"`
	Dosage            string    `json:"dosage"`
	Instructions      string    `gorm:"type:text" json:"instructions"`
	PrescribingDoctor string    `json:"prescribing_doctor"`
	DoctorSpecialty   string    `json:"doctor_specialty"`
	Hospital          string    `json:"hospital"`
	PrescriptionDate  Date      `gorm:"type:date;not null" json:"prescription_date"`
	AttachmentURL     string    `json:"attachment_url"` // S3 URL for PDF/photo
	AttachmentType    string    `json:"attachment_type"` // pdf, jpg, png
	IsActive          bool      `gorm:"default:true" json:"is_active"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	User              User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// Appointment represents a medical appointment
type Appointment struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID            uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	DoctorName        string    `gorm:"not null" json:"doctor_name"`
	Specialty         string    `json:"specialty"`
	Hospital          string    `json:"hospital"`
	Location          string    `json:"location"`
	AppointmentDate   DateTime  `gorm:"type:timestamp;not null;index" json:"appointment_date"`
	Notes             string    `gorm:"type:text" json:"notes"`
	IsCompleted       bool      `gorm:"default:false" json:"is_completed"`
	ReminderSent      bool      `gorm:"default:false" json:"reminder_sent"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	User              User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// LabReport represents a lab test report
type LabReport struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID            uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	TestType          string    `gorm:"not null" json:"test_type"`
	LabName           string    `json:"lab_name"`
	TestDate          Date      `gorm:"type:date;not null;index" json:"test_date"`
	ReportURL         string    `gorm:"not null" json:"report_url"` // S3 URL
	ReportType        string    `json:"report_type"` // pdf, jpg, png
	Notes             string    `gorm:"type:text" json:"notes"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	User              User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// Medication represents regular medications tracked by pharmacy
type Medication struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID            uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	MedicineName      string    `gorm:"not null" json:"medicine_name"`
	Dosage            string    `json:"dosage"`
	Frequency         string    `json:"frequency"` // daily, twice daily, etc.
	PharmacyName      string    `json:"pharmacy_name"`
	PharmacyPhone     string    `json:"pharmacy_phone"`
	PharmacyAddress   string    `json:"pharmacy_address"`
	LastRefillDate    *Date     `gorm:"type:date" json:"last_refill_date"`
	NextRefillDate    *Date     `gorm:"type:date" json:"next_refill_date"`
	RefillReminderDays int      `gorm:"default:7" json:"refill_reminder_days"`
	IsActive          bool      `gorm:"default:true" json:"is_active"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	User              User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// Reminder represents health check-up reminders
type Reminder struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID            uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Title             string    `gorm:"not null" json:"title"`
	Description       string    `gorm:"type:text" json:"description"`
	ReminderDate      DateTime  `gorm:"type:timestamp;not null;index" json:"reminder_date"`
	ReminderType      string    `json:"reminder_type"` // checkup, vaccination, test, etc.
	IsCompleted       bool      `gorm:"default:false" json:"is_completed"`
	IsRecurring       bool      `gorm:"default:false" json:"is_recurring"`
	RecurrenceInterval string   `json:"recurrence_interval"` // monthly, quarterly, yearly
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	User              User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// SharedRecord represents a shared medical record with time-limited access
type SharedRecord struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID            uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	ShareToken        string    `gorm:"uniqueIndex;not null" json:"share_token"`
	RecordType        string    `gorm:"not null" json:"record_type"` // prescription, appointment, lab_report, bundle
	RecordIDs         string    `gorm:"type:text" json:"record_ids"` // JSON array of record IDs
	ExpiresAt         time.Time `gorm:"not null;index" json:"expires_at"`
	MaxAccessCount    int       `gorm:"default:0" json:"max_access_count"` // 0 = unlimited
	CurrentAccessCount int      `gorm:"default:0" json:"current_access_count"`
	AllowDownload     bool      `gorm:"default:false" json:"allow_download"`
	RecipientEmail    string    `json:"recipient_email"`
	RecipientPhone    string    `json:"recipient_phone"`
	ShareMethod       string    `json:"share_method"` // email, sms, link
	IsActive          bool      `gorm:"default:true" json:"is_active"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	User              User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	AccessLogs        []AuditLog `gorm:"foreignKey:SharedRecordID" json:"access_logs,omitempty"`
}

// AuditLog tracks access to shared records
type AuditLog struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	SharedRecordID    uuid.UUID `gorm:"type:uuid;not null;index" json:"shared_record_id"`
	IPAddress         string    `json:"ip_address"`
	UserAgent         string    `json:"user_agent"`
	AccessedAt        time.Time `gorm:"not null" json:"accessed_at"`
	Action            string    `json:"action"` // viewed, downloaded

	SharedRecord      SharedRecord `gorm:"foreignKey:SharedRecordID" json:"shared_record,omitempty"`
}

