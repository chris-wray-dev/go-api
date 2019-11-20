package main

import (
	"fmt"
	"go-api/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	fmt.Println(port)

	router.HandleFunc("/api/locations/{location_id}", controllers.FetchLocationById).Methods("GET")
	router.HandleFunc("/api/locations", controllers.FetchAllLocations).Methods("GET")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
