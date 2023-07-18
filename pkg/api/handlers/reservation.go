package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"mark_emailchaser/pkg/api/responses"
	"mark_emailchaser/pkg/models"
	"mark_emailchaser/pkg/service"
)

func CreateReservation(ctx *gin.Context) {
	// Parse the request and validate the input
	var newReservation models.Reservation
	if err := ctx.ShouldBindJSON(&newReservation); err != nil {
		responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	reservation, err := service.CreateReservation(newReservation)
	if err != nil {
		responses.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"reservation": reservation})
}

func GetReservation(ctx *gin.Context) {
	reservationIDStr := ctx.Param("id")
	// Convert the reservation ID string to an integer.
	reservationIDInt, err := strconv.Atoi(reservationIDStr)
	if err != nil {
		// If the conversion failed, return an error response.
		responses.ErrorResponse(ctx, http.StatusBadRequest, "Invalid reservation ID format.")
		return
	}
	reservationID := uint(reservationIDInt)
	// Use the ReservationService to fetch the reservation from the database.
	reservation, err := service.GetReservationByID(reservationID)
	if err != nil {
		// If there was an error fetching the reservation, return an error response.
		responses.ErrorResponse(ctx, http.StatusInternalServerError, "Error fetching reservation information.")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"reservation": reservation})
}

func DeleteReservation(ctx *gin.Context) {
	reservationIDStr := ctx.Param("id")
	// Convert the reservation ID string to an integer.
	reservationIDInt, err := strconv.Atoi(reservationIDStr)
	if err != nil {
		// If the conversion failed, return an error response.
		responses.ErrorResponse(ctx, http.StatusBadRequest, "Invalid reservation ID format.")
		return
	}
	reservationID := uint(reservationIDInt)
	// Use the ReservationService to delete the reservation from the database.
	err = service.DeleteReservation(reservationID)
	if err != nil {
		// If there was an error deleting the reservation, return an error response.
		responses.ErrorResponse(ctx, http.StatusInternalServerError, "Error deleting reservation.")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation deleted successfully"})
}

func UpdateReservation(ctx *gin.Context) {
	var updatedReservation models.Reservation
	if err := ctx.ShouldBindJSON(&updatedReservation); err != nil {
		// If there was an error parsing the request body, return an error response.
		responses.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	// Use the ReservationService to update the reservation in the database.
	err := service.UpdateReservation(updatedReservation)
	if err != nil {
		// If there was an error updating the reservation, return an error response.
		responses.ErrorResponse(ctx, http.StatusInternalServerError, "Error updating reservation.")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "update reservation successfully"})
}
