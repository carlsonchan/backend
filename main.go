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
	Gender                 int
	Id, Name, Dob, Address string
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

var patient []Person

func GetPatientEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	db, err := gorm.Open("postgres", "postgresql://janitor_dev@ip-172-31-6-7.us-west-2.compute.internal:26257/NWHACKS?sslcert=/home/ubuntu/certs/janitor_dev.cert&sslkey=/home/ubuntu/certs/janitor_dev.key")
	if err != nil {
		log.Fatalf("error connection to the database: %s", err)
	}
	defer db.Close()

	var rawPatient Patient
	db.Table("patients").Where("id = ?", params["id"]).Find(&rawPatient)

	var gender string
	if rawPatient.Gender == 0 {
		gender = "M"
	} else if rawPatient.Gender == 1 {
		gender = "F"
	} else {
		gender = "O"
	}

	birthTime, err := time.Parse(rawPatient.Dob, "2011-01-19")
	birth := &Birth{
		Month: int(birthTime.Month()),
		Day:   int(birthTime.Day()),
		Year:  birthTime.Year(),
	}

	var contact EmergencyContact
	var contactList []EmergencyContact
	// db.Table("nwhacks.patients").Where("id = ?", params["id"]).Find(&pat)
	// db.Model(&pat).Related(&emer, "EmergencyContact")

	db.Find(&contactList)
	db.Table("NWHACKS.emergency_contacts").Where("pid = ?", rawPatient.Id).Find(&contact)

	patient := Person {
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

// func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
//     params := mux.Vars(req)
//     var person Person
//     _ = json.NewDecoder(req.Body).Decode(&person)
//     person.ID = params["id"]
//     patient = append(patient, person)
//     json.NewEncoder(w).Encode(patient)
// }

// func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
//     params := mux.Vars(req)
//     for index, item := range patient {
//         if item.ID == params["id"] {
//             patient = append(patient[:index], patient[index+1:]...)
//             break
//         }
//     }
//     json.NewEncoder(w).Encode(patient)
// }

func main() {

	router := mux.NewRouter()
	/*

	   patient = append(patient, Person{
	     ID: "1",
	     Information: &Information{Fullname: "Jacky Chao", Gender: "M", Address: "1234 UBC w.e.", Birth: &Birth{Day: 12, Month: 9, Year: 1993}},
	     Emergencycontact: &Emergencycontact{Econtact: "Carlson Chan", Phone: "123-456-7890"},

	     // TODO: Fix this
	     // HistoryArray: &HistoryInfo {HospitalName: ""}
	     // HistoryArray: HistoryArray{Collection: a [10]&HistoryInfo}
	     // HistoryArray: [1]HistoryInfo{&HospitalName: ""},
	   })

	*/
	// router.HandleFunc("/patient", GetpatientEndpoint).Methods("GET")
	router.HandleFunc("/patient/{id}", GetPatientEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8787", router))
}
