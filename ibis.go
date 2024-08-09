package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port    = ":8080"
	version = "0.1.0"
)

var (
	twilioAccountSid  = "YOUR_TWILIO_ACCOUNT_SID"
	twilioAuthToken   = "YOUR_TWILIO_AUTH_TOKEN"
	twilioPhoneNumber = "YOUR_TWILIO_PHONE_NUMBER"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("IBIS - Injured Bird Information System v" + version))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<html><head><title>IBIS Health</title><body><h3>Health check</h3>"))
	if twilioAccountSid == "" {
		w.Write([]byte("<p>Twilio Account Sid not found</p>"))
	} else {
		w.Write([]byte("<p>Twilio Account Sid OK</p>"))
	}
	if twilioAuthToken == "" {
		w.Write([]byte("<p>Twilio Auth Token not found</p>"))
	} else {
		w.Write([]byte("<p>Twilio Auth Token OK</p>"))
	}
	if twilioPhoneNumber == "" {
		w.Write([]byte("<p>Twilio Phone Number not found</p>"))
	} else {
		w.Write([]byte("<p>Twilio Phone Number OK</p>"))
	}
	w.Write([]byte("</body></html>"))
}

func main() {
	fmt.Println("IBIS")
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/sms", handleSMS)
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
