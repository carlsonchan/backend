package main

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

var database *gorm.DB

func InitializeDbConnection() {
	dbConnection := "postgresql://" + config.Database.User + "@" +
		config.Database.Ip + ":" + config.Database.Port +
		"/NWHACKS?sslcert=" + config.Database.SslCertLocation +
		"&sslkey=" + config.Database.SslKeyLocation +
		"&parseTime=true"

	err := errors.New("")
	database, err = gorm.Open("postgres", dbConnection)
	if err != nil {
		log.Fatalf("Error connection to the database: %s", err)
	}
}

func CloseDbConnection() {
	database.Close()
}

func GetPatientById(id string) (DatabasePatient, *gorm.DB) {
	var patient DatabasePatient
	result := database.
		Table("patients").
		Where("id = ?", id).
		Find(&patient)
	return patient, result
}

func GetEmergencyContactsByPatientId(id string) ([]DatabaseEmergencyContact, *gorm.DB) {

	var emergencyContacts []DatabaseEmergencyContact
	result := database.
		Table("emergency_contacts").
		Where("pid = ?", id).
		Find(&emergencyContacts)
	return emergencyContacts, result
}
