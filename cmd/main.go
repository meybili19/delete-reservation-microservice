package main

import (
	"log"
	"net/http"

	"github.com/meybili19/delete-reservation-microservice/config"
	"github.com/meybili19/delete-reservation-microservice/routes"
)

func main() {
	databases, err := config.InitDatabases()
	if err != nil {
		log.Fatalf("Error initializing databases: %v", err)
	}
	defer func() {
		for _, db := range databases {
			db.Close()
		}
	}()
	log.Println("All databases connected successfully!")

	http.HandleFunc("/reservations/delete", routes.DeleteReservationHandler(databases))
	log.Println("Server running on port 8082")
	log.Println("http://localhost:8082/reservations/delete")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
