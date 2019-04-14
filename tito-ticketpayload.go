package main

// Contains models for tito 'ticket.complete' webhook payload

// TitoTicket payload data sruct
type TitoTicket struct {
	Text  string `json:"text"`
	Id    int64  `json:"id"`
	Event struct {
		Id    int64  `json:"id"`
		Title string `json:"title"`
		Url   string `json:"url"`
	} `json:"event"`
	Name         string   `json:"name"`
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone_number"`
	Company      string   `json:"company_name"`
	Reference    string   `json:"reference"` // Ticket purchase reference e.g OXBD-1
	Slug         string   `json:"slug"`
	StateName    string   `json:"state_name"` // State of ticket. e.g. complete
	DiscountCode string   `json:"discount_code_used"`
	Release      struct { // Ticket type ("the released ticket for purchase")
		Id    int64  `json:"id"`
		Title string `json:"title"`
		Slug  string `json:"slug"`
	} `json:"release"`
	RegistrationId   string `json:"registration_id"`
	RegistrationSlug string `json:"registration_slug`
	Custom           string `json:"custom"`
}
