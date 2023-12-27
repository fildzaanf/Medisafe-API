package request

type Message struct {
	ID 	  uint `json:"id"`
	Body  string `json:"body"`
	UserId uint  `json:"user_id"`
}

type CreateRoom struct {
	UserId uint `json:"user_id"`
}