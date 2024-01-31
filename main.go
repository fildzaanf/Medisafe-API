package main

import (
	"log"
	"os"
	"talkspace/app/configs"
	"talkspace/app/databases/firebase"
	"talkspace/app/databases/mysql"
	"talkspace/app/databases/redis"
	"talkspace/app/routes"
	"talkspace/middlewares"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {

	// load configurations
	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load configuration: %v", err)
	}

	// connect to mysql
	db := mysql.ConnectMySQL()

	// connect to redis
	redisClient := redis.ConnectRedis()
	defer redisClient.Close()

	// connect to firebase
	fcmClient := firebase.ConnectFirebase()

	e := echo.New()

	// load middlewares
	middlewares.RemoveTrailingSlash(e)
	middlewares.Logger(e)
	middlewares.RateLimiter(e)
	middlewares.Recover(e)
	middlewares.CORS(e)

	// register routes
	routes.SetupRoutes(e, db)

	// get the port from env
	godotenv.Load()
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}

	// start the server
	log.Printf("server is running on port %s...", port)
	if err := e.Start(port); err != nil {
		logrus.Fatalf("error starting server: %v", err)
	}
}
