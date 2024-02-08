package handler

import (
	"net/http"
	"strings"
	"talkspace/features/doctor/dto/request"
	"talkspace/features/doctor/dto/response"
	"talkspace/features/doctor/entity"
	"talkspace/middlewares"
	"talkspace/utils/constant"
	"talkspace/utils/helper/generator"
	"talkspace/utils/responses"

	"github.com/labstack/echo/v4"
)

type doctorHandler struct {
	doctorService entity.DoctorServiceInterface
}

func NewDoctorHandler(ds entity.DoctorServiceInterface) *doctorHandler {
	return &doctorHandler{
		doctorService: ds,
	}
}

func (dh *doctorHandler) Register(c echo.Context) error {
	registerRequest := request.DoctorRegisterRequest{}

	errBind := c.Bind(&registerRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	request := request.DoctorRegisterRequestToDoctorCore(registerRequest)

	_, errCreate := dh.doctorService.Register(request)
	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errCreate.Error()))
	}

	response := response.DoctorCoreToDoctorRegisterResponse(request)

	return c.JSON(http.StatusCreated, responses.SuccessResponse(constant.SUCCESS_REGISTER, response))
}

func (dh *doctorHandler) Login(c echo.Context) error {
	loginRequest := request.DoctorLoginRequest{}

	errBind := c.Bind(&loginRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	request, token, errLogin := dh.doctorService.Login(loginRequest.Email, loginRequest.Password)
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errLogin.Error()))
	}

	response := response.DoctorCoreToDoctorLoginResponse(request, token)

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_LOGIN, response))
}

func (dh *doctorHandler) GetDoctorByID(c echo.Context) error {
	doctorID, role, errExtract := middlewares.ExtractToken(c)

	if role != constant.DOCTOR {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
	}

	result, errGetID := dh.doctorService.GetByID(doctorID)
	if errGetID != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
	}

	response := response.DoctorCoreToDoctorProfileResponse(result)

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_PROFILE_RETRIEVED, response))
}

func (dh *doctorHandler) UpdateByID(c echo.Context) error {
	requestUpdateProfile := request.DoctorProfileRequest{}

	errBind := c.Bind(&requestUpdateProfile)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	doctorID, role, errExtract := middlewares.ExtractToken(c)

	if role != constant.DOCTOR {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
	}

	request := request.DoctorProfileRequestToDoctorCore(requestUpdateProfile)

	errUpdate := dh.doctorService.UpdateByID(doctorID, request)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_PROFILE_UPDATED, nil))
}

func (dh *doctorHandler) UpdatePassword(c echo.Context) error {
	requestUpdatePassword := request.DoctorUpdatePasswordRequest{}

	errBind := c.Bind(&requestUpdatePassword)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	doctorID, role, errExtractToken := middlewares.ExtractToken(c)

	if role != constant.DOCTOR {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	if errExtractToken != nil {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtractToken.Error()))
	}

	request := request.DoctorUpdatePasswordRequestToDoctorCore(requestUpdatePassword)

	errUpdate := dh.doctorService.UpdatePassword(doctorID, request)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_PASSWORD_UPDATED, nil))
}

func (dh *doctorHandler) VerifyAccount(c echo.Context) error {
	token := c.QueryParam("token")

	doctorVerified, errVerified := dh.doctorService.VerifyDoctor(token)
	if errVerified != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errVerified.Error()))
	}

	if doctorVerified {
		email, errGenerateTemplate := generator.GenerateEmailTemplate("verification-account.html", nil)
		if errGenerateTemplate != nil {
			return c.JSON(http.StatusInternalServerError, responses.ErrorResponse(constant.ERROR_TEMPLATE_READER))
		}
		return c.HTML(http.StatusOK, email)
	}

	email, errGenerateTemplate := generator.GenerateEmailTemplate("verification-account-success.html", nil)
	if errGenerateTemplate != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse(constant.ERROR_TEMPLATE_READER))
	}
	return c.HTML(http.StatusOK, email)
}

func (dh *doctorHandler) ForgotPassword(c echo.Context) error {
	requestOTP := request.DoctorSendOTPRequest{}

	errBind := c.Bind(&requestOTP)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	request := request.DoctorSendOTPRequestToDoctorCore(requestOTP)

	errSendOTP := dh.doctorService.SendOTP(request.Email)
	if errSendOTP != nil {
		if strings.Contains(errSendOTP.Error(), constant.ERROR_EMAIL_NOTFOUND) {
			return c.JSON(http.StatusNotFound, responses.ErrorResponse(errSendOTP.Error()))
		}
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errSendOTP.Error()))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_OTP_SENT, nil))
}

func (dh *doctorHandler) VerifyOTP(c echo.Context) error {
	requestVerifyOTP := request.DoctorVerifyOTPRequest{}

	errBind := c.Bind(&requestVerifyOTP)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	request := request.DoctorVerifyOTPRequestToDoctorCore(requestVerifyOTP)

	token, errVerify := dh.doctorService.VerifyOTP(request.Email, request.OTP)
	if errVerify != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse(constant.ERROR_OTP_VERIFY+errVerify.Error()))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_OTP_VERIFIED, token))
}

func (dh *doctorHandler) NewPassword(c echo.Context) error {
	requestNewPassword := request.DoctorNewPasswordRequest{}

	errBind := c.Bind(&requestNewPassword)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	email, errExtract := middlewares.ExtractVerifyToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
	}

	request := request.DoctorNewPasswordRequestToDoctorCore(requestNewPassword)

	err := dh.doctorService.NewPassword(email, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_PASSWORD_UPDATED, nil))
}
