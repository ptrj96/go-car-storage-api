package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ptrj96/go-car-storage-api/listings"
	"github.com/ptrj96/go-car-storage-api/logging"
)

func main() {
	logger := logging.GetLogger()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Print("hit root route")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"welcome"}`))
	}).Methods("GET")
	r.HandleFunc("/listings", listings.FindListingsHandler).Methods("POST")

	port := getRouterPort()

	logger.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getRouterPort() string {
	if value, exists := os.LookupEnv("APP_PORT"); exists {
		return value
	}
	return "8083"
}
