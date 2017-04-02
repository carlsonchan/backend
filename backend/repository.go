package main

import (
	"github.com/jinzhu/gorm"
	"log"
)

var database *gorm.DB

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
