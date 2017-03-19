package main

import (
    "encoding/json"
    "log"
    "net/http"
    //"database/sql"

    // "fmt"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
)

type Patient struct{
  Gender int
  Id,Name,Dob,Address string
}
type Person struct {
    ID        string   `json:"id,omitempty"`
    Information   *Information `json:"information,omitempty"`
    Emergencycontact  *Emergencycontact   `json:"emergencycontact,omitempty"`
    HistoryArray []HistoryInfo `json:"historyarray,omitempty"`
}

type Information struct {
    Fullname string   `json:"fullname,omitempty"`
    Gender string   `json:"gender,omitempty"`
    Address string   `json:"address,omitempty"`
    Birth   *Birth `json:"birth,omitempty"`
}

type Birth struct {
  Month int `json:"month,omitempty"`
  Day int `json:"day,omitempty"`
  Year int `json:"year,omitempty"`
}

type Emergencycontact struct {
  Econtact string   `json:"econtact,omitempty"`
  Phone string   `json:"phone,omitempty"`
}

type HistoryArray struct{
  Collection []HistoryInfo `json:"historyarray"`
}

type HistoryInfo struct {
    HospitalName string `json:"hospitalname"`
}
/*
type PublicKey struct {
    Id int
    Key string
}

type KeysResponse struct {
    Collection []PublicKey
}

type YourJson struct {
    YourSample []struct {
        data map[string]string
    }
}
*/
var patient []Person


func GetPatientEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)

    db, err := gorm.Open("postgres", "postgresql://janitor_dev@ip-172-31-6-7.us-west-2.compute.internal:26257?sslcert=/home/ubuntu/certs/janitor_dev.cert&sslkey=/home/ubuntu/certs/janitor_dev.key")
    if err != nil {
      log.Fatalf("error connection to the database: %s", err)
    }

    var pat Patient
    db.Table("nwhacks.patients").Where("id = ?", params["id"]).Find(&pat)


    // fmt.Printf("%d", item.ID)
    // fmt.Printf("%s", pat.Id)
    json.NewEncoder(w).Encode(pat)
     defer db.Close()
    // json.NewEncoder(w).Encode(patient)
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
