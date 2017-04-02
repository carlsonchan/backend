package main

import (
	"time"
)

type DatabasePatient struct {
	Id, Name, Address string
	Gender            int
	Dob               time.Time
}

type DatabaseEmergencyContact struct {
	Id, Pid, Name, Phone string
}

func (dbModel DatabaseEmergencyContact) toView() EmergencyContactView {
	var view EmergencyContactView

	view.Id = dbModel.Id
	view.Pid = dbModel.Pid
	view.Name = dbModel.Name
	view.Phone = dbModel.Phone

	return view
}
