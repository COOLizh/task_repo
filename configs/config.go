// Package config returns configuration parameters
package config

import "os"

// Config struct
type Config struct {
	DatabaseURL     string
	APIPort         string
	FileLogName     string
	JwtSalt         string
	ExecutionerPort string
}

// New function returns new struct
func New() *Config {
	return &Config{
		DatabaseURL:     getEnv("DATABASE_URL"),
		APIPort:         getEnv("APIPort"),
		FileLogName:     getEnv("FileLogName"),
		JwtSalt:         getEnv("JWT_SALT"),
		ExecutionerPort: getEnv("EXECUTIONER_PORT"),
	}
}

// Function return the value from env.example
func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return ""
}
