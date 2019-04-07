package main

import (
	"encoding/json"
	"fmt"
)

// AcTag typical tag structure returned when requesting a tag
type AcTag struct {
	Tag         string `json:"tag"`
	Description string `json:"description"`
	TagType     string `json:"tagType"`
	Cdate       string `json:"cdate"`
	Id          string `json:"id"`
}

// AcNewTagPayload used in request payload to create new tag
type AcNewTagPayload struct {
	Tag acNewTag `json:"tag"`
}

type acNewTag struct {
	Tag         string `json:"tag"`
	TagType     string `json:"tagType"`
	Description string `json:"description"`
}

// AcTagList is a slice of AcTag structs
type AcTagList struct {
	Tags []AcTag `json:"tags"`
}

func acLookupContactTag(name string) (actag AcTag, err error) {
	fmt.Println("acLookupContactTag")
	endpoint := fmt.Sprintf("/tags?filters[tag]=%s", name)
	body, err := acGetRequest(endpoint)
	if err != nil {
		fmt.Printf("Error requesting tag:\n%s\n", err)
		return
	}

	var taglist AcTagList
	err1 := json.Unmarshal(body, &taglist)
	if err1 != nil {
		fmt.Printf("Error unmarshal to taglist struct:\n%s\n", err1)
		return actag, err1
	}

	if len(taglist.Tags) < 1 {
		fmt.Printf("No tags found\n")
		return
	}

	// find our tag
	for _, value := range taglist.Tags {
		if value.Tag == name && value.TagType == "contact" {
			actag = value
			break
		}
	}

	return
}

func acAddContactTag(name string) (actag AcTag, err error) {
	fmt.Println("acAddContactTag")
	newtagpayload := AcNewTagPayload{
		Tag: acNewTag{
			TagType:     "contact",
			Tag:         name,
			Description: "",
		},
	}
	payload, err := json.Marshal(newtagpayload)
	if err != nil {
		fmt.Printf("Error marshalling stuct to json:\n%s\n", err)
		return
	}
	fmt.Printf("New Tag Payload:\n%s\n", string(payload))

	body, err := acPostRequest("/tags", string(payload))
	if err != nil {
		fmt.Printf("Error while creating new contact tag:\n%s\n", err)
		return
	}

	err = json.Unmarshal(body, &actag)
	if err != nil {
		fmt.Printf("Error unmarshalling json to struct:\n%s\n", err)
		return
	}

	return
}

/*
 */
