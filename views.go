package main

type BirthView struct {
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
	Year  int `json:"year,omitempty"`
}

type InformationView struct {
	Fullname string     `json:"fullname,omitempty"`
	Gender   string     `json:"gender,omitempty"`
	Address  string     `json:"address,omitempty"`
	Birth    *BirthView `json:"birth,omitempty"`
}

type EmergencyContactView struct {
	Id, Pid, Name, Phone string
}

type HistoryInfoView struct {
	HospitalName string `json:"hospitalname"`
}

type PersonView struct {
	ID                string                 `json:"id,omitempty"`
	Information       *InformationView       `json:"information,omitempty"`
	EmergencyContacts []EmergencyContactView `json:"emergency_contacts,omitempty"`
	HistoryArray      []HistoryInfoView      `json:"historyarray,omitempty"`
}
