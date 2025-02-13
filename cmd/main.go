package main

import (
	"log"
	"net/http"

	"github.com/meybili19/delete-reservation-microservice/config"
	"github.com/meybili19/delete-reservation-microservice/routes"
	"github.com/rs/cors"
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

	// Crear un mux (router) para manejar rutas
	mux := http.NewServeMux()
	mux.HandleFunc("/reservations/delete", routes.DeleteReservationHandler(databases))

	// 🟢 Habilitar CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Permitir solicitudes solo desde frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Envolver el mux con el middleware de CORS
	handler := corsHandler.Handler(mux)

	log.Println("Server running on port 4002")
	log.Println("http://localhost:4002/reservations/delete")
	log.Fatal(http.ListenAndServe(":4002", handler))
}
