package profileapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"jwt-go/model"

	"github.com/gorilla/mux"
)

// CreateProfile create the profile
// POST /profile
func (api *api) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var profile model.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	defer r.Body.Close()
	if err != nil {
		w.Write([]byte("ERROR: Error while decoding JSON data"))
		return
	}

	err = api.ProfileService.CreateProfile(profile)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// GetProfile return the profile according to passed id
// GET /profile/{id}
func (api *api) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	profile, err := api.ProfileService.GetProfile(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// UpdateProfile updates the profile using PUT method according to passed id
// PUT /profile/{id}
func (api *api) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var profile model.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		w.Write([]byte("ERROR: Error while decoding JSON data"))
	}

	id := mux.Vars(r)["id"]

	err = api.ProfileService.UpdateProfile(profile, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// DeleteProfile delete the profile according to passed id
// DELETE /profile/{id}
func (api *api) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := api.ProfileService.DeleteProfile(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Profile Deleted with ID: " + id))
}

// GetAllTask return all the task attached with particular user
// GET /profile/task/{id}
func (api *api) GetAllTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	tasks, err := api.ProfileService.GetAllTask(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for _, task := range tasks {
		w.Write([]byte(fmt.Sprintf("Title: %v\nBody: %v\nMade By: %v\n\n", task.Title, task.Body, task.ProfileId)))
	}
}
