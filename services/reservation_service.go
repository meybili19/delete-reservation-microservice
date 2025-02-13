package services

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/meybili19/delete-reservation-microservice/repositories"
)

// IncrementParkingCapacity increases the parking lot capacity after deleting a reservation
func IncrementParkingCapacity(parkingLotID int) error {
	incrementURL := os.Getenv("PARKINGLOT_SERVICE_INCREMENT_URL")
	if incrementURL == "" {
		return fmt.Errorf("PARKINGLOT_SERVICE_INCREMENT_URL is not set")
	}

	url := fmt.Sprintf("%s/%d", incrementURL, parkingLotID)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		return fmt.Errorf("error creating request to increment capacity: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request to increment capacity: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error incrementing capacity: received status %d", resp.StatusCode)
	}

	return nil
}

func DeleteReservationService(db *sql.DB, id int) error {
	if db == nil {
		return fmt.Errorf("database not initialized")
	}

	reservation, err := repositories.GetReservationByID(id)
	if err != nil {
		return fmt.Errorf("error fetching reservation details: %w", err)
	}

	err = repositories.DeleteReservation(db, id)
	if err != nil {
		return fmt.Errorf("error deleting reservation: %w", err)
	}

	err = IncrementParkingCapacity(reservation.Data.GetReservationById.ParkingLotID)
	if err != nil {
		return fmt.Errorf("reservation deleted, but failed to increment capacity: %w", err)
	}

	return nil
}
