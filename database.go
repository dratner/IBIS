package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Message struct {
	ID      int
	UUID    string
	Message string
	City    sql.NullString
	State   sql.NullString
	Zipcode sql.NullString
	Caller  string
	Contact sql.NullString
	Created time.Time
	Updated time.Time
	Deleted sql.NullTime
}

func ConnectDatabase() (*sql.DB, error) {
	var dbURL string

	// Check for IBIS_DATABASE_URL first
	if dbURL = os.Getenv("IBIS_DATABASE_URL"); dbURL == "" {
		// If not found, check for DATABASE_URL
		if dbURL = os.Getenv("DATABASE_URL"); dbURL == "" {
			return nil, fmt.Errorf("no database URL found in environment variables")
		}
	}

	// Open the database connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		db.Close() // Make sure to close the connection if Ping fails
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
}

func LoadMessageByUUID(db *sql.DB, uuid string) (*Message, error) {
	query := `
		SELECT id, uuid, message, city, state, zipcode, caller, contact, created, updated, deleted
		FROM messages
		WHERE uuid = $1
	`

	var m Message
	err := db.QueryRow(query, uuid).Scan(
		&m.ID, &m.UUID, &m.Message, &m.City, &m.State, &m.Zipcode,
		&m.Caller, &m.Contact, &m.Created, &m.Updated, &m.Deleted,
	)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func LoadAllMessages(db *sql.DB) ([]Message, error) {
	query := `
		SELECT id, uuid, message, city, state, zipcode, caller, contact, created, updated, deleted
		FROM messages
		WHERE deleted IS NULL
		ORDER BY created DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var m Message
		err := rows.Scan(
			&m.ID, &m.UUID, &m.Message, &m.City, &m.State, &m.Zipcode,
			&m.Caller, &m.Contact, &m.Created, &m.Updated, &m.Deleted,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (m *Message) Save(db *sql.DB) error {
	if m.UUID == "" {
		m.UUID = uuid.New().String()
	}

	query := `
		INSERT INTO messages (uuid, message, city, state, zipcode, caller, contact, created, updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	now := time.Now()
	err := db.QueryRow(query,
		m.UUID, m.Message, m.City, m.State, m.Zipcode,
		m.Caller, m.Contact, now, now,
	).Scan(&m.ID)

	if err != nil {
		return err
	}

	m.Created = now
	m.Updated = now
	return nil
}
