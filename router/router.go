package router

import (
	"demo/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(server *handlers.Server) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/employees", server.HandleCreateEmployee).Methods("POST")
	r.HandleFunc("/employees", server.HandleListEmployees).Methods("GET")
	r.HandleFunc("/employee/{id}", server.HandleGetEmployeeByID).Methods("GET")
	r.HandleFunc("/employee/{id}", server.HandleUpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{id}", server.HandleDeleteEmployee).Methods("DELETE")

	return r
}
