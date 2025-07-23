package routes

import (
	"anime/handlers"
	"anime/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.MainHandler).Methods("GET")
	r.HandleFunc("/detail/{id:[0-9]+}", handlers.DetailHandler).Methods("GET")
	r.Handle("/", middleware.AuthMiddleware(http.HandlerFunc(handlers.HomeHandler))).Methods("GET")

	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")

	return r

}
