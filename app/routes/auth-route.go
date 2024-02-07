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

func AuthRoutes(e *echo.Group, db *gorm.DB, rdb *redis.Client) {

	userRepository := repository.NewUserRepository(db, rdb)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	e.POST("register", userHandler.Register)
	e.GET("verify-account", userHandler.VerifyAccount)
	e.POST("login", userHandler.Login)


	e.POST("forgot-password", userHandler.ForgotPassword)
	e.POST("verify-otp", userHandler.VerifyOTP)
	e.PATCH("new-password", userHandler.NewPassword, middlewares.JWTMiddleware())

}
