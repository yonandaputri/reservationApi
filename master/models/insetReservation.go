package models

type InsertReservation struct {
	IdentityNumb string `json:"identity_numb"`
	NameVisitor string `json:"name_visitor"`
	Address string `json:"address"`
	HpNumb string `json:"hp_numb"`
	CheckIn string `json:"check_in"`
	CheckOut string `json:"check_out"`
	RoomId int `json:"room_id"`
}
