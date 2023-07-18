package main

import (
	"log"

	"mark_emailchaser/pkg/common"
	"mark_emailchaser/pkg/models"
)

func main() {
	common.LoadEnvVariables()
	common.GetDBConnection()
	err := common.DB.AutoMigrate(
		&models.User{},
		&models.Reservation{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
