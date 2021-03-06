package main

import (
	"os"
	"time"

	"github.com/kumarabd/go-http/api/http"
	"github.com/kumarabd/go-http/internal/platform/configs"
	"github.com/kumarabd/go-http/internal/platform/logger"
)

var (
	ApplicationName = "test"
)

func main() {

	// Initialize Logger instance
	log := logger.New(ApplicationName)
	log.Info("Logger Initialized")

	// Initialize application specific configs and dependencies
	// App and request config
	appconfig, err := configs.NewConfig()
	if err != nil {
		log.Err("Config Init Failed", err.Error())
		os.Exit(1)
	}
	log.Info("Config Initiliazed")

	// Server Initialization
	service := &http.Service{
		StartedAt: time.Now(),
	}
	err = appconfig.Get(&service)
	if err != nil {
		log.Err("Error getting application config", err.Error())
		os.Exit(1)
	}

	log.Info("Server Started")
	err = http.Start(service)
	if err != nil {
		log.Err("Server crashed!!", err.Error())
		os.Exit(1)
	}
}
