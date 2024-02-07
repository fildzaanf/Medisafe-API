package routes

import (
	"talkspace/features/user/handler"
	"talkspace/features/user/repository"
	"talkspace/features/user/service"
	"talkspace/middlewares"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(e *echo.Group, db *gorm.DB, rdb *redis.Client) {

	userRepository := repository.NewUserRepository(db, rdb)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	user := e.Group("/profile", middlewares.JWTMiddleware())
	user.GET("", userHandler.GetUserByID)
	user.PUT("", userHandler.UpdateByID)
	user.PATCH("/change-password", userHandler.UpdatePassword)

}
