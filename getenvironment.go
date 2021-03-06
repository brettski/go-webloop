package main

import (
	"errors"
	"log"
	"os"
	"strings"
)

// Environment value struct
type Environment struct {
	AcAccountName   string
	SlackWebhookURL string
	AcApiKey        string
	AirtableAcctId  string
	AirtableApiKey  string
	AcBaseUrl       string
	CounselorSlugs  []string
}

func getEnvironmentInfo() (*Environment, error) {

	accountname := os.Getenv("AC_ACCOUNT_NAME")
	slackwebhookurl, urlexists := os.LookupEnv("SLACK_WEBHOOK_URL")
	if !urlexists {
		log.Fatal("Slack webhook URL not provided. Stopping.")
		return nil, errors.New("Slack webhook URL not provided. Stopping")
	}
	acapikey := os.Getenv("AC_API_KEY")
	airtableacctid := os.Getenv("AIRTABLE_ACCOUNT_ID")
	airtableapikey := os.Getenv("AIRTABLE_API_KEY")
	slug := strings.Replace(os.Getenv("COUNSELOR_SLUGS"), " ", "", -1)
	counselorslugs := strings.Split(slug, ",")

	return &Environment{
		AcAccountName:   accountname,
		SlackWebhookURL: slackwebhookurl,
		AcApiKey:        acapikey,
		AirtableAcctId:  airtableacctid,
		AirtableApiKey:  airtableapikey,
		AcBaseUrl:       "https://%s.api-us1.com/api/3",
		CounselorSlugs:  counselorslugs,
	}, nil
}
