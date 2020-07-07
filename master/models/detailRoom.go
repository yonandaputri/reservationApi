package models

type DetailRooms struct {
	RoomType string `json:"room_type"`
	RoomNumb string `json:"room_numb"`
	Price int `json:"price"`
	Availability string `json:"availability"`
}