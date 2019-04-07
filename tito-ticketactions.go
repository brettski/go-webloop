package main

import "log"

func isTicketCounselor(slug string) bool {
	env, err := getEnvironmentInfo()
	if err != nil {
		log.Printf("Error getting environment variables:\n%s\n", err)
		return false
	}

	slugfound := false
	for _, v := range env.CounselorSlugs {
		if v == slug {
			slugfound = true
			break
		}
	}

	return slugfound
}
