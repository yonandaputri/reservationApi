package roomUsecase

import (
	"hotelApp/master/models"
	"hotelApp/master/repositories/roomRepo"
)

type RoomUsecaseImpl struct{
	roomRepo roomRepo.RoomRepository
}

func (r RoomUsecaseImpl) GetRooms() ([]*models.DetailRooms, error) {
	rooms, err := r.roomRepo.GetAllRoom()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r RoomUsecaseImpl) GetRoomsByType(id int) ([]*models.DetailRooms, error) {
	rooms, err := r.roomRepo.GetRoomByType(id)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r RoomUsecaseImpl) PostRoom(room *models.Room) error {
	err := r.roomRepo.AddRoom(room)
	if err != nil {
		return err
	}
	return nil
}

func (r RoomUsecaseImpl) PutRoom(id int, room *models.Room) error {
	err := r.roomRepo.UpdateRoom(id, room)
	if err != nil {
		return err
	}
	return nil
}

func (r RoomUsecaseImpl) DeleteRoom(id int) error {
	err := r.roomRepo.DeleteRoom(id)
	if err != nil {
		return err
	}
	return nil
}

func InitRoomUseCase(roomRepo roomRepo.RoomRepository) RoomUseCase {
	return &RoomUsecaseImpl{roomRepo}
}