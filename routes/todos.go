package routes

import (
	h "latihan_golang/handlers"

	"github.com/gorilla/mux"
)

func TodoRoutes(route *mux.Router) {
	route.HandleFunc("/todos", h.FindTodos).Methods("GET")
	route.HandleFunc("/todo/{id}", h.GetTodo).Methods("GET")
	route.HandleFunc("/todo", h.CreateTodo).Methods("POST")
	route.HandleFunc("/todo/{id}", h.UpdateTodo).Methods("PATCH")
	route.HandleFunc("/todo/{id}", h.DeleteTodo).Methods("DELETE")
}
