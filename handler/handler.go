package handler

import (
	"elibrary/service"
	"elibrary/transport"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupHandlers sets up all the routes for the application
func SetupHandlers(services *service.ServiceImpl, router *mux.Router) *http.ServeMux {
	handler := http.NewServeMux()

	// Registering HTTP routes
	//GET
	router.HandleFunc("/Book", func(w http.ResponseWriter, r *http.Request) {
		transport.GetBookHandler(services, w, r)
	}).Methods("GET")
	// Add more routes

	return handler
}
