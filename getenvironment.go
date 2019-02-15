package main

import (
	"errors"
	"log"
	"os"
)

type Environment struct {
	AcAccountName   string
	SlackWebhookURL string
}

func getEnvironmentInfo() (*Environment, error) {

	accountname := os.Getenv("AC_ACCOUNT_NAME")
	slackwebhookurl, urlexists := os.LookupEnv("SLACK_WEBHOOK_URL")
	if !urlexists {
		log.Fatal("Slack webhook URL not provided. Stopping.")
		return nil, errors.New("Slack webhook URL not provided. Stopping")
	}

	return &Environment{
		AcAccountName:   accountname,
		SlackWebhookURL: slackwebhookurl,
	}, nil
}
