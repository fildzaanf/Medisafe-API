package response

type RoomRes struct {
	ID uint `json:"id"`
	CreatorId uint `json:"creator_id"`
	UserName string `json:"user_name"`
	UserPhoto string `json:"user_photo"`
}