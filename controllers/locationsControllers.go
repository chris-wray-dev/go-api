package controllers

import (
	"go-api/models"
	u "go-api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var FetchAllLocations = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllLocations()
	resp := u.Message(true, "success")
	resp["locations"] = data
	u.Respond(w, resp)
}

var FetchLocationById = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["location_id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetLocationById(int(id))
	resp := u.Message(true, "success")
	resp["location"] = data
	u.Respond(w, resp)
}
