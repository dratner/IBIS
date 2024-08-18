package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	TwilioAccountSID   string
	TwilioAccountToken string
	TwilioPhoneNumber  string
	Port               string
}

var (
	Conf Config
	DB   *sql.DB
)

const (
	Version = "0.1.1"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("IBIS - Injured Bird Information System v" + Version))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<html><head><title>IBIS Health</title><body><h3>Health check</h3>"))
	if Conf.TwilioAccountSID == "" {
		w.Write([]byte("<p>Twilio Account Sid not found</p>"))
	} else {
		w.Write([]byte("<p>Twilio Account Sid OK</p>"))
	}
	if Conf.TwilioAccountToken == "" {
		w.Write([]byte("<p>Twilio Auth Token not found</p>"))
	} else {
		w.Write([]byte("<p>Twilio Auth Token OK</p>"))
	}
	if Conf.TwilioPhoneNumber == "" {
		w.Write([]byte("<p>Twilio Phone Number not found</p>"))
	} else {
		w.Write([]byte("<p>Twilio Phone Number OK</p>"))
	}
	w.Write([]byte("</body></html>"))
}

func main() {
	fmt.Println("IBIS")

	Conf.TwilioAccountSID = os.Getenv("TWILIO_ACCOUNT_SID")
	Conf.TwilioAccountToken = os.Getenv("TWILIO_ACCOUNT_TOKEN")
	Conf.TwilioPhoneNumber = os.Getenv("TWILIO_ACCOUNT_NUMBER")
	Conf.Port = ":8080"

	DB, err := ConnectDatabase()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer DB.Close()

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/sms", handleSMS)

	log.Printf("Server is running on port %s", Conf.Port)
	log.Printf("Phone: %s", Conf.TwilioPhoneNumber)
	log.Fatal(http.ListenAndServe(Conf.Port, nil))
}
