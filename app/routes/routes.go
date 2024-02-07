package routes

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, rdb *redis.Client) {

	user := e.Group("/users")
	doctor := e.Group("/doctors")
	transaction := e.Group("/transactions")
	consultation := e.Group("/consultations")
	chatbot := e.Group("/chatbots")

	UserRoutes(user, db, rdb)
	DoctorRoutes(doctor, db, rdb)
	TransactionRoutes(transaction, db, rdb)
	ConsultationRoutes(consultation, db, rdb)
	ChatbotRoutes(chatbot, db, rdb)

}
