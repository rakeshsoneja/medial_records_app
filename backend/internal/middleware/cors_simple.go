package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// SimpleCORSMiddleware provides a simple, guaranteed-to-work CORS solution
// This is more permissive but ensures CORS errors are resolved
func SimpleCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		// Get allowed origins from environment
		frontendURL := os.Getenv("FRONTEND_URL")
		allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")
		
		// Build allowed origins list
		var allowedOrigins []string
		if frontendURL != "" {
			allowedOrigins = append(allowedOrigins, frontendURL)
		}
		if allowedOriginsEnv != "" {
			origins := strings.Split(allowedOriginsEnv, ",")
			for _, o := range origins {
				o = strings.TrimSpace(o)
				if o != "" {
					allowedOrigins = append(allowedOrigins, o)
				}
			}
		}
		
		// Determine which origin to allow
		allowOrigin := ""
		if origin != "" {
			// Check if origin is in allowed list
			for _, allowed := range allowedOrigins {
				if origin == allowed || origin == strings.TrimSuffix(allowed, "/") || origin == allowed+"/" {
					allowOrigin = origin
					break
				}
			}
			// If not in list but we have allowed origins, use first one
			// Otherwise, allow the requesting origin (permissive mode)
			if allowOrigin == "" {
				if len(allowedOrigins) > 0 {
					allowOrigin = allowedOrigins[0]
					log.Printf("CORS: Origin %s not in allowed list, using %s", origin, allowOrigin)
				} else {
					allowOrigin = origin
					log.Printf("CORS: No allowed origins configured, allowing requesting origin: %s", origin)
				}
			}
		} else if len(allowedOrigins) > 0 {
			allowOrigin = allowedOrigins[0]
		} else {
			// Fallback: allow all (use * but this won't work with credentials)
			allowOrigin = "*"
		}
		
		// Set CORS headers
		if allowOrigin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept, Authorization, X-Requested-With, X-CSRF-Token, Cache-Control")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "43200") // 12 hours
		
		// Handle OPTIONS preflight
		if c.Request.Method == "OPTIONS" {
			log.Printf("CORS: Handling OPTIONS preflight - Origin: %s, Allow-Origin: %s", origin, allowOrigin)
			c.AbortWithStatus(204)
			return
		}
		
		log.Printf("CORS: Request - Origin: %s, Allow-Origin: %s, Method: %s, Path: %s", 
			origin, allowOrigin, c.Request.Method, c.Request.URL.Path)
		
		c.Next()
	}
}

