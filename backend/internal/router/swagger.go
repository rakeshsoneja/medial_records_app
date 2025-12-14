package router

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupSwagger configures Swagger documentation
func SetupSwagger(r *gin.Engine) {
	// Check if docs directory exists - try multiple paths
	docsPath := ""
	possiblePaths := []string{
		"./docs",
		"./backend/docs",
		"../docs",
		"docs",
		"/opt/render/project/src/docs",
		"/opt/render/project/src/backend/docs",
	}

	// Find docs directory
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			// Check if swagger.json exists in this path
			docJSONPath := filepath.Join(path, "swagger.json")
			if _, err := os.Stat(docJSONPath); err == nil {
				docsPath = path
				log.Printf("✅ Swagger docs found at: %s", docsPath)
				log.Printf("✅ Swagger doc.json found at: %s", docJSONPath)
				break
			}
		}
	}

	if docsPath == "" {
		log.Printf("❌ WARNING: Swagger docs not found in any of these paths: %v", possiblePaths)
		log.Printf("❌ Swagger documentation will not be available")
		log.Printf("💡 To generate docs locally, run: cd backend && swag init -g cmd/server/main.go -o ./docs")
		
		// List current directory for debugging
		if wd, err := os.Getwd(); err == nil {
			log.Printf("📁 Current working directory: %s", wd)
			if entries, err := os.ReadDir("."); err == nil {
				log.Printf("📁 Files in current directory:")
				for _, entry := range entries {
					log.Printf("   - %s (dir: %v)", entry.Name(), entry.IsDir())
				}
			}
		}
		
		// Still set up the route but it will return 404/500
		// This way we can see the error in browser
		r.GET("/swagger/*any", func(c *gin.Context) {
			c.JSON(500, gin.H{
				"error": "Swagger documentation not available",
				"message": "Swagger docs were not generated during build. Check build logs for 'swag init' errors.",
				"paths_checked": possiblePaths,
			})
		})
		return
	}

	// Configure Swagger handler with explicit URL
	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Printf("✅ Swagger UI configured at /swagger/index.html")
}

