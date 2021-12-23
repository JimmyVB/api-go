package main

import (
	"api-go/connection"
	"api-go/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	driver := connection.Postgres
	connection.New(driver)

	//	connection.DB().AutoMigrate(&models.Person{})

	router := mux.NewRouter()
	routes.RoutePerson(router)

	server := http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	log.Println("Servidor iniciado en el puerto 8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
