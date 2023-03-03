package configs

import (
	"log"
	"os"

	"github.com/naofel1/api-golang-template/internal/primitive"
)

// Init will get the application mode from environment variable
func Init() *Config {
	// Check application mode
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		log.Println("Application mode not set, switching to default 'prod'")

		mode = primitive.AppModeProd.String()
	}

	conf := &Config{
		AppInfo: &AppInfo{
			Mode: mode,
		},
	}

	return conf
}
