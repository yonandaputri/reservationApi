package roomRepo

import "hotelApp/master/models"

type RoomRepository interface {
	GetAllRoom() ([]*models.DetailRooms, error)
	GetRoomByType(id int) ([]*models.DetailRooms, error)
	AddRoom(room *models.Room) error
	UpdateRoom(id int, room *models.Room) error
	DeleteRoom(id int) error
}