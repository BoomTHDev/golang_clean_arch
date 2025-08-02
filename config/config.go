package config

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server   *Server   `validate:"required"`
		Database *Database `validate:"required"`
		Redis    *Redis    `validate:"required"`
	}

	Server struct {
		Port         int           `validate:"required"`
		AllowOrigins []string      `validate:"required"`
		BodyLimit    int           `validate:"required"`
		TimeOut      time.Duration `validate:"required"`
		JWTSecret    string        `validate:"required"`
	}

	Database struct {
		Host     string `validate:"required"`
		Port     int    `validate:"required"`
		User     string `validate:"required"`
		Password string `validate:"required"`
		DBName   string `validate:"required"`
		SSLMode  string `validate:"required"`
		Schema   string `validate:"required"`
	}

	Redis struct {
		URL      string
		Host     string
		Port     int
		Password string
		DB       int
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		godotenv.Load()

		configInstance = &Config{
			Server:   &Server{},
			Database: &Database{},
			Redis:    &Redis{},
		}

		// Server
		port, err := strconv.Atoi(getEnv("SERVER_PORT", "10000"))
		if err != nil {
			port = 10000
		}
		bodyLimit, err := strconv.Atoi(getEnv("SERVER_BODY_LIMIT", "10485760")) // 10MB
		if err != nil {
			bodyLimit = 10485760
		}
		timeOut, err := time.ParseDuration(getEnv("SERVER_TIMEOUT", "30s"))
		if err != nil {
			timeOut = 30 * time.Second
		}
		configInstance.Server = &Server{
			Port:         port,
			AllowOrigins: strings.Split(getEnv("SERVER_ALLOW_ORIGINS", "*"), ","),
			BodyLimit:    bodyLimit,
			TimeOut:      timeOut,
			JWTSecret:    getEnv("SERVER_JWT_SECRET", "secret"),
		}

		// Database
		dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
		if err != nil {
			dbPort = 5432
		}
		configInstance.Database = &Database{
			Host:     getEnv("DB_HOST", ""),
			Port:     dbPort,
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", ""),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
			Schema:   getEnv("DB_SCHEMA", "public"),
		}

		// Redis
		redisPort, err := strconv.Atoi(getEnv("REDIS_PORT", "6379"))
		if err != nil {
			redisPort = 6379
		}
		redisDb, err := strconv.Atoi(getEnv("REDIS_DB", "0"))
		if err != nil {
			redisDb = 0
		}
		configInstance.Redis = &Redis{
			URL:      getEnv("REDIS_URL", ""),
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     redisPort,
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       redisDb,
		}
	})

	validating := validator.New()

	if err := validating.Struct(configInstance); err != nil {
		panic(err)
	}

	return configInstance

}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
