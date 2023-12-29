package handler

import (
	"net/http"
	"talkspace/features/user/dto/request"
	"talkspace/features/user/dto/response"
	"talkspace/features/user/entity"
	"talkspace/utils/constant"
	"talkspace/utils/responses"

	"github.com/labstack/echo"
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

	return c.JSON(http.StatusCreated, responses.SuccessResponse(constant.SUCCESS_CREATED, response))
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
