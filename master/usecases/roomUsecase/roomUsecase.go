package roomUsecase

import "hotelApp/master/models"

type RoomUseCase interface {
	GetRooms() ([]*models.DetailRooms, error)
	GetRoomsByType(id int) ([]*models.DetailRooms, error)
	PostRoom(room *models.Room) error
	PutRoom(id int, room *models.Room) error
	DeleteRoom(id int) error
}