package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	
	auth := e.Group("/")
	user := e.Group("/users")

	AuthRoutes(auth, db)
	UserRoutes(user, db)

}
