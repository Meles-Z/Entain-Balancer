package server

import (
	"net/http"
	"strings"

	"github.com/meles-z/entainbalancer/internal/handlers"
)

func Route(h *handlers.Handler) {
	// Handles /user/{userId}/transaction and /user/{userId}/balance
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/") // Remove leading/trailing slashes
		parts := strings.Split(path, "/")     // Split by "/"

		if len(parts) == 3 && parts[0] == "user" {
			switch parts[2] {
			case "transaction":
				if r.Method == http.MethodPost {
					h.UpdateTransaction(w, r)
					return
				}
			case "balance":
				if r.Method == http.MethodGet {
					h.GetUserBalance(w, r)
					return
				}
			}
		}

		http.Error(w, "Not Found", http.StatusNotFound)
	})
}
