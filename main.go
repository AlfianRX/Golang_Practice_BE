package main

import (
	"fmt"
	"latihan_golang/database"
	"latihan_golang/pkg/mysql"
	"latihan_golang/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mysql.DatabaseInit()

	database.RunMigration()

	route := mux.NewRouter()

	routes.RouteInit(route.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running")
	http.ListenAndServe("localhost:5000", route)
}
