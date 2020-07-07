package reservationRepo

import (
	"database/sql"
	"hotelApp/master/models"
	"hotelApp/utils"
)

type ReservationRepoImpl struct{
	db *sql.DB
}

func (rv ReservationRepoImpl) GetAllReservation() ([]*models.Reservation, error) {
	data, err := rv.db.Query(utils.SELECT_RESERVATION)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var reservations []*models.Reservation

	for data.Next() {
		var reservation = new(models.Reservation)
		var err = data.Scan(&reservation.IdentityNumb, &reservation.NameVisitor, &reservation.Address, &reservation.CheckIn, &reservation.CheckOut, &reservation.RoomNum, &reservation.TotalPrice)

		if err != nil {
			return nil, err
		}

		reservations = append(reservations, reservation)
	}

	if err = data.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (rv ReservationRepoImpl) AddReservation(reservation *models.InsertReservation) error {
	data, err := rv.db.Begin()

	if err != nil {
		return err
	}

	row, err := rv.db.Exec(utils.INSERT_VISITOR, reservation.IdentityNumb, reservation.NameVisitor, reservation.Address, reservation.HpNumb)

	if err != nil {
		return err
	}

	idVisitor, err := row.LastInsertId()

	if err != nil {
		return err
	}

	row, err = rv.db.Exec(utils.INSERT_RENT, reservation.CheckIn, reservation.CheckOut, idVisitor, reservation.RoomId)

	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}

	return nil
}

