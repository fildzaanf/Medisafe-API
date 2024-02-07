package routes

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, rdb *redis.Client) {

	auth := e.Group("/")
	user := e.Group("/users")

	AuthRoutes(auth, db, rdb)
	UserRoutes(user, db, rdb)

}
