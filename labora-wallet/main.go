package main

import (
	"log"

	"github.com/Zumari/labora-wallet/labora-wallet/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// router.HandleFunc("/CreateWallet", controller.CreateWallet).Methods("POST")
	// router.HandleFunc("/UpdateWallet", controller.UpdateWallet).Methods("PUT")
	// router.HandleFunc("/DeleteWallet", controller.DeleteWallet).Methods("DELETE")
	// router.HandleFunc("/WalletStatus", controller.WalletStatus).Methods("GET")
	// router.HandleFunc("/GetLogs", controller.GetLogs).Methods("GET")

	// Configure CORS middleware
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
	)

	// Add CORS middleware to all routes
	handler := corsOptions(router)

	//Server connection
	if err := config.StartServer(handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
