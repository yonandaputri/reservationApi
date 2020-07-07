package reservationRepo

import "hotelApp/master/models"

type ReservationRepository interface {
	GetAllReservation() ([]*models.Reservation, error)
	AddReservation(reservation *models.UpdateReservation) error
}