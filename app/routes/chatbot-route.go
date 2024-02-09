package routes

func ChatbotRoutes(e *echo.Group, db *gorm.DB) {

	chatbotRepository := repository.NewChatbotRepository(db)
	chatbotService := service.NewChatbotService(chatbotRepository)
	chatbotHandler := handler.NewChatbotHandler(chatbotService)

	e.POST("", chatbotHandler.CreateMessageChatBot)

}
