package api

import (
	"workplaces/server/internals/app/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes(userHandler *handlers.UsersHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/create", userHandler.Create).Methods("POST")
	r.HandleFunc("/users/update/{id:[0-9]+}", userHandler.Update).Methods("PUT")
	r.HandleFunc("/users/delete/{id:[0-9]+}", userHandler.Delete).Methods("DELETE")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()
	return r
}
