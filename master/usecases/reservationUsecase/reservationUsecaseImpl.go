package reservationUsecase

import (
	"hotelApp/master/models"
	"hotelApp/master/repositories/reservationRepo"
)

type ReservationUsecaseImpl struct{
	reservRepo reservationRepo.ReservationRepository
}

func (rv ReservationUsecaseImpl) GetReservation() ([]*models.Reservation, error) {
	reservation, err := rv.reservRepo.GetAllReservation()
	if err != nil {
		return nil, err
	}
	return reservation, nil
}

func (rv ReservationUsecaseImpl) PostReservation(reservation *models.InsertReservation) error {
	err := rv.reservRepo.AddReservation(reservation)
	if err != nil {
		return err
	}
	return nil
}
