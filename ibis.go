package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

type Config struct {
	TwilioAccountSID   string
	TwilioAccountToken string
	TwilioPhoneNumber  string
	Port               string
}

var (
	Conf Config
)

const (
	Version = "0.1.2"
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

func handleStatic(w http.ResponseWriter, r *http.Request) {
	assetName := r.URL.Path[len("/static/"):]
	extension := filepath.Ext(assetName)
	extMime := map[string]string{
		".css":  "text/css",
		".png":  "image/png",
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".gif":  "image/gif",
		".svg":  "image/svg+xml",
		".ico":  "image/x-icon",
		".html": "text/html",
		".js":   "application/javascript",
	}

	log.Printf("asset %s extension %s mime %s", assetName, extension, extMime[extension])

	mime, ok := extMime[extension]
	if !ok {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	assetPath := filepath.Join(os.Getenv("IBIS_DIR"), "static", filepath.Clean(assetName))
	if _, err := os.Stat(assetPath); os.IsNotExist(err) {
		if os.Getenv("IBIS_DIR") == "" {
			log.Println("IBIS_DIR environment variable not set")
		}
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", mime)
	http.ServeFile(w, r, assetPath)
}

func handleOptIn(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/optin.tpl")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	fmt.Println("IBIS")

	Conf.TwilioAccountSID = os.Getenv("TWILIO_ACCOUNT_SID")
	Conf.TwilioAccountToken = os.Getenv("TWILIO_ACCOUNT_TOKEN")
	Conf.TwilioPhoneNumber = os.Getenv("TWILIO_ACCOUNT_NUMBER")
	Conf.Port = ":8080"

	db, err := ConnectDatabase()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/optin", handleOptIn)
	http.HandleFunc("/static/", handleStatic)
	http.HandleFunc("/sms", smsHandler(db))
	http.HandleFunc("/messages", messagesHandler(db))

	log.Printf("Server is running on port %s", Conf.Port)
	log.Fatal(http.ListenAndServe(Conf.Port, nil))
}
