package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func atCreateRecord(name string, email string, issue string) bool {
	log.Println("atCreateRecord")
	payload := fmt.Sprintf(`
	{"fields": {
		"Name": "%s",
		"Email": "%s",
		"Issue": "%s"
	}}`, name, email, issue)

	_, err := atPostRequest("/Issues", payload)
	if err != nil {
		log.Printf("Error writing new record into Airtable. payload: %s\n", payload)
		return false
	}

	return true
}

func atPostRequest(endpoint string, payload string) (body []byte, err error) {
	log.Println("atPostRequest")
	env, err := getEnvironmentInfo()
	if err != nil {
		log.Fatal(err)
	}

	atapikey := env.AirtableApiKey
	atacctid := env.AirtableAcctId

	// "https://%s.api.-us1.com/api/3"
	url := fmt.Sprintf("https://api.airtable.com/v0/%s%s", atacctid, endpoint)
	//log.Printf("Full POST req url:\n%s\n", url)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		msg := fmt.Sprintf("Error setting up request:\n%s\n", err)
		log.Println(msg)
		return
	}
	req.Header.Add("Authorization", "Bearer "+atapikey)
	req.Header.Add("content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error requesting data from AC:\n%s\n", err)
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 399 {
		msg := fmt.Sprintf("Response didn't return 2xx-3xx. Status received: %d", resp.StatusCode)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("Error reading body from request\n%s\n", err)
		return
	}

	return

}
