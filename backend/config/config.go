package config

import (
	"fmt"
	"log"
	"os"

	"reflect"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ServerEnvironment string

const (
	ServerEnvironmentDevelopment ServerEnvironment = "development"
	ServerEnvironmentProduction  ServerEnvironment = "production"
)

type Config struct {
	PORT            string
	DB_NAME         string
	FRONTEND_APPS   string
	SERVER_BASE_URL string
	JWT_SECRET      string
	XRATE_LIMIT_MAX int
	APP_ENV         ServerEnvironment
}

var EnvConfig Config

func InitEnvSchema() *logrus.Logger {
	loadENV()
	loadConfig()
	return initLogger()
}

func loadENV() {
	if goEnv := os.Getenv("GO_ENV"); goEnv == "" {
		if err := godotenv.Load(); err != nil {
			fmt.Println("No .env file found")
		}
	}
}

func loadConfig() {
	envConfigReflection := reflect.ValueOf(&EnvConfig).Elem()
	envConfigType := envConfigReflection.Type()

	for i := 0; i < envConfigReflection.NumField(); i++ {
		field := envConfigType.Field(i)
		fieldName := field.Name
		envVariableValue := os.Getenv(fieldName)

		if envVariableValue == "" {
			log.Fatalf("You must set your %s environment variable.", fieldName)
		}

		switch field.Type.Kind() {
		case reflect.String:
			envConfigReflection.FieldByName(fieldName).SetString(envVariableValue)
		case reflect.Int:
			val, err := strconv.Atoi(envVariableValue)
			if err != nil {
				log.Fatalf("Invalid value for %s: %v", fieldName, err)
			}
			envConfigReflection.FieldByName(fieldName).SetInt(int64(val))
		default:
			log.Fatalf("Unsupported field type %s for field %s", field.Type.Kind(), fieldName)
		}
	}
}

func initLogger() *logrus.Logger {
	Logger := logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return Logger
}
