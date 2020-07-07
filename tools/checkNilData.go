package tools

import "hotelApp/master/models"

func CheckNilData(room *models.Room) bool {
	if room.RoomNumb == "" || room.Availability == "" {
		return false
	} else {
		return true
	}
}
