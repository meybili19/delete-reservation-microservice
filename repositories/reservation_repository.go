package repositories

import (
	"database/sql"
	"fmt"
)

func DeleteReservation(db *sql.DB, id int) error {
	query := "DELETE FROM Reservations WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting reservation: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error fetching rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no reservation found with ID %d", id)
	}
	return nil
}
