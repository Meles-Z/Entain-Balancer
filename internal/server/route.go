package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meles-z/entainbalancer/internal/handlers"
)

func Route(h *handlers.Handler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/user/{userId}/transaction", h.UpdateTransaction).Methods("POST")
	r.HandleFunc("/user/{userId}/balance", h.GetUserBalance).Methods("GET")

	return r
}
