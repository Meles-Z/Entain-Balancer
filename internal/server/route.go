package server

import (
	"net/http"

	"github.com/gorilla/mux"
	httpHandler "github.com/meles-z/entainbalancer/internal/delivery/http"
)

func Route(h *httpHandler.Handler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/user/{userId}/transaction", h.UpdateTransaction).Methods("POST")
	r.HandleFunc("/user/{userId}/balance", h.GetUserBalance).Methods("GET")

	return r
}
