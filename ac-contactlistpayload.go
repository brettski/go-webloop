package main

// Active Campain payload when filtering contacts by email

// AcContact Active Campaign contact
type AcContact struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Id        string `json:"id"`
	Org       string `json:"organization"`
}

// AcContactListPayload struct modeling payload received from AC filtered contact list
type AcContactListPayload struct {
	Contacts []AcContact `json:"contacts"`
}
