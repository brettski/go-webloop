package main

// Active Campain payload when filtering contacts by email

// AcContactListPayload struct modeling payload received from AC filtered contact list
type AcContactListPayload struct {
	Contacts []struct {
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Id        string `json:"id"`
		Org       string `json:"organization"`
	} `json:"contacts"`
}
