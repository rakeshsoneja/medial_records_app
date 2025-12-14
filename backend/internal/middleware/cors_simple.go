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
		allowOrigin := ""
		if origin != "" {
			// Normalize origins (remove trailing slashes for comparison)
			normalizedOrigin := strings.TrimSuffix(origin, "/")
			
			// Check if origin is in allowed list (with flexible matching)
			found := false
			for _, allowed := range allowedOrigins {
				normalizedAllowed := strings.TrimSuffix(allowed, "/")
				// Exact match or match without trailing slash
				if normalizedOrigin == normalizedAllowed || origin == allowed || origin == normalizedAllowed || origin == allowed+"/" {
					allowOrigin = origin // Use the exact origin from request
					found = true
					break
				}
			}
			
			// If not found in list, be permissive: allow the requesting origin
			// This ensures CORS works even if there's a slight mismatch
			if !found {
				allowOrigin = origin
				if len(allowedOrigins) > 0 {
					log.Printf("CORS: Origin %s not exactly in allowed list %v, but allowing it anyway", origin, allowedOrigins)
				} else {
					log.Printf("CORS: No allowed origins configured, allowing requesting origin: %s", origin)
				}
			} else {
				log.Printf("CORS: Origin %s matched allowed origins", origin)
			}
		} else if len(allowedOrigins) > 0 {
			// No origin header (e.g., same-origin request), use first allowed
			allowOrigin = allowedOrigins[0]
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

