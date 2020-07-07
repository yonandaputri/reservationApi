package models

type Reservation struct {
	IdentityNumb string `json:"identity_numb"`
	NameVisitor string `json:"name_visitor"`
	Address string `json:"address"`
	HpNumb string `json:"hp_numb"`
	CheckIn string `json:"check_in"`
	CheckOut string `json:"check_out"`
	RoomNum int `json:"room_num"`
	TotalPrice int `json:"total_price"`
}
