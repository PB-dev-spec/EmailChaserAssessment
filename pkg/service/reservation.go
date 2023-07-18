package service

import (
	"mark_emailchaser/pkg/models"
	"mark_emailchaser/pkg/repository"
)

func CreateReservation(newReservation models.Reservation) (*models.Reservation, error) {

	if err := repository.CreateReservation(&newReservation); err != nil {
		return nil, err
	}

	return &newReservation, nil
}

func GetReservationByID(id uint) (*models.Reservation, error) {
	return repository.FindReservationByID(id)
}

func DeleteReservation(id uint) error {
	return repository.DeleteReservation(id)
}

func UpdateReservation(reservation models.Reservation) error {
	return repository.UpdateReservation(&reservation)
}
