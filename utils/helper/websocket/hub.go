package websocket

type Room struct {
	ID       uint             `json:"id"`
	CreatorId uint			 `json:"creator_id"`
	Client   map[uint]*Client `json:"client"`
}

type Hub struct {
	Rooms      map[uint]*Room `json:"rooms"`
	Register    chan *Client   `json:"register"`
	Unregister chan *Client   `json:"unregister"`	
	Broadcast  chan *Message  `json:"broadcast"`
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[uint]*Room),
		Register:    make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			room, ok := h.Rooms[client.RoomId]
			if !ok {
				room = &Room{
					ID:     client.ID,
					CreatorId:   client.SenderId,
					Client: make(map[uint]*Client),
				}
				h.Rooms[client.RoomId] = room
			}

			room.Client[client.ID] = client
		case client := <-h.Unregister:
			room, ok := h.Rooms[client.RoomId]
			if ok {
				if _, ok := room.Client[client.ID]; ok {
					delete(room.Client, client.ID)
					close(client.Message)
				}
			}
		case message := <-h.Broadcast:
			room, ok := h.Rooms[message.RoomId]
			if ok {
				for _, client := range room.Client {
					select {
					case client.Message <- message:
					default:
						close(client.Message)
						delete(room.Client, client.ID)
					}
				}
			}
		}
	}
}
