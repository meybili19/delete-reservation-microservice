package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/meybili19/delete-reservation-microservice/services"
)

func DeleteReservationHandler(databases map[string]*sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		idParam := r.URL.Query().Get("id")
		if idParam == "" {
			http.Error(w, "Missing reservation ID", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid reservation ID", http.StatusBadRequest)
			return
		}

		if err := services.DeleteReservationService(databases["reservations"], id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Reservation deleted successfully",
		})
	}
}
