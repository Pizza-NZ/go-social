package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Hello, World"))
			return
		case "/users":
			w.Write([]byte("users page"))
			return
		default:
			w.Write([]byte("404 page"))
		}
	default:
		w.Write([]byte("404 page"))
	}
}

func main() {
	api := &api{addr: ":8080"}

	srv := &http.Server{
		Addr:    api.addr,
		Handler: api,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
