package main

import (
	"fmt"
	"net/http"
)

func trackCounelorTicket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entering trackCounelorTicket")
	if r.Method != http.MethodPost {
		fmt.Printf("Non-post request received: %s", r.Method)
		return
	}

	// Fields
}
