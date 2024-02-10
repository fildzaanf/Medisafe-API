package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	POSTGRESQL   PostgreSQLConfig
	MYSQL        MySQLConfig
	REDIS        RedisConfig
	CLOUDSTORAGE CloudStorageConfig
	FIREBASE     FirebaseConfig
	MIDTRANS     MidtransConfig
	SMTP         SMTPConfig
	OPENAI       OpenAIConfig
	JWT          JWTConfig
	SERVER       ServerConfig
}

type MySQLConfig struct {
	MYSQL_USER string
	MYSQL_PASS string
	MYSQL_HOST string
	MYSQL_PORT string
	MYSQL_NAME string
}

type PostgreSQLConfig struct {
	POSTGRESQL_USER string
	POSTGRESQL_PASS string
	POSTGRESQL_HOST string
	POSTGRESQL_PORT string
	POSTGRESQL_NAME string
}

type RedisConfig struct {
	REDIS_HOST string
	REDIS_PORT string
	REDIS_DB   string
	REDIS_PASS string
}

type CloudStorageConfig struct {
	GOOGLE_CLOUD_STORAGE_SERVICE_ACCOUNT string
	GOOGLE_CLOUD_PROJECT_ID              string
	GOOGLE_CLOUD_STORAGE_BUCKET          string
}

type FirebaseConfig struct {
	FIREBASE_API_KEY string
}

type MidtransConfig struct {
	MIDTRANS_SERVER_KEY string
	MIDTRANS_CLIENT_KEY string
}

type OpenAIConfig struct {
	OPENAI_API_KEY string
}

type SMTPConfig struct {
	SMTP_USER string
	SMTP_PASS string
	SMTP_PORT string
	SMTP_HOST string
}

type ServerConfig struct {
	SERVER_HOST string
	SERVER_PORT string
}

type JWTConfig struct {
	JWT_SECRET string
}

func LoadConfig() (*Configuration, error) {

	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("failed to load environment variables from .env file: %w", err)
		}
	} else {
		fmt.Println("warning: .env file not found. make sure environment variables are set")
	}

	return &Configuration{
		MYSQL: MySQLConfig{
			MYSQL_USER: os.Getenv("MYSQL_USER"),
			MYSQL_PASS: os.Getenv("MYSQL_PASS"),
			MYSQL_HOST: os.Getenv("MYSQL_HOST"),
			MYSQL_PORT: os.Getenv("MYSQL_PORT"),
			MYSQL_NAME: os.Getenv("MYSQL_NAME"),
		},
		POSTGRESQL: PostgreSQLConfig{
			POSTGRESQL_USER: os.Getenv("POSTGRESQL_USER"),
			POSTGRESQL_PASS: os.Getenv("POSTGRESQL_PASS"),
			POSTGRESQL_HOST: os.Getenv("POSTGRESQL_HOST"),
			POSTGRESQL_PORT: os.Getenv("POSTGRESQL_PORT"),
			POSTGRESQL_NAME: os.Getenv("POSTGRESQL_DB"),
		},
		REDIS: RedisConfig{
			REDIS_HOST: os.Getenv("REDIS_HOST"),
			REDIS_PORT: os.Getenv("REDIS_PORT"),
			REDIS_DB:   os.Getenv("REDIS_DB"),
			REDIS_PASS: os.Getenv("REDIS_PASS"),
		},
		CLOUDSTORAGE: CloudStorageConfig{
			GOOGLE_CLOUD_STORAGE_SERVICE_ACCOUNT: os.Getenv("GOOGLE_CLOUD_STORAGE_SERVICE_ACCOUNT"),
			GOOGLE_CLOUD_PROJECT_ID:              os.Getenv("GOOGLE_CLOUD_PROJECT_ID"),
			GOOGLE_CLOUD_STORAGE_BUCKET:          os.Getenv("GOOGLE_CLOUD_STORAGE_BUCKET"),
		},
		FIREBASE: FirebaseConfig{
			FIREBASE_API_KEY: os.Getenv("FIREBASE_API_KEY"),
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
			OPENAI_API_KEY: os.Getenv("OPENAI_API_KEY"),
		},
		SERVER: ServerConfig{
			SERVER_HOST: os.Getenv("SERVER_HOST"),
			SERVER_PORT: os.Getenv("SERVER_PORT"),
		},
		JWT: JWTConfig{
			JWT_SECRET: os.Getenv("JWT_SECRET"),
		},
	}, nil
}
