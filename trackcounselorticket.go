package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	fmt.Println("result:")
	//fmt.Printf("%+v\n", ticket)
	rbody := fmt.Sprintf("Hello %s\n%s", ticket.Name, ticket.Email)
	w.Write([]byte(rbody))
}
