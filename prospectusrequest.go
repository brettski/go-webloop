package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/lytics/slackhook"
)

// application/x-www-form-urlencoded
// application/json

func prospectusRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entering prospectusRequest")
	if r.Method != http.MethodPost {
		fmt.Printf("Non-post request received: %s", r.Method)
		fmt.Fprint(w, "that's a get")
		return
	}
	contentType := strings.Split(r.Header.Get("Content-Type"), ";")[0]
	if "application/x-www-form-urlencoded" != contentType {
		fmt.Printf("content type %s is not supported. skipping", contentType)
		return
	}

	acFields, err := parseAcPostHook(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	firstname := acFields.FirstName
	lastname := acFields.LastName
	email := acFields.Email
	contactid := acFields.ContactId
	yearsSponsor := acFields.YearsSponsor
	// Query string value
	requestSource := r.FormValue("from")
	fmt.Printf("Prospectus requested by %s %s, %s, sponsored:%s, from:%s\n", firstname, lastname, email, yearsSponsor, requestSource)

	env, err1 := getEnvironmentInfo()
	if err1 != nil {
		log.Fatal(err1)
	}
	accountname := env.AcAccountName
	slackwebhookurl := env.SlackWebhookURL

	msg := slackhook.Message{
		Text:      "Prospectus requested!",
		IconEmoji: ":scroll:",
		IconURL:   "",
	}

	att := slackhook.Attachment{
		Fallback:  fmt.Sprintf("Prospectus requested by %s %s, %s", firstname, lastname, email),
		Color:     "4286f4",
		Title:     fmt.Sprintf("%s %s", firstname, lastname),
		TitleLink: fmt.Sprintf("https://%s.activehosted.com/app/contacts/%s", accountname, contactid),
		Text:      email,
	}

	fields := []slackhook.Field{
		slackhook.Field{
			Title: "Request Source",
			Value: requestSource,
		},
	}

	if len(yearsSponsor) > 0 {
		fields = append(fields, slackhook.Field{
			Title: "Years Sponsored",
			Value: yearsSponsor,
		})
	}
	att.Fields = fields

	msg.AddAttachment(&att)

	slack := slackhook.New(slackwebhookurl)
	slack.Send(&msg)

}
