package routes

import (
	"talkspace/features/doctor/handler"
	"talkspace/features/doctor/repository"
	"talkspace/features/doctor/service"
	"talkspace/middlewares"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DoctorRoutes(e *echo.Group, db *gorm.DB, rdb *redis.Client) {

	doctorRepository := repository.NewDoctorRepository(db, rdb)
	doctorService := service.NewDoctorService(doctorRepository)
	doctorHandler := handler.NewDoctorHandler(doctorService)

	account := e.Group("/account")
	account.POST("register", doctorHandler.Register)
	account.GET("verify-account", doctorHandler.VerifyAccount)
	account.POST("login", doctorHandler.Login)

	password := e.Group("/password")
	password.POST("forgot-password", doctorHandler.ForgotPassword)
	password.POST("verify-otp", doctorHandler.VerifyOTP)
	password.PATCH("new-password", doctorHandler.NewPassword, middlewares.JWTMiddleware())

	profile := e.Group("/profile", middlewares.JWTMiddleware())
	profile.GET("", doctorHandler.GetDoctorByID)
	profile.PUT("", doctorHandler.UpdateByID)
	profile.PATCH("/change-password", doctorHandler.UpdatePassword)

}
