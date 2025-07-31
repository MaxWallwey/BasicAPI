package users

import (
	"github.com/gocql/gocql"
	"time"
)

// User represents data about a record User.
type User struct {
	ID                   gocql.UUID `cql:"id" json:"id"`
	Name                 string     `cql:"name" json:"name"`
	EmailAddress         string     `cql:"emailAddress" json:"emailAddress"`
	LastUpdatedTimestamp time.Time  `cql:"lastUpdatedTimestamp" json:"lastUpdatedTimestamp"`
}
