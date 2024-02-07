package routes

func ChatbotRoutes(e *echo.Group, db *gorm.DB, rdb *redis.Client) {

	chatbotRepository := repository.NewChatbotRepository(db, rdb)
	chatbotService := service.NewChatbotService(chatbotRepository)
	chatbotHandler := handler.NewChatbotHandler(chatbotService)

}
