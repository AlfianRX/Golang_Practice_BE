package routes

import (
	"latihan_golang/handlers"
	"latihan_golang/pkg/mysql"
	"latihan_golang/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(route *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	route.HandleFunc("/users", h.FindUsers).Methods("GET")
	route.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	route.HandleFunc("/user", h.CreateUser).Methods("POST")
	route.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	route.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}
