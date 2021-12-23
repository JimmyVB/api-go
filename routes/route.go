package routes

import (
	"api-go/handler"
	"github.com/gorilla/mux"
)

func RoutePerson(router *mux.Router) {

	subRoute := router.PathPrefix("/persons/api").Subrouter()
	subRoute.HandleFunc("/all", handler.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", handler.Save).Methods("POST")
	subRoute.HandleFunc("/update/{id}", handler.Update).Methods("PUT")
	subRoute.HandleFunc("/delete/{id}", handler.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", handler.Get).Methods("GET")

}
