package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Standard web requests for Active Campaign

// acGetRequest: standard get requests to AC
// endpoint: endpoing to request after `api/3` INCLUDING any query string params
func acGetRequest(endpoint string) (body []byte, err error) {
	env, err := getEnvironmentInfo()
	if err != nil {
		log.Fatal(err)
	}

	accountname := env.AcAccountName
	acbaseurl := env.AcBaseUrl
	acapikey := env.AcApiKey

	// "https://%s.api.-us1.com/api/3"
	url := fmt.Sprintf(acbaseurl+endpoint, accountname)
	//log.Printf("Full GET req url:\n%s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		msg := fmt.Sprintf("Error setting up request:\n%s\n", err)
		log.Println(msg)
		return
	}
	req.Header.Add("Api_Token", acapikey)
	req.Header.Add("content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error requesting data from AC:\n%s\n", err)
		return nil, errors.New("Error requesting data from AC")
	}

	if resp.StatusCode < 200 || resp.StatusCode > 399 {
		msg := fmt.Sprintf("Response didn't return 2xx-3xx. Status received: %d", resp.StatusCode)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	contact, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("Error reading body from request:\n%s\n", err)
		return nil, errors.New("Error reading body from request")
	}

	return contact, nil
}

func acPostRequest(endpoint string, payload string) (body []byte, err error) {
	env, err := getEnvironmentInfo()
	if err != nil {
		log.Fatal(err)
	}

	accountname := env.AcAccountName
	acbaseurl := env.AcBaseUrl
	acapikey := env.AcApiKey

	// "https://%s.api.-us1.com/api/3"
	url := fmt.Sprintf(acbaseurl+endpoint, accountname)
	//log.Printf("Full POST req url:\n%s\n", url)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		msg := fmt.Sprintf("Error setting up request:\n%s\n", err)
		log.Println(msg)
		return
	}
	req.Header.Add("Api_Token", acapikey)
	req.Header.Add("content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error requesting data from AC:\n%s\n", err)
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
