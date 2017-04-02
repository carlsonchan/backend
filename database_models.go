package main

import (
	"time"
)

type Patient struct {
	Id, Name, Address string
	Gender            int
	Dob               time.Time
}

type EmergencyContact struct {
	Id, Pid, Name, Phone string
}
