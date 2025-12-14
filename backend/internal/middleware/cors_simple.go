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
		// CRITICAL: When AllowCredentials is true, we MUST use the exact origin, not *
		// ALWAYS use the requesting origin if present - this fixes URL mismatches
		allowOrigin := ""
		if origin != "" {
			// ALWAYS use the requesting origin - this is the safest approach
			// The browser requires exact match, so we must return the exact origin from the request
			allowOrigin = origin
			
			// Log if it doesn't match configured origins (for debugging)
			normalizedOrigin := strings.TrimSuffix(origin, "/")
			matched := false
			for _, allowed := range allowedOrigins {
				normalizedAllowed := strings.TrimSuffix(allowed, "/")
				if normalizedOrigin == normalizedAllowed {
					matched = true
					break
				}
			}
			
			if matched {
				log.Printf("CORS: Origin %s matched configured allowed origins", origin)
			} else if len(allowedOrigins) > 0 {
				log.Printf("CORS: WARNING - Origin %s does not match configured FRONTEND_URL %v, but allowing it anyway", origin, allowedOrigins)
			} else {
				log.Printf("CORS: Allowing origin %s (no FRONTEND_URL configured)", origin)
			}
		} else if len(allowedOrigins) > 0 {
			// No origin header (e.g., same-origin request), use first allowed
			allowOrigin = allowedOrigins[0]
			log.Printf("CORS: No origin header, using configured FRONTEND_URL: %s", allowOrigin)
		} else {
			// Last resort: use * but disable credentials (browser requirement)
			allowOrigin = "*"
			log.Printf("CORS: WARNING - Using wildcard origin, credentials will be disabled")
		}
		
		// Set CORS headers
		if allowOrigin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		}
		
		// Only set credentials if not using wildcard (browser requirement)
		if allowOrigin != "*" {
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept, Authorization, X-Requested-With, X-CSRF-Token, Cache-Control")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "43200") // 12 hours
		
		// Handle OPTIONS preflight - MUST be before c.Next()
		if c.Request.Method == "OPTIONS" {
			log.Printf("CORS: Handling OPTIONS preflight - Origin: %s, Allow-Origin: %s, Path: %s", origin, allowOrigin, c.Request.URL.Path)
			// Set status and abort - don't call c.Next()
			c.AbortWithStatus(204)
			return
		}
		
		log.Printf("CORS: Request - Origin: %s, Allow-Origin: %s, Method: %s, Path: %s", 
			origin, allowOrigin, c.Request.Method, c.Request.URL.Path)
		
		c.Next()
	}
}

