package tdtg

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
)

type DB struct {
	db *badger.DB
}

func NewDB(dbPath string) (*DB, error) {
	// Create the BadgerDB connection
	opts := badger.DefaultOptions(dbPath)
	if dbPath == "" {
		opts.InMemory = true
	}
	b, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("unable to open badger db: %w", err)
	}

	// Return the DB
	return &DB{
		db: b,
	}, nil
}
