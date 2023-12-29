package entity

import "github.com/labstack/echo"

type UserRepositoryInterface interface {
	Register(userCore User) (User, error)
	GetByID(id string) (User, error)
	UpdateByID(id string, userCore User) error
	FindByEmail(email string) (User, error)
}

type UserServiceInterface interface {
	Register(userCore User) (User, error)
	Login(email, password string) (User, string, error)
	GetByID(id string) (User, error)
	UpdateByID(id string, userCore User) error
}

type UserHandlerInterface interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetUserByID(c echo.Context) error
	UpdateByID(c echo.Context) error
}
