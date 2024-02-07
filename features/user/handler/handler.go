package handler

import (
	"net/http"
	"strings"
	"talkspace/features/user/dto/request"
	"talkspace/features/user/dto/response"
	"talkspace/features/user/entity"
	"talkspace/middlewares"
	"talkspace/utils/constant"
	"talkspace/utils/helper/generator"
	"talkspace/utils/responses"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService entity.UserServiceInterface
}

func NewUserHandler(us entity.UserServiceInterface) *userHandler {
	return &userHandler{
		userService: us,
	}
}

func (uh *userHandler) Register(c echo.Context) error {
	registerRequest := request.UserRegisterRequest{}

	errBind := c.Bind(&registerRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	request := request.UserRegisterRequestToUserCore(registerRequest)

	_, errCreate := uh.userService.Register(request)
	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errCreate.Error()))
	}

	response := response.UserCoreToUserRegisterResponse(request)

	return c.JSON(http.StatusCreated, responses.SuccessResponse(constant.SUCCESS_REGISTER, response))
}

func (uh *userHandler) Login(c echo.Context) error {
	loginRequest := request.UserLoginRequest{}

	errBind := c.Bind(&loginRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	request, token, errLogin := uh.userService.Login(loginRequest.Email, loginRequest.Password)
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errLogin.Error()))
	}

	response := response.UserCoreToUserLoginResponse(request, token)

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_LOGIN, response))
}

func (uh *userHandler) GetUserByID(c echo.Context) error {
	userID, role, errExtract := middlewares.ExtractToken(c)

	if role != constant.USER {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
	}

	result, errGetID := uh.userService.GetByID(userID)
	if errGetID != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
	}

	response := response.UserCoreToUserProfileResponse(result)

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_PROFILE_RETRIEVED, response))
}

func (uh *userHandler) UpdateByID(c echo.Context) error {
	requestUpdateProfile := request.UserUpdateProfileRequest{}

	errBind := c.Bind(&requestUpdateProfile)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	userID, role, errExtract := middlewares.ExtractToken(c)
	
	if role != constant.USER {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
	}

	request := request.UserUpdateProfileRequestToUserCore(requestUpdateProfile)

	errUpdate := uh.userService.UpdateByID(userID, request)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_PROFILE_UPDATED, nil))
}

func (uh *userHandler) UpdatePassword(c echo.Context) error {
	requestUpdatePasword := request.UserUpdatePasswordRequest{}

	errBind := c.Bind(&requestUpdatePasword)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	userID, role, errExtractToken := middlewares.ExtractToken(c)

	if role != constant.USER {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	if errExtractToken != nil {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtractToken.Error()))
	}

	request := request.UserUpdatePasswordRequestToUserCore(requestUpdatePasword)

	errUpdate := uh.userService.UpdatePassword(userID, request)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_PASSWORD_UPDATED, nil))
}

func (uh *userHandler) VerifyAccount(c echo.Context) error {
	token := c.QueryParam("token")

	userVerified, errVerified := uh.userService.VerifyUser(token)
	if errVerified != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errVerified.Error()))
	}

	if userVerified {
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

func (uh *userHandler) ForgotPassword(c echo.Context) error {
	requestOTP := request.UserSendOTPRequest{}

	errBind := c.Bind(&requestOTP)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	request := request.UserSendOTPRequestToUserCore(requestOTP)

	errSendOTP := uh.userService.SendOTP(request.Email)
	if errSendOTP != nil {
		if strings.Contains(errSendOTP.Error(), constant.ERROR_EMAIL_NOTFOUND) {
			return c.JSON(http.StatusNotFound, responses.ErrorResponse(errSendOTP.Error()))
		}
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errSendOTP.Error()))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_OTP_SENT, nil))
}

func (uh *userHandler) VerifyOTP(c echo.Context) error {
	requestVerifyOTP := request.UserVerifyOTPRequest{}

	errBind := c.Bind(&requestVerifyOTP)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	request := request.UserVerifyOTPRequestToUserCore(requestVerifyOTP)

	token, errVerify := uh.userService.VerifyOTP(request.Email, request.OTP)
	if errVerify != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse(constant.ERROR_OTP_VERIFY+errVerify.Error()))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_OTP_VERIFIED, token))
}

func (uh *userHandler) NewPassword(c echo.Context) error {
	requestNewPassword := request.UserNewPasswordRequest{}

	errBind := c.Bind(&requestNewPassword)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
	}

	email, errExtract := middlewares.ExtractVerifyToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
	}

	request := request.UserNewPasswordRequestToUserCore(requestNewPassword)

	err := uh.userService.NewPassword(email, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, responses.SuccessResponse(constant.SUCCESS_PASSWORD_UPDATED, nil))
}
