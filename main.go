package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n\n\n", r.Header)

	webhookData := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&webhookData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("got webhook payload: ")
	for k, v := range webhookData {
		fmt.Printf("%s : %v\n", k, v)
	}
}

func testPayload(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n\n\n", r.Header)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("r.Form", r.Form)
	fmt.Println("r.PostForm", r.PostForm)
	fmt.Printf("\n----\n")
	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}

}

func main() {
	defaultport := "8088"
	sport, pexist := os.LookupEnv("PORT")
	if !pexist {
		sport = defaultport
	}
	http.HandleFunc("/webhook", handleWebhook)
	http.HandleFunc("/testpayload", testPayload)
	http.HandleFunc("/newcontact", newContactHook)
	http.HandleFunc("/prospectusrequest", prospectusRequest)
	http.HandleFunc("/newsponsoronboard", newSponsorOnBoard)
	http.HandleFunc("/trackcounselorticket", trackCounelorTicket)
	log.Println("server started on port " + sport)
	log.Fatal(http.ListenAndServe(":"+sport, nil))
}
