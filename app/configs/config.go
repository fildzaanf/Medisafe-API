package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MYSQL        MySQLConfig
	REDIS        RedisConfig
	CLOUDSTORAGE CloudStorageConfig
	FIREBASE     FirebaseConfig
	MIDTRANS     MidtransConfig
	SMTP         SMTPConfig
	OPENAI       OpenAIConfig
}

type MySQLConfig struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
}

type RedisConfig struct {
	REDIS_HOST string
	REDIS_PORT string
}

type CloudStorageConfig struct {
	GOOGLE_APPLICATION_CREDENTIALS string
	CLOUD_PROJECT_ID               string
	CLOUD_BUCKET_NAME              string
}

type FirebaseConfig struct {
	FIREBASE_API_KEY          string
	FIREBASE_CREDENTIALS_FILE string
}

type MidtransConfig struct {
	MIDTRANS_SERVER_KEY string
	MIDTRANS_CLIENT_KEY string
}

type OpenAIConfig struct {
	OPENAI_KEY string
}

type SMTPConfig struct {
	SMTP_USER string
	SMTP_PASS string
	SMTP_PORT string
	SMTP_HOST string
}

func LoadConfig() (*AppConfig, error) {

	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("failed to load environment variables from .env file: %w", err)
		}
	} else {
		fmt.Println("warning: .env file not found. make sure environment variables are set")
	}

	return &AppConfig{
		MYSQL: MySQLConfig{
			DB_USER: os.Getenv("DB_USER"),
			DB_PASS: os.Getenv("DB_PASS"),
			DB_HOST: os.Getenv("DB_HOST"),
			DB_PORT: os.Getenv("DB_PORT"),
			DB_NAME: os.Getenv("DB_NAME"),
		},
		REDIS: RedisConfig{
			REDIS_HOST: os.Getenv("REDIS_HOST"),
			REDIS_PORT: os.Getenv("REDIS_PORT"),
		},
		CLOUDSTORAGE: CloudStorageConfig{
			GOOGLE_APPLICATION_CREDENTIALS: os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
			CLOUD_PROJECT_ID:               os.Getenv("CLOUD_PROJECT_ID"),
			CLOUD_BUCKET_NAME:              os.Getenv("CLOUD_BUCKET_NAME"),
		},
		FIREBASE: FirebaseConfig{
			FIREBASE_API_KEY:          os.Getenv("FIREBASE_API_KEY"),
			FIREBASE_CREDENTIALS_FILE: os.Getenv("FIREBASE_CREDENTIALS_FILE"),
		},
		MIDTRANS: MidtransConfig{
			MIDTRANS_SERVER_KEY: os.Getenv("MIDTRANS_SERVER_KEY"),
			MIDTRANS_CLIENT_KEY: os.Getenv("MIDTRANS_CLIENT_KEY"),
		},
		SMTP: SMTPConfig{
			SMTP_USER: os.Getenv("SMTP_USER"),
			SMTP_PASS: os.Getenv("SMTP_PASS"),
			SMTP_PORT: os.Getenv("SMTP_PORT"),
			SMTP_HOST: os.Getenv("SMTP_HOST"),
		},
		OPENAI: OpenAIConfig{
			OPENAI_KEY: os.Getenv("OPENAI_KEY"),
		},
	}, nil
}
