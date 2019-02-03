package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for key, value := range r.Form {
		fmt.Printf("%s :: %s\n", key, value)
	}
	firstname := r.FormValue("contact[first_name]")
	lastname := r.FormValue("contact[last_name]")
	email := r.FormValue("contact[email]")
	contactid := r.FormValue("contact[id]")
	yearsSponsor := r.FormValue("contact[fields][yearssponsor]")
	yearsSponsor = strings.TrimPrefix(yearsSponsor, "||")
	yearsSponsor = strings.TrimSuffix(yearsSponsor, "||")
	//ysVals := strings.Split(yearsSponsor, "||")
	fmt.Printf("Prospectus requested by %s %s, %s, sponsored:%s\n", firstname, lastname, email, yearsSponsor)

	accountname := os.Getenv("AC_ACCOUNT_NAME")
	slackwebhookurl, urlexists := os.LookupEnv("SLACK_WEBHOOK_URL")
	if !urlexists {
		log.Fatal("Slack webhook URL not provided. Stopping.")
	}

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
	if len(yearsSponsor) > 0 {
		fields := []slackhook.Field{
			slackhook.Field{
				Title: "Years Sponsored",
				Value: yearsSponsor,
			},
		}
		att.Fields = fields
	}
	msg.AddAttachment(&att)
	slack := slackhook.New(slackwebhookurl)
	slack.Send(&msg)

}
