package taskapi

import (
	"encoding/json"
	"jwt-go/model"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateTask create the task
// POST /task
func (api *api) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	defer r.Body.Close()
	if err != nil {
		w.Write([]byte("ERROR: Error while decoding JSON data"))
	}

	err = api.TaskService.CreateTask(task)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// GetOneTask return the task according to passed id
// GET /task/{id}
func (api *api) GetTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	task, err := api.TaskService.GetTask(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// UpdateTask updates the task using PUT method according to passed id
// PUT /task/{id}
func (api *api) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.Write([]byte("ERROR: Error while decoding JSON data"))
	}

	id := mux.Vars(r)["id"]

	err = api.TaskService.UpdateTask(task, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// DeleteTask delete the task according to passed id
// DELETE /task/{id}
func (api *api) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := api.TaskService.DeleteTask(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Task Deleted with ID: " + id))
}
