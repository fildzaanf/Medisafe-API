package websocket

import (
	"log"
	"talkspace/app/databases/mysql"
	"talkspace/features/consultation/model"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Message chan *Message
	ID 	uint
	RoomId  uint
	SenderId  uint
	Username string
}

type Message struct {
	Content  string `json:"content"`
	RoomId   uint   `json:"room_id"`
	SenderId uint   `json:"user_id"`
}

func (c *Client) ReadMessage(hub *Hub){
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content : string(m),
			RoomId: c.RoomId,
			SenderId: c.SenderId,
		}
		
		message := model.Message{
			Content: msg.Content,
			RoomId: msg.RoomId,
			SenderId: msg.SenderId,
		}

		mysql.DB.Create(&message)

		hub.Broadcast <- msg
	}
}

func (c *Client) WriteMessage(RoomId uint){
	defer func(){
		c.Conn.Close()
	}()

	var message []model.Message
	mysql.DB.Where("room_id = ?", RoomId).Find(&message)
	c.Conn.WriteJSON(message)

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}