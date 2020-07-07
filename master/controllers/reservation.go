package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"hotelApp/master/models"
	"hotelApp/master/usecases/reservationUsecase"
	"net/http"
)

type ReservatioHandler struct {
	ReservationUseCase reservationUsecase.ReservationUseCase
}

func ReservationController(r *mux.Router, service reservationUsecase.ReservationUseCase) {
	reservationHandler := ReservatioHandler{service}
	r.HandleFunc("/reservations", reservationHandler.ListReservation).Methods(http.MethodGet)
	r.HandleFunc("/reservation", reservationHandler.PostReservation).Methods(http.MethodPost)
}

func (rv ReservatioHandler) ListReservation(w http.ResponseWriter, r *http.Request) {
	rooms, err := rv.ReservationUseCase.GetReservation()

	var response models.MessageResponse
	response.Message = "Success Get Data"
	response.Data = rooms

	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfRooms, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfRooms))
}

func (rv ReservatioHandler) PostReservation(w http.ResponseWriter, r *http.Request) {
	var reservation []*models.InsertReservation
	err := json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	for _, value := range reservation {
			err := rv.ReservationUseCase.PostReservation(value)
			if err != nil {
				w.Write([]byte("Cannot add data"))
			}
	}
	w.Write([]byte("Succes add data"))
}