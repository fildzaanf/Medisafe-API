package entity

import "github.com/labstack/echo/v4"

type DoctorRepositoryInterface interface {
	Register(doctorCore Doctor) (Doctor, error)
	Login(email, password string) (Doctor, error)
	GetByID(id string) (Doctor, error)
	UpdateByID(id string, doctorCore Doctor) error
	UpdatePassword(id string, doctorCore Doctor) error
	NewPassword(email string, doctorCore Doctor) (Doctor, error)
	FindByEmail(email string) (Doctor, error)
	GetByVerificationToken(token string) (Doctor, error)
	UpdateIsVerified(id string, isVerified bool) error
	SendOTP(email string, otp string, expired int64) (Doctor, error)
	VerifyOTP(email, otp string) (Doctor, error)
	ResetOTP(otp string) (Doctor, error)
}

type DoctorServiceInterface interface {
	Register(doctorCore Doctor) (Doctor, error)
	Login(email, password string) (Doctor, string, error)
	GetByID(id string) (Doctor, error)
	UpdateByID(id string, doctorCore Doctor) error
	UpdatePassword(id string, doctorCore Doctor) error
	NewPassword(email string, doctorCore Doctor) error
	VerifyDoctor(token string) (bool, error)
	UpdateIsVerified(id string, isVerified bool) error
	SendOTP(email string) error
	VerifyOTP(email, otp string) (string, error)
}

type DoctorHandlerInterface interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetDoctorByID(c echo.Context) error
	UpdateByID(c echo.Context) error
	UpdatePassword(c echo.Context) error
	ForgotPassword(c echo.Context) error
	NewPassword(c echo.Context) error
	VerifyAccount(c echo.Context) error
	VerifyOTP(c echo.Context) error
}
