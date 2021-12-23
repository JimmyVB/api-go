// Package classification Petstore API.
//
// Documentation for Person API
//
//     Schemes: http
//     Host: localhost
//     BasePath: /v1
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
// swagger:meta
package handler

import (
	"api-go/connection"
	"api-go/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// A list of persons return in the response
// swagger:response personsResponse
type personsResponse struct {
	// All persons in the db
	// in: body
	Body []models.Person
}

// swagger:route GET /persons/api all listPersons
// Return a list of persons from the database
// responses:
//	200: personsResponse

// GetAll returns the persons from the db
func GetAll(writer http.ResponseWriter, request *http.Request) {
	var persons []models.Person
	db := connection.DB()

	db.Find(&persons)
	json, _ := json.Marshal(persons)
	SendReponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	person := models.Person{}

	id := mux.Vars(request)["id"]

	db := connection.DB()
	db.Find(&person, id)

	if person.ID > 0 {
		json, _ := json.Marshal(person)
		SendReponse(writer, http.StatusOK, json)
	} else {
		SendError(writer, http.StatusNotFound)
	}
}

func Save(writer http.ResponseWriter, request *http.Request) {
	person := models.Person{}

	db := connection.DB()
	error := json.NewDecoder(request.Body).Decode(&person)

	if error != nil {
		log.Fatal(error)
		SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&person).Error
	if error != nil {
		log.Fatal(error)
		SendError(writer, http.StatusInternalServerError)
		return
	}
	json, _ := json.Marshal(person)
	SendReponse(writer, http.StatusCreated, json)
}

func Update(writer http.ResponseWriter, request *http.Request) {

	person := models.Person{}
	newPerson := models.Person{}

	id := mux.Vars(request)["id"]
	db := connection.DB()

	db.Find(&person, id)

	if person.ID > 0 {
		error := json.NewDecoder(request.Body).Decode(&newPerson)
		if error != nil {
			log.Fatal(error)
			SendError(writer, http.StatusBadRequest)
			return
		}
		db.Model(&person).Updates(&newPerson)
		json, _ := json.Marshal(person)
		SendReponse(writer, http.StatusCreated, json)
	} else {
		SendError(writer, http.StatusInternalServerError)
		return
	}
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	person := models.Person{}

	db := connection.DB()
	id := mux.Vars(request)["id"]

	db.Find(&person, id)

	if person.ID > 0 {
		db.Delete(person)
		SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		SendError(writer, http.StatusNotFound)
	}
}
