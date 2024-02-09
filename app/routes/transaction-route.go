package routes

import (
	"talkspace/features/user/handler"
	"talkspace/features/user/repository"
	"talkspace/features/user/service"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRoutes(e *echo.Group, db *gorm.DB, rdb *redis.Client) {

	transactionRepository := repository.NewTransactionRepository(db, rdb)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	transaction := e.Group("/consultation", middlewares.JWTMiddleware())
	transaction.POST("", transactionHandler.CreateTransaction)
	transaction.GET("", transactionHandler.GetAllTransactions)
	transaction.GET("/:transaction_id", transactionHandler.GetTransactionByID)

}

