package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"uala/internal/command/event"
	"uala/internal/command/repository"
)

const (
	InsertEventQuery = "INSERT INTO events (event_id, event_type, event_data, event_date_created) VALUES ($1, $2)"
)

type postgresEventStore struct {
	db *sqlx.DB
}

func NewPostgresEventStore(db *sqlx.DB) repository.EventStore {
	return &postgresEventStore{
		db: db,
	}
}

func (pes *postgresEventStore) SaveEvent(event event.Event) error {

	tx, err := pes.db.Beginx()
	if err != nil {
		return fmt.Errorf("could not begin transaction for saving event %d: %w", event.ID, err)
	}
	_, err = tx.Exec(InsertEventQuery, event.ID, event.Type, event.Content, event.DateCreated)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("could not save event %d: %w", event.ID, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction on saving event %d: %w", event.ID, err)
	}

	return nil
}
