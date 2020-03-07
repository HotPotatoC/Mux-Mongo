package userhandler

import "github.com/gorilla/mux"

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")

	return r
}
