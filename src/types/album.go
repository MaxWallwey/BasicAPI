package types

import (
	"github.com/gocql/gocql"
)

// Album represents data about a record Album.
type Album struct {
	ID     gocql.UUID `json:"id"`
	Title  string     `json:"title"`
	Artist string     `json:"artist"`
	Price  float64    `json:"price"`
}
