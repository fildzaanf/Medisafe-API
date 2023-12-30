package entity

import "github.com/labstack/echo"

type UserRepositoryInterface interface {
	Register(userCore User) (User, error)
	GetByID(id string) (User, error)
	UpdateByID(id string, userCore User) error
	FindByEmail(email string) (User, error)
	GetByVerificationToken(token string) (User, error)
	UpdateIsVerified(id string, isVerified bool) error
	SendOTP(email string, otp string, expired int64) (User, error)
	VerifyOTP(email, otp string) (User, error)
	ResetOTP(otp string) (User, error)
}

type UserServiceInterface interface {
	Register(userCore User) (User, error)
	Login(email, password string) (User, string, error)
	GetByID(id string) (User, error)
	UpdateByID(id string, userCore User) error
	VerifyUser(token string) (bool, error)
	UpdateIsVerified(id string, isVerified bool) error
	SendOTP(email string) error
	VerifyOTP(email, otp string) (string, error)
}

type UserHandlerInterface interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetUserByID(c echo.Context) error
	UpdateByID(c echo.Context) error
	VerifyAccount(c echo.Context) error
	VerifyOTP(c echo.Context) error
}
