create database db_hotel;

create table room(
	id int not null auto_increment,
    type_id int,
    room_numb varchar(4),
    price int,
	availability_id int,
    primary key(id),
    foreign key(type_id) references room_type(id),
    foreign key(availability_id) references availability(id)
);

SELECT rt.name_room, r.room_numb, r.price, a.name FROM room r JOIN room_type rt ON r.type_id = rt.id JOIN availability a ON a.id = r.availability_id where r.availability_id = 1;

create table availability(
	id int not null auto_increment,
    name varchar(50),
    primary key(id)
);

INSERT INTO availability(name) values
	('available'),
    ('not available');

SELECT v.identity_numb, v.name, v.address, re.check_in, re.check_out, ro.room_numb, ro.price*(day(re.check_out)-day(re.check_in)) FROM rent re JOIN visitor v ON re.id_visitor = v.id JOIN room ro ON re.id_room = ro.id;

insert into room(type_id, room_numb, price, availability_id) values
	(1, 'ST01', 500000, 2),
    (1, 'ST02', 500000, 1),
    (1, 'ST03', 500000, 1),
    (2, 'SP01', 1500000, 1),
    (2, 'SP02', 1500000, 1),
    (2, 'SP03', 1500000, 2),
    (3, 'PR01', 3000000, 2),
    (3, 'PR02', 3000000, 2),
    (3, 'PR03', 3000000, 1);

INSERT INTO room(room_numb, price, type_id, available) VALUES (?, ?, ?, ?);

UPDATE room SET room_numb = ?, price = ?, type_id = ?, available = ? WHERE id = ?;

select * from room;

create table room_type(
	id int not null auto_increment,
	name_room varchar(100),
    primary key(id)
);

create table visitor(
	id int not null auto_increment,
    identity_numb varchar(100),
    name varchar(100),
    address varchar(100),
    hp_numb varchar(20),
    primary key(id)
);

INSERT INTO visitor(identity_numb, name, address, hp_numb) VALUES
	


create table rent(
	id int not null auto_increment,
    check_in date,
    check_out date,
    id_visitor int,
    id_room int,
    primary key(id),
    foreign key(id_visitor) references visitor(id),
    foreign key(id_room) references room(id)
);

insert into room_type(name_room) values
	('Standart Room'),
    ('Superior Room'),
    ('Presidential Room');

SELECT rt.name_room, r.room_numb, r.price, r.available FROM room r JOIN room_type rt ON r.type_id = rt.id;

select * from room;