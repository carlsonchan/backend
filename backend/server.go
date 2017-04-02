package main

import (
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	_ "strings"
)

func GetPatientEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Print("Mapped " + req.Method + " " + req.URL.Path)

	params := mux.Vars(req)

	rawPatient, result := GetPatientById(params["id"])
	if result.RecordNotFound() {
		log.Print("Patient " + params["id"] + " not found.")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var gender string
	if rawPatient.Gender == 0 {
		gender = "M"
	} else if rawPatient.Gender == 1 {
		gender = "F"
	} else {
		gender = "O"
	}

	birth := &BirthView{
		Month: int(rawPatient.Dob.Month()),
		Day:   int(rawPatient.Dob.Day()),
		Year:  rawPatient.Dob.Year(),
	}

	contactList, _ := GetEmergencyContactsByPatientId(params["id"])

	var contactViewList []EmergencyContactView
	for i := 0; i < len(contactList); i++ {
		contactViewList = append(contactViewList, contactList[i].toView())
	}

	patient := PersonView{
		ID: rawPatient.Id,
		Information: &InformationView{
			Fullname: rawPatient.Name,
			Gender:   gender,
			Address:  rawPatient.Address,
			Birth:    birth,
		},
		EmergencyContacts: contactViewList,
	}

	json.NewEncoder(w).Encode(patient)
}

func main() {
	log.Print("Initializing server.")

	log.Print("Initializing configuration.")
	InitializeConfiguration()

	log.Print("Initializing database connection.")
	database = InitializeDbConnection()
	defer database.Close()

	router := mux.NewRouter()
	router.HandleFunc("/patient/{id}", GetPatientEndpoint).Methods("GET")

	log.Print("Starting server on port " + config.Port + ".")
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
