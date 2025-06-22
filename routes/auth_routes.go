package routes

import (
	"github.com/gorilla/mux"
	"belajar-golang-api/controllers"
)

func SetupAuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	r.HandleFunc("/checkdata", controllers.CheckDataHandler).Methods("GET")
}
