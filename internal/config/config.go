package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Configuration

type Configuration struct {
	Service Service
}

type Service struct {
	Name      string
	Version   string
	HTTP_Port int
}

// load all the configuration from the env file.
func loadConfiguration() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("No .env file found falling back to system environment variables")
		os.Exit(1)
	}

	svrName := os.Getenv("SERVICE_NAME")
	if svrName == "" {
		slog.Error("Service name not found")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		slog.Error("Service name not found")
		os.Exit(1)
	}

	httpPort, err := strconv.ParseInt(os.Getenv("HTTP_PORT"), 10, 32)
	if err != nil {
		slog.Error("Could not able to parsing HTTP port address")
		os.Exit(1)
	}

	service := Service{
		Name:      svrName,
		Version:   version,
		HTTP_Port: int(httpPort),
	}

	// keeping the configurations
	configurations = &Configuration{
		Service: service,
	}
}

// return the confis from the env file. if need loads them.
func GetConfiguration() *Configuration {
	if configurations == nil {
		loadConfiguration()
	}

	return configurations
}
