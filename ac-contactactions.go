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
	log.Println("getAcContactByEmail")
	if !strings.Contains(email, "@") {
		log.Println("Provided email not valid")
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
		log.Println(msg)
		return nil, errors.New(msg)
	}
	req.Header.Add("Api_Token", acapikey)
	req.Header.Add("content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error requesting data from AC:\n%s\n", err)
		return nil, errors.New("Error requesting data from AC")
	}

	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("Response didn't return 200. Status received: %d", resp.StatusCode)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	contact, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body from request:\n%s\n", err)
		return nil, errors.New("Error reading body from request")
	}

	var acContact AcContactListPayload
	err2 := json.Unmarshal(contact, &acContact)
	if err != nil {
		msg := fmt.Sprintf("Unable to unmarshal json into struct:\n%s\n", err2)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	if len(acContact.Contacts) > 1 {
		log.Printf("WARN: Multiple contacts received for email address: %s\n", email)
	}

	//fmt.Printf("Well here's our struct:\n%+v\n", acContact)

	return &acContact, nil
}

func addTagToAcContact(contact AcContact, tagname string) bool {
	log.Println("addTagToAcContact")
	/*
		get tag id
		create tag if no exist
		add tag to contact
	*/
	actag, err := acLookupContactTag(tagname)
	if err != nil {
		log.Printf("Error getting tag info from AC:\n%s\n", err)
		return false
	}
	if len(actag.Id) == 0 {
		// Tag doesn't exist, create it
		actag, err = acAddContactTag(tagname)
		if err != nil {
			log.Printf("Error creating new tag:\n%s\n", err)
			return false
		}
	}
	// add tag to contact
	contactTag := fmt.Sprintf(`{"contactTag": {"contact":"%s", "tag":"%s"}}`, contact.Id, actag.Id)
	_, err = acPostRequest("/contactTags", contactTag)
	if err != nil {
		log.Printf("Error while adding tag to contact:\n%s\n", err)
		return false
	}

	return true
}
