package reservationUsecase

import "hotelApp/master/models"

type ReservationUseCase interface {
	GetReservation() ([]*models.DetailRooms, error)
	PostRoom(room *models.Room) error
}