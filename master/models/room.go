package models

type Room struct {
	RoomNumb string `json:"room_numb"`
	Price int `json:"price"`
	TypeId int `json:"type_id"`
	Availability string `json:"availability"`
}
