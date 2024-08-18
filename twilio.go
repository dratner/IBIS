package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwiMLResponse struct {
	XMLName xml.Name `xml:"Response"`
	Message string   `xml:"Message"`
}

func smsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Printf("Error parsing form: %v", err)
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		from := r.FormValue("From")
		body := r.FormValue("Body")

		log.Printf("Received message from %s: %s", from, body)

		// Send thank you message using Twilio API
		go SMSRouter(from, body, db)

		// Respond to Twilio with TwiML
		twiml := TwiMLResponse{Message: "Thank you for your message!"}
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(twiml)
	}
}

func SendMessage(to string, body string) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: Conf.TwilioAccountSID,
		Password: Conf.TwilioAccountToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(Conf.TwilioPhoneNumber)
	params.SetBody(body)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS: " + err.Error())
	} else {
		fmt.Println("Thank you message sent successfully!")
	}
}
