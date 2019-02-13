package main

import (
	"net/http"
	"strings"
)

type acPostFields struct {
	firstName    string
	lastName     string
	email        string
	contactid    string
	yearsSponsor string
}

// Handles parsing of the common Active Campaign form values we use
func parseAcPostHook(r *http.Request) (*acPostFields, error) {
	err1 := r.ParseForm()
	if err1 != nil {
		return nil, err1
	}

	fields := acPostFields{}
	fields.firstName = r.FormValue("contact[first_name]")
	fields.lastName = r.FormValue("contact[last_name]")
	fields.email = r.FormValue("contact[email]")
	fields.contactid = r.FormValue("contact[id]")

	yearsSponsor := r.FormValue("contact[fields][yearssponsor]")
	yearsSponsor = strings.TrimPrefix(yearsSponsor, "||")
	yearsSponsor = strings.TrimSuffix(yearsSponsor, "||")
	fields.yearsSponsor = yearsSponsor

	return &fields, nil
}
