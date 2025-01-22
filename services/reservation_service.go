package services

import (
	"database/sql"
	"fmt"

	"github.com/meybili19/delete-reservation-microservice/repositories"
)

func DeleteReservationService(db *sql.DB, id int) error {
	if db == nil {
		return fmt.Errorf("database not initialized")
	}
	return repositories.DeleteReservation(db, id)
}
