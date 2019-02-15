package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lytics/slackhook"
)

func newSponsorOnBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entering newSponsorOnBoard")
	if r.Method != http.MethodPost {
		fmt.Printf("Non-post request received: %s", r.Method)
		return
	}

	fields, err := parseAcPostHook(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(fields)

	env, err1 := getEnvironmentInfo()
	if err1 != nil {
		log.Fatal(err1)
	}

	msg := slackhook.Message{
		Text:      "We have a new client!",
		IconEmoji: ":moneybag:",
		IconURL:   "",
	}

	attach := slackhook.Attachment{
		Fallback:  fmt.Sprintf("New Sponsor %s represented by %s %s has paid!", fields.OrgName, fields.FirstName, fields.LastName),
		Color:     "85bb65",
		Title:     fmt.Sprintf("%s -- %s %s", fields.OrgName, fields.FirstName, fields.LastName),
		TitleLink: fmt.Sprintf("https://%s.activehosted.com/app/contacts/%s", env.AcAccountName, fields.ContactId),
	}

	msg.AddAttachment(&attach)
	fmt.Println(msg)
	slack := slackhook.New(env.SlackWebhookURL)
	slack.Send(&msg)
}
