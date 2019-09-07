package models

import (
	"errors"
	"time"
)

// ErrNoRecord defines error to Snippet type
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet top level data type to database
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
