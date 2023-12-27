package handler

import (
	"net/http"
	"strconv"
	"talkspace/app/databases/mysql"
	"talkspace/features/consultation/dto/response"
	userService "talkspace/features/auth/service"
	"talkspace/features/consultation/dto/request"
	"talkspace/features/consultation/model"
	doctorService "talkspace/features/doctor/service"
	"talkspace/utils/helper/websocket"

	"github.com/golang-jwt/jwt"
	ws "github.com/gorilla/websocket"

	"github.com/labstack/echo/v4"
)

type ConsultationHandlerImpl struct {
	UserService userService.UserService
	DoctorService doctorService.DoctorService
	hub *websocket.Hub
}

type ConsultationHandler interface {
	JoinRoom(c echo.Context) error
	CreateRoom(c echo.Context) error
	GetRoom(c echo.Context) (uint, error)
}

func NewConsultationHandler(userService userService.UserService, doctorService doctorService.DoctorService, hub *websocket.Hub) *ConsultationHandlerImpl {
	return &ConsultationHandlerImpl{
		UserService: userService,
		DoctorService: doctorService,
		hub: hub,
	}
}

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (handler *ConsultationHandlerImpl) CreateRoom(c echo.Context) error {
	var reqRoom request.CreateRoom
	if err := c.Bind(&reqRoom); err != nil {
		return err
	}

	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)

	user, err := handler.UserService.GetUserByID(ID)
	if err != nil {
		return err
	}

	room := model.Room{
		CreatorId: user.ID,
	}

	mysql.DB.Create(&room)
	handler.hub.Rooms[room.ID] = &websocket.Room{
		ID:       room.ID,
		CreatorId: room.CreatorId,
		Client: make(map[uint]*websocket.Client),
	}

	return c.JSON(http.StatusOK, "success")
}

func (handler *ConsultationHandlerImpl) JoinRoom(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)

	var clientId uint
	if user, err := handler.UserService.FindUserByID(ID); err != nil {
		doctor, err := handler.DoctorService.FindDoctorByID(ID)
		if err != nil {
			return err
		}
		clientId = doctor.ID
	} else {
		clientId = user.ID
	}
	
	roomId, _ := strconv.Atoi(c.QueryParam("room_id"))

	cl := &websocket.Client{
		Conn:   conn,
		Message: make(chan *websocket.Message, 10),
		ID: clientId,
		RoomId: uint(roomId),
		SenderId: clientId,
	}

	handler.hub.Register <- cl

	go cl.WriteMessage(uint(roomId))
	cl.ReadMessage(handler.hub)

	return nil
}

func (handler *ConsultationHandlerImpl) GetRoom(c echo.Context)  error {
	rooms := make([]response.RoomRes, 0)
	var user model.User

	for _, room := range handler.hub.Rooms {
		mysql.DB.Find(&user, " id = ? ", room.CreatorId)
		rooms = append(rooms, response.RoomRes{
			ID: room.ID,
			CreatorId: room.CreatorId,
			UserName: user.Name,
			UserPhoto: user.Photo,
		})
	}

	return c.JSON(http.StatusOK, rooms)
}
