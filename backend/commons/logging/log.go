package logging

import (
	"log"

	"go.uber.org/zap"
)

var AppLogger *zap.Logger

func LogInit() {
	var err error
	AppLogger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
}

// Call this at shutdown to prevent file descriptor leaks
func ShutdownLogger() {
	if AppLogger != nil {
		err := AppLogger.Sync()
		if err != nil {
			log.Printf("Error syncing logger: %v", err)
		}
	}
}
