package utils

const (
	SELECT_ROOMS = `SELECT rt.name_room, r.room_numb, r.price, r.available FROM room r JOIN room_type rt ON r.type_id = rt.id`
	SELECT_ROOMS_BY_TYPE = `SELECT rt.name_room, r.room_numb, r.price, r.available FROM room r JOIN room_type rt ON r.type_id = rt.id where rt.id = ?`
	INSERT_ROOM = `INSERT INTO room(room_numb, price, type_id, available) VALUES (?, ?, ?, ?)`
	UPDATE_ROOM = `UPDATE room SET room_numb = ?, price = ?, type_id = ?, available = ? WHERE id = ?`
	DELETE_ROOM = `DELETE FROM room WHERE id = ?`
	SELECT_RESERVATION = `SELECT v.identity_numb, v.name, v.address, re.check_in, re.check_out, ro.room_numb, ro.price*(day(re.check_out)-day(re.check_in)) FROM rent re JOIN visitor v ON re.id_visitor = v.id JOIN room ro ON re.id_room = ro.id`
	INSERT_VISITOR = `INSERT INTO visitor(identity_numb, name, address, hp_numb) VALUES (?, ?, ?, ?)`
	INSERT_RENT = `INSERT INTO rent(check_in, check_out, id_visitor, id_room) VALUES (?, ?, ?, ?)`
)
