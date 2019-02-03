package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lytics/slackhook"
)

func newContactHook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entering newContact")
	if r.Method != http.MethodPost {
		fmt.Println("not a post request, skipping")
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("First name received", r.Form["contact[first_name]"])
	firstname := r.FormValue("contact[first_name]")
	lastname := r.FormValue("contact[last_name]")
	email := r.FormValue("contact[email]")
	contactid := r.FormValue("contact[id]")

	accountname := os.Getenv("AC_ACCOUNT_NAME")
	slackwebhookurl, urlexists := os.LookupEnv("SLACK_WEBHOOK_URL")
	if !urlexists {
		log.Fatal("Slack webhook URL not provided. Stopping.")
	}

	msg := slackhook.Message{
		Text:      "New Prospectus Request!",
		IconEmoji: ":scroll:",
		IconURL:   "",
	}

	att := slackhook.Attachment{
		Fallback:   "New Prospectus Request!",
		Color:      "0000FF",
		AuthorName: "Active Campaign",
		Title:      fmt.Sprintf("%s %s", firstname, lastname),
		TitleLink:  fmt.Sprintf("https://%s.activehosted.com/app/contacts/%s", accountname, contactid),
		Text:       email,
	}

	/*
		fields := []slackhook.Field{
			slackhook.Field{
				Title: "Email",
				Value: email,
			},
		}
		att.Fields = fields
	*/

	msg.AddAttachment(&att)

	slack := slackhook.New(slackwebhookurl)
	slack.Send(&msg)
}
