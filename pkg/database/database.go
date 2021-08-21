package database

import "time"

func New(name string) *Database {
	return &Database{
		Name:        name,
		DumpedBytes: nil,
		AddedAt:     time.Now().UTC(),
	}
}
