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
	// Check if docs directory exists
	docsPath := "./docs"
	if _, err := os.Stat(docsPath); os.IsNotExist(err) {
		log.Printf("WARNING: Swagger docs directory not found at %s", docsPath)
		// Try alternative paths
		altPaths := []string{"./backend/docs", "../docs", "docs"}
		for _, altPath := range altPaths {
			if _, err := os.Stat(altPath); err == nil {
				docsPath = altPath
				log.Printf("Found Swagger docs at: %s", docsPath)
				break
			}
		}
	} else {
		log.Printf("Swagger docs found at: %s", docsPath)
	}

	// Check if doc.json exists
	docJSONPath := filepath.Join(docsPath, "swagger.json")
	if _, err := os.Stat(docJSONPath); os.IsNotExist(err) {
		log.Printf("ERROR: Swagger doc.json not found at %s", docJSONPath)
		log.Printf("Swagger documentation will not be available")
		log.Printf("To generate docs, run: swag init -g cmd/server/main.go -o ./docs")
		return
	}

	log.Printf("Swagger doc.json found at: %s", docJSONPath)

	// Configure Swagger handler
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

