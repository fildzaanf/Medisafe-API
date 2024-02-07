package routes

func ConsultationRoutes(e *echo.Group, db *gorm.DB, rdb *redis.Client) {

	consultationRepository := repository.NewConsultationRepository(db, rdb)
	consultationService := service.NewConsultationService(consultationRepository)
	consultationHandler := handler.NewConsultationHandler(consultationService)

}
