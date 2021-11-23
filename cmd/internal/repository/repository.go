package repository

import (
	"learningGo/cmd/internal/modules"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res modules.Reservation) (int, error)
	InsertRoomRestriction(r modules.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomId int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]modules.Room, error)
	GetRoomByID(id int) (modules.Room, error)

	UpdateUser(u modules.User) error
	GetUserByID(id int) (modules.User, error)
	Authenticate(email, testPassword string) (int, string, error)

	AllReservations() ([]modules.Reservation, error)
	AllNewReservations() ([]modules.Reservation, error)
}
