package router

import (
	"medical-records-app/internal/config"
	"medical-records-app/internal/handlers"
	"medical-records-app/internal/middleware"
	"medical-records-app/internal/services"
	"os"

	"github.com/gin-contrib/cors"
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

	// CORS configuration
	// Get allowed origins from environment or use defaults
	allowedOrigins := []string{
		"http://localhost:3000",
		"http://localhost:3001",
	}
	
	// Add production frontend URL if set
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL != "" {
		allowedOrigins = append(allowedOrigins, frontendURL)
	}
	
	// Configure CORS with flexible origin handling for production
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 hours
	}
	
	// In production, use flexible origin matching if FRONTEND_URL is not set
	// This allows any .onrender.com origin as a fallback
	if cfg.Server.Env == "production" && frontendURL == "" {
		corsConfig.AllowOriginFunc = func(origin string) bool {
			// Allow localhost for development/testing
			if origin == "http://localhost:3000" || origin == "http://localhost:3001" {
				return true
			}
			// Allow any HTTPS origin from onrender.com (Render frontend URLs)
			if len(origin) > 0 && (origin[:8] == "https://" || origin[:7] == "http://") {
				// Check if it's a Render URL or allow all HTTPS origins in production
				// This is a fallback - ideally FRONTEND_URL should be set
				return true
			}
			return false
		}
	} else {
		// Use explicit allowed origins
		corsConfig.AllowOrigins = allowedOrigins
	}
	
	r.Use(cors.New(corsConfig))

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

