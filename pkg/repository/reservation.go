package repository

import (
	"mark_emailchaser/pkg/common"
	"mark_emailchaser/pkg/models"
)

func CreateReservation(reservation *models.Reservation) error {
	return common.DB.Create(reservation).Error
}

func FindReservationByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation
	err := common.DB.First(&reservation, id).Error
	return &reservation, err
}

func UpdateReservation(updateReservation *models.Reservation) error {
	var reservation models.Reservation
	reservation.UserID = updateReservation.UserID
	reservation.Restaurant = updateReservation.Restaurant
	reservation.ReservationTime = updateReservation.ReservationTime
	reservation.NumGuests = updateReservation.NumGuests
	return common.DB.Save(reservation).Error
}

func DeleteReservation(id uint) error {
	return common.DB.Delete(&models.Reservation{}, id).Error
}
