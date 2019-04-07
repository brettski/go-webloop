package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func trackCounselorTicket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entering trackCounelorTicket")
	if r.Method != http.MethodPost {
		fmt.Printf("Non-post request received: %s\n", r.Method)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Printf("Error reading body:\n%s", err)
		return
	}

	var ticket TitoTicket
	err = json.Unmarshal(body, &ticket)
	if err != nil {
		fmt.Printf("Error unmarshaling json:\n%s", err)
		return
	}

	var contactlist *AcContactListPayload
	if isTicketCounselor(ticket.Release.Slug) {
		contactlist, err = getAcContactByEmail(ticket.Email)
		if err != nil {
			fmt.Printf("Error looking up AC contact using email:\n%s\n", err)
			return
		}

		if len(contactlist.Contacts) < 1 {
			// write to list of missing emails
		} else if len(contactlist.Contacts) > 1 {
			// write to list email return multiple
		} else {
			// tag contact at AC
			year := strconv.Itoa(time.Now().Year())
			tagName := "SpeakerRegistered" + year
			if !addTagToAcContact(contactlist.Contacts[0], tagName) {
				fmt.Printf("Failed to add tag, %s, to contact, %+v\n", tagName, contactlist.Contacts[0])
			} else {
				fmt.Printf("Tag (%s) added to contact (%s)\n", tagName, contactlist.Contacts[0].Email)
			}
		}

	}
	rbody := fmt.Sprintf("Hello %s (%s)", ticket.Name, ticket.Email)
	w.Write([]byte(rbody))

}
