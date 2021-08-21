package database

import "time"

// New returns a new database instance
func New(name string) *Database {
	return &Database{
		Name:        name,
		DumpedBytes: nil,
		AddedAt:     time.Now().UTC(),
	}
}
