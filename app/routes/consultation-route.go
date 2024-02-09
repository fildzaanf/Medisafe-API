package routes

import (
	"talkspace/middlewares"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ConsultationRoutes(e *echo.Group, db *gorm.DB, rdb *redis.Client) {

	consultationRepository := repository.NewConsultationRepository(db, rdb)
	consultationService := service.NewConsultationService(consultationRepository)
	consultationHandler := handler.NewConsultationHandler(consultationService)

	doctor := e.Group("doctor")
	doctor.GET("", consultationHandler.GetAllDoctors)
	doctor.GET("/:doctor_id", consultationHandler.GetDoctorByID)

	roomchat := e.Group("/roomchat", middlewares.JWTMiddleware())
	roomchat.POST("/:transaction_id", consultationHandler.CreateRoomchat)
	roomchat.GET("", consultationHandler.GetAllRoomchats)
	roomchat.GET("/:roomchat_id", consultationHandler.GetRoomchatByID)

	message := e.Group("/message", middlewares.JWTMiddleware())
	message.POST("/:roomchat_id", consultationHandler.CreateMessage)

}


