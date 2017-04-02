package main

import (
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	_ "strings"
	"time"
)

type Patient struct {
	Id, Name, Address string
	Gender            int
	Dob               time.Time
}

type Birth struct {
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
	Year  int `json:"year,omitempty"`
}

type Information struct {
	Fullname string `json:"fullname,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Address  string `json:"address,omitempty"`
	Birth    *Birth `json:"birth,omitempty"`
}

type EmergencyContact struct {
	Id, Pid, Name, Phone string
}

type HistoryInfo struct {
	HospitalName string `json:"hospitalname"`
}

type Person struct {
	ID                string             `json:"id,omitempty"`
	Information       *Information       `json:"information,omitempty"`
	EmergencyContacts []EmergencyContact `json:"emergency_contacts,omitempty"`
	HistoryArray      []HistoryInfo      `json:"historyarray,omitempty"`
}

const port string = "8787"

const osUser string = "jleung"
const dbUser string = "janitor_dev"
const dbIp string = "localhost"
const dbPort string = "26257"

var database *gorm.DB

func InitializeDbConnection() *gorm.DB {
	const sslCertLocation = "/home/" + osUser +
		"/cockroach/certs/janitor_dev.cert"
	const sslKeyLocation = "/home/" + osUser +
		"/cockroach/certs/janitor_dev.key"

	dbConnection := "postgresql://" + dbUser + "@" + dbIp + ":" + dbPort +
		"/NWHACKS?sslcert=" + sslCertLocation +
		"&sslkey=" + sslKeyLocation +
		"&parseTime=true"

	db, err := gorm.Open("postgres", dbConnection)
	if err != nil {
		log.Fatalf("error connection to the database: %s", err)
	}
	return db
}

func GetPatientEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Print("Mapped " + req.Method + " " + req.URL.Path)

	params := mux.Vars(req)

	var rawPatient Patient
	database.Table("patients").Where("id = ?", params["id"]).Find(&rawPatient)

	var gender string
	if rawPatient.Gender == 0 {
		gender = "M"
	} else if rawPatient.Gender == 1 {
		gender = "F"
	} else {
		gender = "O"
	}

	birth := &Birth{
		Month: int(rawPatient.Dob.Month()),
		Day:   int(rawPatient.Dob.Day()),
		Year:  rawPatient.Dob.Year(),
	}

	var contactList []EmergencyContact
	database.Table("emergency_contacts").Where("pid = ?", rawPatient.Id).Find(&contactList)

	patient := Person{
		ID: rawPatient.Id,
		Information: &Information{
			Fullname: rawPatient.Name,
			Gender:   gender,
			Address:  rawPatient.Address,
			Birth:    birth,
		},
		EmergencyContacts: contactList,
	}

	json.NewEncoder(w).Encode(patient)
}

func main() {
	log.Print("Initializing server.")

	log.Print("Initializing database connection.")
	database = InitializeDbConnection()
	defer database.Close()

	router := mux.NewRouter()
	router.HandleFunc("/patient/{id}", GetPatientEndpoint).Methods("GET")

	log.Print("Starting server on port " + port + ".")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
