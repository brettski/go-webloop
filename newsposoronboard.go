package main

import (
	"fmt"
	"net/http"
)

func newSponsorOnBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entering newSponsorOnBoard")
	if r.Method != http.MethodPost {
		fmt.Printf("Non-post request received: %s", r.Method)
		return
	}

	fields, err := parseAcPostHook(r)
	if err != nil {
		//respond error
	}
	fmt.Println(fields)

}
