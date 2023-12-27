package model

type Message struct {
	ID 	  uint `gorm:"primaryKey"`
	Content string 
	RoomId uint 
	SenderId uint
}

type Room struct {
	ID uint `gorm:"primaryKey"`
	CreatorId uint
}