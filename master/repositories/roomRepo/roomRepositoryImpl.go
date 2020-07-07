package roomRepo

import (
	"database/sql"
	"hotelApp/master/models"
	"hotelApp/utils"
)

type RoomRepoImpl struct{
	db *sql.DB
}

func (r RoomRepoImpl) GetAllRoom() ([]*models.DetailRooms, error) {
	data, err := r.db.Query(utils.SELECT_ROOMS)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var rooms []*models.DetailRooms

	for data.Next() {
		var room = new(models.DetailRooms)
		var err = data.Scan(&room.RoomType, &room.RoomNumb, &room.Price, &room.Availability)

		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	if err = data.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r RoomRepoImpl) GetRoomByType(id int) ([]*models.DetailRooms, error) {
	data, err := r.db.Query(utils.SELECT_ROOMS_BY_TYPE)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var rooms []*models.DetailRooms

	for data.Next() {
		var room = new(models.DetailRooms)
		var err = data.Scan(&room.RoomType, &room.RoomNumb, &room.Price, &room.Availability, id)

		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	if err = data.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r RoomRepoImpl) AddRoom(room *models.Room) error {
	data, err := r.db.Begin()

	if err != nil {
		return err
	}

	row, err := r.db.Prepare(utils.INSERT_ROOM)

	if err != nil {
		return err
	}

	_, err = row.Exec(room.RoomNumb, room.Price, room.TypeId, room.Availability)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	return nil
}

func (r RoomRepoImpl) UpdateRoom(id int, room *models.Room) error {
	data, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, _ = r.db.Exec("SET FOREIGN_KEY_CHECKS=0;")

	row, err := r.db.Prepare(utils.UPDATE_ROOM)
	if err != nil {
		return err
	}

	_, err = row.Exec(room.RoomNumb, room.Price, room.TypeId, room.Availability, id)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	_, _ = r.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return nil
}

func (r RoomRepoImpl) DeleteRoom(id int) error {
	data, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, _ = r.db.Exec("SET FOREIGN_KEY_CHECKS=0;")

	row, err := r.db.Prepare(utils.DELETE_ROOM)
	if err != nil {
		return err
	}

	_, err = row.Exec(id)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	_, _ = r.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return nil
}

func InitRoomRepoImpl(db *sql.DB) RoomRepository {
	return &RoomRepoImpl{db}
}