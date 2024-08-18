package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

func messagesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		messages, err := LoadAllMessages(db)
		if err != nil {
			http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
			log.Printf("Error retrieving messages: %v", err)
			return
		}

		tmpl, err := template.New("messages").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Messages</title>
    <style>
        table {
            border-collapse: collapse;
            width: 100%;
        }
        th, td {
            border: 1px solid black;
            padding: 8px;
            text-align: left;
        }
        th {
            background-color: #f2f2f2;
        }
    </style>
</head>
<body>
    <h1>Messages</h1>
    <table>
        <tr>
            <th>ID</th>
            <th>UUID</th>
            <th>Message</th>
            <th>City</th>
            <th>State</th>
            <th>Zipcode</th>
            <th>Caller</th>
            <th>Contact</th>
            <th>Created</th>
            <th>Updated</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.ID}}</td>
            <td>{{.UUID}}</td>
            <td>{{.Message}}</td>
            <td>{{if .City.Valid}}{{.City.String}}{{end}}</td>
            <td>{{if .State.Valid}}{{.State.String}}{{end}}</td>
            <td>{{if .Zipcode.Valid}}{{.Zipcode.String}}{{end}}</td>
            <td>{{.Caller}}</td>
            <td>{{if .Contact.Valid}}{{.Contact.String}}{{end}}</td>
            <td>{{.Created.Format "2006-01-02 15:04:05"}}</td>
            <td>{{.Updated.Format "2006-01-02 15:04:05"}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>
`)
		if err != nil {
			http.Error(w, "Failed to parse template", http.StatusInternalServerError)
			log.Printf("Error parsing template: %v", err)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		err = tmpl.Execute(w, messages)
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			log.Printf("Error rendering template: %v", err)
			return
		}
	}
}
