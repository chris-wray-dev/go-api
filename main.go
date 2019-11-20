package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "GET called"}`))
}

func main() {

	router := mux.NewRouter()
	router.Use() //attach JWT auth middleware

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "9080" //localhost
	}

	fmt.Println(port)

	router.HandleFunc("/api", get).Methods("GET")

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
