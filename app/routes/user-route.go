package routes

import (
	"talkspace/features/user/handler"
	"talkspace/middlewares"
	"talkspace/features/user/repository"
	"talkspace/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(e *echo.Group, db *gorm.DB) {

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	user := e.Group("/profile", middlewares.JWTMiddleware())
	user.GET("", userHandler.GetUserByID)
	user.PUT("", userHandler.UpdateByID)
	user.PATCH("/reset-password", userHandler.UpdatePassword)

}
