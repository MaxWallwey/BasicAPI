package users

import (
	"github.com/gocql/gocql"
	"time"
)

// User represents data about a record User.
type User struct {
	ID                   gocql.UUID `json:"id"`
	Name                 string     `json:"name"`
	EmailAddress         string     `json:"emailAddress"`
	LastUpdatedTimestamp time.Time  `json:"lastUpdatedTimestamp"`
}
