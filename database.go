package main

import (
	"database/sql"
	"encoding/json"
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

type Person struct {
	ID          int64           `db:"id" json:"id"`
	UUID        uuid.UUID       `db:"uuid" json:"uuid"`
	Name        string          `db:"name" json:"name"`
	Email       sql.NullString  `db:"email" json:"email,omitempty"`
	Password    sql.NullString  `db:"password" json:"-"`
	Phone       string          `db:"phone" json:"phone"`
	Preferences json.RawMessage `db:"preferences" json:"preferences,omitempty"`
	OnDuty      sql.NullBool    `db:"onduty" json:"onduty,omitempty"`
	City        string          `db:"city" json:"city"`
	State       string          `db:"state" json:"state"`
	Zipcode     sql.NullString  `db:"zipcode" json:"zipcode,omitempty"`
	Created     time.Time       `db:"created" json:"created"`
	Updated     time.Time       `db:"updated" json:"updated"`
	Deleted     sql.NullTime    `db:"deleted" json:"deleted,omitempty"`
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

func (p *Person) Save(db *sql.DB) error {
	if p.UUID == uuid.Nil {
		p.UUID = uuid.New()
	}

	query := `
		INSERT INTO people (uuid, name, email, password, phone, preferences, onduty, city, state, zipcode, created, updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`

	now := time.Now()
	err := db.QueryRow(query,
		p.UUID, p.Name, p.Email, p.Password, p.Phone, p.Preferences, p.OnDuty,
		p.City, p.State, p.Zipcode, now, now,
	).Scan(&p.ID)

	if err != nil {
		return fmt.Errorf("error saving person: %v", err)
	}

	p.Created = now
	p.Updated = now
	return nil
}

func LoadPersonByUUID(db *sql.DB, uuidStr string) (*Person, error) {
	query := `
		SELECT id, uuid, name, email, password, phone, preferences, onduty, city, state, zipcode, created, updated, deleted
		FROM people
		WHERE uuid = $1
	`

	var p Person

	err := db.QueryRow(query, uuidStr).Scan(
		&p.ID, &uuidStr, &p.Name, &p.Email, &p.Password, &p.Phone, &p.Preferences, &p.OnDuty,
		&p.City, &p.State, &p.Zipcode, &p.Created, &p.Updated, &p.Deleted,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no person found with UUID %s", uuidStr)
		}
		return nil, fmt.Errorf("error loading person: %v", err)
	}

	p.UUID, err = uuid.Parse(uuidStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing UUID: %v", err)
	}

	return &p, nil
}

func LoadAllPeople(db *sql.DB) ([]Person, error) {
	query := `
		SELECT id, uuid, name, email, password, phone, preferences, onduty, city, state, zipcode, created, updated, deleted
		FROM people
		WHERE deleted IS NULL
		ORDER BY created DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying people: %v", err)
	}
	defer rows.Close()

	var people []Person
	for rows.Next() {
		var p Person
		var uuidStr string

		err := rows.Scan(
			&p.ID, &uuidStr, &p.Name, &p.Email, &p.Password, &p.Phone, &p.Preferences, &p.OnDuty,
			&p.City, &p.State, &p.Zipcode, &p.Created, &p.Updated, &p.Deleted,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning person row: %v", err)
		}

		p.UUID, err = uuid.Parse(uuidStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing UUID: %v", err)
		}

		people = append(people, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating people rows: %v", err)
	}

	return people, nil
}

func DeletePerson(db *sql.DB, uuidStr string) error {
	query := `
		UPDATE people
		SET deleted = $1, updated = $1
		WHERE uuid = $2 AND deleted IS NULL
	`

	now := time.Now()
	result, err := db.Exec(query, now, uuidStr)
	if err != nil {
		return fmt.Errorf("error executing soft delete: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no person found with UUID %s or already deleted", uuidStr)
	}

	return nil
}
