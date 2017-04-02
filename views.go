package main

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

type HistoryInfo struct {
	HospitalName string `json:"hospitalname"`
}

type Person struct {
	ID                string             `json:"id,omitempty"`
	Information       *Information       `json:"information,omitempty"`
	EmergencyContacts []EmergencyContact `json:"emergency_contacts,omitempty"`
	HistoryArray      []HistoryInfo      `json:"historyarray,omitempty"`
}
