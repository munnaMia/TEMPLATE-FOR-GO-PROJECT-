package config

import "github.com/joho/godotenv"

var configurations *Configuration

type Configuration struct{
	Service Service
}

type Service struct {
	Name      string
	Version   string
	HTTP_Port int
}

func loadConfiguration() {
	err := godotenv.Load()
	if err != nil {
		// msg and os exit 1
	}

	// load the envs.
	// write the envs value on configurations variable
}

func GetConfiguration() *Configuration {
	if configurations == nil {
		loadConfiguration()
	}

	return configurations
}
