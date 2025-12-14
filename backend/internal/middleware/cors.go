package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GetCORSConfig returns a CORS configuration based on environment
// This handles both development and production scenarios, including Render deployments
func GetCORSConfig() cors.Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// Get frontend URL from environment (Render provides this)
	frontendURL := os.Getenv("FRONTEND_URL")
	
	// Get allowed origins from environment (comma-separated list)
	allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")
	
	config := cors.Config{
		// Allow common HTTP methods
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"PATCH",
			"OPTIONS",
			"HEAD",
		},
		
		// Allow common headers
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Content-Length",
			"Accept",
			"Authorization",
			"X-Requested-With",
			"X-CSRF-Token",
			"Cache-Control",
		},
		
		// Expose headers that frontend might need
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Type",
			"Authorization",
		},
		
		// Allow credentials (cookies, authorization headers)
		AllowCredentials: true,
		
		// Cache preflight requests for 12 hours
		MaxAge: 12 * 3600,
	}

	// Production configuration
	if env == "production" {
		// Build allowed origins list
		var allowedOrigins []string
		
		// Add frontend URL if provided
		if frontendURL != "" {
			// Handle both http and https
			allowedOrigins = append(allowedOrigins, frontendURL)
			// Also add without trailing slash
			if strings.HasSuffix(frontendURL, "/") {
				allowedOrigins = append(allowedOrigins, strings.TrimSuffix(frontendURL, "/"))
			} else {
				allowedOrigins = append(allowedOrigins, frontendURL+"/")
			}
		}
		
		// Add any additional allowed origins from environment
		if allowedOriginsEnv != "" {
			origins := strings.Split(allowedOriginsEnv, ",")
			for _, origin := range origins {
				origin = strings.TrimSpace(origin)
				if origin != "" {
					allowedOrigins = append(allowedOrigins, origin)
				}
			}
		}
		
		if len(allowedOrigins) > 0 {
			// Use explicit allowed origins (more secure)
			config.AllowOrigins = allowedOrigins
			log.Printf("CORS: Production mode with allowed origins: %v", allowedOrigins)
		} else {
			// Fallback: Allow all origins (less secure but works for Render)
			// This is a fallback - ideally FRONTEND_URL should be set
			config.AllowOriginFunc = func(origin string) bool {
				// Log the origin for debugging
				log.Printf("CORS: Allowing origin (fallback mode): %s", origin)
				return true
			}
			log.Println("CORS: Production mode - WARNING: Allowing all origins (FRONTEND_URL not set)")
		}
	} else {
		// Development configuration
		allowedOrigins := []string{
			"http://localhost:3000",
			"http://localhost:3001",
			"http://127.0.0.1:3000",
			"http://127.0.0.1:3001",
		}
		
		// Add frontend URL if provided (for testing production frontend locally)
		if frontendURL != "" {
			allowedOrigins = append(allowedOrigins, frontendURL)
		}
		
		// Add any additional allowed origins from environment
		if allowedOriginsEnv != "" {
			origins := strings.Split(allowedOriginsEnv, ",")
			for _, origin := range origins {
				origin = strings.TrimSpace(origin)
				if origin != "" {
					allowedOrigins = append(allowedOrigins, origin)
				}
			}
		}
		
		config.AllowOrigins = allowedOrigins
		log.Printf("CORS: Development mode with allowed origins: %v", allowedOrigins)
	}

	return config
}

// CORSMiddleware returns a Gin middleware handler for CORS
// This is a wrapper that ensures OPTIONS requests are handled correctly
func CORSMiddleware() gin.HandlerFunc {
	config := GetCORSConfig()
	
	// Create the CORS middleware
	corsHandler := cors.New(config)
	
	// Return a custom handler that ensures OPTIONS requests are handled
	return func(c *gin.Context) {
		// Handle preflight OPTIONS requests
		if c.Request.Method == "OPTIONS" {
			// Let the CORS middleware handle it
			corsHandler(c)
			// If the middleware didn't abort, we should return 204
			if !c.IsAborted() {
				c.Status(204)
			}
			return
		}
		
		// For all other requests, apply CORS headers
		corsHandler(c)
		
		// Continue to next handler
		if !c.IsAborted() {
			c.Next()
		}
	}
}

