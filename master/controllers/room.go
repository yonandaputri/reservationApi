package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"hotelApp/master/models"
	"hotelApp/master/usecases/roomUsecase"
	"hotelApp/tools"
	"net/http"
	"strconv"
)

type RoomHandler struct {
	RoomUseCase roomUsecase.RoomUseCase
}

func RoomController(r *mux.Router, service roomUsecase.RoomUseCase) {
	roomHandler := RoomHandler{service}
	r.HandleFunc("/rooms", roomHandler.ListRooms).Methods(http.MethodGet)
	r.HandleFunc("/rooms/{id}", roomHandler.GetRoomsByTypeId).Methods(http.MethodGet)
	r.HandleFunc("/room", roomHandler.PostRoom).Methods(http.MethodPost)
	r.HandleFunc("/room/{id}", roomHandler.PutRoom).Methods(http.MethodPut)
	r.HandleFunc("/room/{id}", roomHandler.DeleteRoom).Methods(http.MethodDelete)
}

func (rh RoomHandler) ListRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := rh.RoomUseCase.GetRooms()

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

func (rh RoomHandler) GetRoomsByTypeId(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := tools.GetPathVar(key, r)
	idTypeRoom, _ := strconv.Atoi(id)
	rooms, err := rh.RoomUseCase.GetRoomsByType(idTypeRoom)

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

func (rh RoomHandler) PostRoom(w http.ResponseWriter, r *http.Request) {
	var room []*models.Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	for _, value := range room {
		check := tools.CheckNilData(value)
		if check == false {
			w.Write([]byte("Cannot add data"))
		} else {
			err := rh.RoomUseCase.PostRoom(value)
			if err != nil {
				w.Write([]byte("Cannot add data"))
			}
			w.Write([]byte("Succes add data"))
		}
	}
}

func (rh RoomHandler) PutRoom(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := tools.GetPathVar(key, r)
	idRoom, _ := strconv.Atoi(id)
	var room *models.Room
	check := tools.CheckNilData(room)
	if check == false {
		w.Write([]byte("Cannot Update Data"))
	} else {
		err := json.NewDecoder(r.Body).Decode(&room)
		if err != nil {
			w.Write([]byte("Oops something when wrong"))
		}
		err = rh.RoomUseCase.PutRoom(idRoom, room)
		if err != nil {
			w.Write([]byte("Cannot update data"))
		}
		w.Write([]byte("Succes update data"))
	}
}

func (rh RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := tools.GetPathVar(key, r)
	idRoom, _ := strconv.Atoi(id)
	err := rh.RoomUseCase.DeleteRoom(idRoom)
	if err != nil {
		w.Write([]byte("Cannot delete data"))
	}
	w.Write([]byte("Succes delete data"))
}
