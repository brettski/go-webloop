package main

// Calls to get and send data to Active Campaign

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getAcContactByEmail(email string) (*AcContactListPayload, error) {
	if !strings.Contains(email, "@") {
		fmt.Println("Provided email not valid")
		return nil, errors.New("Provided email not valid")
	}

	env, err := getEnvironmentInfo()
	if err != nil {
		log.Fatal(err)
	}

	accountname := env.AcAccountName
	acbaseurl := env.AcBaseUrl
	acapikey := env.AcApiKey

	// "https://%s.api.-us1.com/api/3"
	url := fmt.Sprintf(acbaseurl+"/contacts?email_like=%s", accountname, email)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		msg := fmt.Sprintf("Error setting up request:\n%s\n", err)
		fmt.Println(msg)
		return nil, errors.New(msg)
	}
	req.Header.Add("Api_Token", acapikey)
	req.Header.Add("content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error requesting data from AC:\n%s\n", err)
		return nil, errors.New("Error requesting data from AC")
	}

	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("Response didn't return 200. Status received: %d", resp.StatusCode)
		fmt.Println(msg)
		return nil, errors.New(msg)
	}

	contact, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body from request:\n%s\n", err)
		return nil, errors.New("Error reading body from request")
	}

	var acContact AcContactListPayload
	err2 := json.Unmarshal(contact, &acContact)
	if err != nil {
		msg := fmt.Sprintf("Unable to unmarshal json into struct:\n%s\n", err2)
		fmt.Println(msg)
		return nil, errors.New(msg)
	}

	if len(acContact.Contacts) > 1 {
		fmt.Printf("WARN: Multiple contacts received for email address: %s\n", email)
	}

	//fmt.Printf("Well here's our struct:\n%+v\n", acContact)

	return &acContact, nil
}
