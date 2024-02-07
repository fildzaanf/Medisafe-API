package main

import (
	"log"
	//"talkspace/app/databases/mysql"
	"talkspace/app/configs"
	"talkspace/app/databases/postgresql"
	"talkspace/app/databases/redis"
	"talkspace/app/routes"
	"talkspace/middlewares"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()
	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load configuration: %v", err)
	}

	// db := mysql.ConnectMySQL()
	db := postgresql.ConnectPostgreSQL()
	rdb := redis.ConnectRedis()
	defer rdb.Close()
	// fcm := firebase.ConnectFirebase()

	e := echo.New()

	middlewares.RemoveTrailingSlash(e)
	middlewares.Logger(e)
	middlewares.RateLimiter(e)
	middlewares.Recover(e)
	middlewares.CORS(e)

	routes.SetupRoutes(e, db, rdb)

	host := config.SERVER.SERVER_HOST
	port := config.SERVER.SERVER_PORT
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "8000"
	}
	address := host + ":" + port

	log.Printf("server is running on address %s...", address)
	if err := e.Start(address); err != nil {
		logrus.Fatalf("error starting server: %v", err)
	}
}
