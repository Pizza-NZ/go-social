package main

import (
	"net/http"

	"github.com/pizza-nz/go-social/users"
)

func main() {
	api := &users.UserAPI{Addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.Addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.GetUsersHandler)
	mux.HandleFunc("POST /users", api.CreateUsersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
