package master

import (
	"database/sql"
	"github.com/gorilla/mux"
	"hotelApp/master/controllers"
	roomRepo2 "hotelApp/master/repositories/roomRepo"
	"hotelApp/master/usecases/roomUsecase"
)

func Init(r *mux.Router, db *sql.DB) {
	roomRepo := roomRepo2.InitRoomRepoImpl(db)
	roomUsecase := roomUsecase.InitRoomUseCase(roomRepo)
	controllers.RoomController(r, roomUsecase)
}
