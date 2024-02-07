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

	account := e.Group("/account")
	account.POST("register", userHandler.Register)
	account.GET("verify-account", userHandler.VerifyAccount)
	account.POST("login", userHandler.Login)

	password := e.Group("/password")
	password.POST("forgot-password", userHandler.ForgotPassword)
	password.POST("verify-otp", userHandler.VerifyOTP)
	password.PATCH("new-password", userHandler.NewPassword, middlewares.JWTMiddleware())

	profile := e.Group("/profile", middlewares.JWTMiddleware())
	profile.GET("", userHandler.GetUserByID)
	profile.PUT("", userHandler.UpdateByID)
	profile.PATCH("/change-password", userHandler.UpdatePassword)

}



