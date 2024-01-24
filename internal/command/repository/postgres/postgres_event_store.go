package postgres

import (
	"database/sql"
	"uala/internal/command/event"
)

type PostgresEventStore struct {
	DB *sql.DB
}

func (pes *PostgresEventStore) SaveEvent(event event.Event) error {
	// TODO: fill with postgres logic
	return nil
}
