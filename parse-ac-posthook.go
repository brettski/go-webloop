package main

import (
	"net/http"
	"strings"
)

type AcPostFields struct {
	FirstName    string
	LastName     string
	Email        string
	ContactId    string
	YearsSponsor string
	OrgName      string
}

// Handles parsing of the common Active Campaign form values we use
func parseAcPostHook(r *http.Request) (*AcPostFields, error) {
	err1 := r.ParseForm()
	if err1 != nil {
		return nil, err1
	}

	fields := AcPostFields{}
	fields.FirstName = r.FormValue("contact[first_name]")
	fields.LastName = r.FormValue("contact[last_name]")
	fields.Email = r.FormValue("contact[email]")
	fields.ContactId = r.FormValue("contact[id]")
	fields.OrgName = r.FormValue("contact[orgname]")

	yearsSponsor := r.FormValue("contact[fields][yearssponsor]")
	yearsSponsor = strings.TrimPrefix(yearsSponsor, "||")
	yearsSponsor = strings.TrimSuffix(yearsSponsor, "||")
	fields.YearsSponsor = yearsSponsor

	return &fields, nil
}
