package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Configuration

type Configuration struct {
	Service  Service
	Database DB
}

// service configuration
type Service struct {
	Name      string
	Version   string
	HTTP_Port int
	SecretKey string
}

// database configuration
type DB struct {
	DB_Host     string
	DB_Port     string
	DB_User     string
	DB_Password string
	DB_Name     string
	SSL_Mode    string
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

	secretKey := os.Getenv("SECRET_KEY")
	if version == "" {
		slog.Error("secret key not found")
		os.Exit(1)
	}

	service := Service{
		Name:      svrName,
		Version:   version,
		HTTP_Port: int(httpPort),
		SecretKey: secretKey,
	}

	// fetching db configurations
	dbHost := os.Getenv("DB_HOST")
	if svrName == "" {
		slog.Error("db host string not found")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if svrName == "" {
		slog.Error("db port  not found")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if svrName == "" {
		slog.Error("db user not found")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if svrName == "" {
		slog.Error("db password not found")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if svrName == "" {
		slog.Error("db name not found")
		os.Exit(1)
	}

	sslMode := os.Getenv("SSL_MODE")
	if svrName == "" {
		slog.Error("sslmode not found")
		os.Exit(1)
	}

	database := DB{
		DB_Host:     dbHost,
		DB_Port:     dbPort,
		DB_User:     dbUser,
		DB_Password: dbPassword,
		DB_Name:     dbName,
		SSL_Mode:    sslMode,
	}

	// keeping the configurations
	configurations = &Configuration{
		Service:  service,
		Database: database,
	}
}

// return the confis from the env file. if need loads them.
func GetConfiguration() *Configuration {
	if configurations == nil {
		loadConfiguration()
	}

	return configurations
}
