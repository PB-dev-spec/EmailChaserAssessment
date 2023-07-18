package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID          uint      `gorm:"not null" json:"user_id"`
	Restaurant      string    `gorm:"not null" json:"restaurant_id"`
	ReservationTime time.Time `gorm:"not null" json:"reservation_time"`
	NumGuests       int       `gorm:"not null" json:"reservation_numguests"`
}
