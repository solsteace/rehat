package services

import (
	"github.com/solsteace/rest/models"
	"github.com/solsteace/rest/repositories"
)

type Reservation struct {
	repositories.Room
	repositories.Reservation
	repositories.Class
}

func (r Reservation) Add(userId, roomId int64) (models.Reservation, error) {
	var reservation models.Reservation
	// room, err := r.Room.GetById(roomId)
	// if err != nil {
	// 	return reservation, err
	// }

	// if !room.IsVacant() {
	// 	return reservation, &ErrIllegalOperation{reason: "Room is not vacant"}
	// }
	return reservation, nil
}

func (r Reservation) EditById(
	reservationId int64,
	reservation models.Reservation,
) error {
	// if _, err := r.Reservation.GetById(reservationId); err != nil {
	// 	return err
	// }

	// if err := r.Reservation.EditById(reservationId, &reservation); err != nil {
	// 	return err
	// }
	return nil
}

func (r Reservation) DeleteById(reservationId int64) error {
	// reservation, err := r.Reservation.GetById(reservationId)
	// if err != nil {
	// 	return err
	// }

	// if err := r.Reservation.DeleteById(reservationId); err != nil {
	// 	return err
	// }

	// roomId := reservation.RoomID
	// room, err := r.Room.GetById(roomId)
	// if err != nil {
	// 	return err
	// }
	// room.Status = "open"

	// if err := r.Room.EditById(roomId, &room); err != nil {
	// 	return err
	// }

	return nil
}
