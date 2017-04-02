package main

import (
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	_ "strings"
	"time"
)

type ConfigDatabase struct {
	User, Ip, Port, SslCertLocation, SslKeyLocation string
}
type Config struct {
	Port, OsUser string
	Database     ConfigDatabase
}

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

var config Config
var database *gorm.DB

func InitializeConfiguration() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Error opening the configuration file: ", err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("Error reading the configuration file: ", err)
	}
}

func InitializeDbConnection() *gorm.DB {
	dbConnection := "postgresql://" + config.Database.User + "@" +
		config.Database.Ip + ":" + config.Database.Port +
		"/NWHACKS?sslcert=" + config.Database.SslCertLocation +
		"&sslkey=" + config.Database.SslKeyLocation +
		"&parseTime=true"

	db, err := gorm.Open("postgres", dbConnection)
	if err != nil {
		log.Fatalf("Error connection to the database: %s", err)
	}
	return db
}

func GetPatientEndpoint(w http.ResponseWriter, req *http.Request) {
	log.Print("Mapped " + req.Method + " " + req.URL.Path)

	params := mux.Vars(req)

	var rawPatient Patient
	result := database.Table("patients").Where("id = ?", params["id"]).Find(&rawPatient)
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
