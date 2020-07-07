package reservationUsecase

import "hotelApp/master/models"

type ReservationUseCase interface {
	GetReservation() ([]*models.Reservation, error)
	PostReservation(reservation *models.InsertReservation) error
}