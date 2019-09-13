package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord defines error to Snippet type
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials defines error to user credentials
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	//ErrDuplicateEmail defines error if a user tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet top level data type to database
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User define a new User type
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
