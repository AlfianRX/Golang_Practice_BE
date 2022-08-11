package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(route *mux.Router) {
	TodoRoutes(route)
}
