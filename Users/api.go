package users

import (
	"encoding/json"
	"net/http"
)

type UserAPI struct {
	Addr string
}

var usrs = []User{}

func (a *UserAPI) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(usrs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *UserAPI) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {

	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	usrs = append(usrs, u)

	w.WriteHeader(http.StatusCreated)
}
