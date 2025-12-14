package router

import (
	"medical-records-app/internal/config"
	"medical-records-app/internal/handlers"
	"medical-records-app/internal/middleware"
	"medical-records-app/internal/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB, cfg *config.Config) *gin.Engine {
	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Apply CORS middleware - using simple middleware for guaranteed CORS fix
	// This middleware properly handles OPTIONS preflight requests
	// Try simple middleware first (more permissive, always works)
	r.Use(middleware.SimpleCORSMiddleware())

	// Initialize services
	userService := services.NewUserService(db)
	recordService := services.NewRecordService(db)
	sharingService := services.NewSharingService(db)
	medicationService := services.NewMedicationService(db)
	reminderService := services.NewReminderService(db)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userService, cfg)
	recordHandler := handlers.NewRecordHandler(recordService)
	sharingHandler := handlers.NewSharingHandler(sharingService)
	dashboardHandler := handlers.NewDashboardHandler(recordService, medicationService, reminderService)
	medicationHandler := handlers.NewMedicationHandler(medicationService)
	reminderHandler := handlers.NewReminderHandler(reminderService)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Explicit OPTIONS handler for all routes (backup for CORS preflight)
	r.OPTIONS("/*path", func(c *gin.Context) {
		// CORS middleware should handle this, but this ensures it works
		c.Status(204)
	})

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := r.Group("/api/v1")
	{
		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/profile", middleware.AuthMiddleware(), authHandler.GetProfile)
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Dashboard
			protected.GET("/dashboard", dashboardHandler.GetDashboard)

			// Prescriptions
			protected.POST("/prescriptions", recordHandler.CreatePrescription)
			protected.GET("/prescriptions", recordHandler.GetPrescriptions)
			protected.GET("/prescriptions/:id", recordHandler.GetPrescription)
			protected.PUT("/prescriptions/:id", recordHandler.UpdatePrescription)
			protected.DELETE("/prescriptions/:id", recordHandler.DeletePrescription)

			// Appointments
			protected.POST("/appointments", recordHandler.CreateAppointment)
			protected.GET("/appointments", recordHandler.GetAppointments)

			// Lab Reports
			protected.POST("/lab-reports", recordHandler.CreateLabReport)
			protected.GET("/lab-reports", recordHandler.GetLabReports)

			// Health Insurance
			protected.POST("/insurance", recordHandler.CreateHealthInsurance)
			protected.GET("/insurance", recordHandler.GetHealthInsurances)

			// Medications
			protected.POST("/medications", medicationHandler.CreateMedication)
			protected.GET("/medications", medicationHandler.GetMedications)
			protected.GET("/medications/refill-needed", medicationHandler.GetMedicationsNeedingRefill)

			// Reminders
			protected.POST("/reminders", reminderHandler.CreateReminder)
			protected.GET("/reminders", reminderHandler.GetReminders)
			protected.GET("/reminders/upcoming", reminderHandler.GetUpcomingReminders)

			// Sharing
			protected.POST("/sharing/create", sharingHandler.CreateShareLink)
			protected.GET("/sharing/my-shares", sharingHandler.GetMySharedRecords)
			protected.POST("/sharing/:id/revoke", sharingHandler.RevokeShareLink)
		}

		// Public share access
		api.GET("/share/:token", sharingHandler.GetSharedRecord)
	}

	return r
}

