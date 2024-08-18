package main

import (
	"database/sql"
	"os"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	// Set up the test database connection
	var err error
	testDB, err = ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer testDB.Close()

	// Run the tests
	code := m.Run()

	// Clean up after tests
	_, _ = testDB.Exec("DELETE FROM messages")

	os.Exit(code)
}

func TestSaveAndLoadMessage(t *testing.T) {
	// Create a test message
	testMessage := &Message{
		Message: "Test message",
		Caller:  "Test Caller",
		City:    sql.NullString{String: "Test City", Valid: true},
		State:   sql.NullString{String: "TS", Valid: true},
		Zipcode: sql.NullString{String: "12345", Valid: true},
		Contact: sql.NullString{String: "test@example.com", Valid: true},
	}

	// Save the message
	err := testMessage.Save(testDB)
	if err != nil {
		t.Fatalf("Failed to save message: %v", err)
	}

	// Check that ID and UUID were set
	if testMessage.ID == 0 {
		t.Error("Expected ID to be set after save, but it's 0")
	}
	if testMessage.UUID == "" {
		t.Error("Expected UUID to be set after save, but it's empty")
	}

	// Load the message
	loadedMessage, err := LoadMessageByUUID(testDB, testMessage.UUID)
	if err != nil {
		t.Fatalf("Failed to load message: %v", err)
	}

	// Compare loaded message with original
	if loadedMessage.Message != testMessage.Message {
		t.Errorf("Loaded message doesn't match: got %v, want %v", loadedMessage.Message, testMessage.Message)
	}
	if loadedMessage.Caller != testMessage.Caller {
		t.Errorf("Loaded caller doesn't match: got %v, want %v", loadedMessage.Caller, testMessage.Caller)
	}
	if loadedMessage.City != testMessage.City {
		t.Errorf("Loaded city doesn't match: got %v, want %v", loadedMessage.City, testMessage.City)
	}
	if loadedMessage.State != testMessage.State {
		t.Errorf("Loaded state doesn't match: got %v, want %v", loadedMessage.State, testMessage.State)
	}
	if loadedMessage.Zipcode != testMessage.Zipcode {
		t.Errorf("Loaded zipcode doesn't match: got %v, want %v", loadedMessage.Zipcode, testMessage.Zipcode)
	}
	if loadedMessage.Contact != testMessage.Contact {
		t.Errorf("Loaded contact doesn't match: got %v, want %v", loadedMessage.Contact, testMessage.Contact)
	}

	// Check that Created and Updated times were set
	if loadedMessage.Created.IsZero() {
		t.Error("Expected Created time to be set, but it's zero")
	}
	if loadedMessage.Updated.IsZero() {
		t.Error("Expected Updated time to be set, but it's zero")
	}
}

func TestLoadNonExistentMessage(t *testing.T) {
	nonExistentUUID := uuid.New().String()
	_, err := LoadMessageByUUID(testDB, nonExistentUUID)
	if err != sql.ErrNoRows {
		t.Errorf("Expected sql.ErrNoRows for non-existent message, got: %v", err)
	}
}

func TestSaveMessageWithoutUUID(t *testing.T) {
	testMessage := &Message{
		Message: "Test message without UUID",
		Caller:  "Test Caller",
	}

	err := testMessage.Save(testDB)
	if err != nil {
		t.Fatalf("Failed to save message without UUID: %v", err)
	}

	if testMessage.UUID == "" {
		t.Error("Expected UUID to be generated and set after save, but it's empty")
	}
}
