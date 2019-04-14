package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func trackCounselorTicket(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering trackCounelorTicket")
	if r.Method != http.MethodPost {
		fmt.Printf("Non-post request received: %s\n", r.Method)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error reading body:\n%s", err)
		return
	}

	var ticket TitoTicket
	err = json.Unmarshal(body, &ticket)
	if err != nil {
		log.Printf("Error unmarshaling json:\n%s", err)
		return
	}

	var contactlist *AcContactListPayload
	if isTicketCounselor(ticket.Release.Slug) {
		contactlist, err = getAcContactByEmail(ticket.Email)
		if err != nil {
			log.Printf("Error looking up AC contact using email:\n%s\n", err)
			return
		}

		if len(contactlist.Contacts) < 1 {
			// write to list of missing emails
			_ = atCreateRecord(ticket.Name, ticket.Email, "Email not found in AC")

		} else if len(contactlist.Contacts) > 1 {
			// write to list email return multiple
		} else {
			// tag contact at AC
			year := strconv.Itoa(time.Now().Year())
			tagName := "SpeakerRegistered" + year
			if !addTagToAcContact(contactlist.Contacts[0], tagName) {
				log.Printf("Failed to add tag, %s, to contact, %+v\n", tagName, contactlist.Contacts[0])
			} else {
				log.Printf("Tag (%s) added to contact (%s)\n", tagName, contactlist.Contacts[0].Email)
			}
		}

	}
	rbody := fmt.Sprintf("Hello %s (%s)", ticket.Name, ticket.Email)
	w.Write([]byte(rbody))
	log.Println("trackCounelorTicket end")
}
