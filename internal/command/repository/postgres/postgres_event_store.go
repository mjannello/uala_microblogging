package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"uala/internal/command/event"
	"uala/internal/command/repository"
)

const (
	InsertEventQuery = "INSERT INTO events (event_user_name, event_type, event_data, event_date_created) VALUES ($1, $2,$3, $4) RETURNING id"
)

type postgresEventStore struct {
	db *sqlx.DB
}

func NewPostgresEventStore(db *sqlx.DB) repository.EventStore {
	return &postgresEventStore{
		db: db,
	}
}

func (pes *postgresEventStore) SaveEvent(event event.Event) (int64, error) {

	tx, err := pes.db.Beginx()
	if err != nil {
		return 0, fmt.Errorf("could not begin transaction for saving event: %w", err)
	}
	var lastInsertID int64
	err = tx.QueryRow(InsertEventQuery, event.GetUserName(), event.GetType(), event.GetContent(), event.GetDate()).Scan(&lastInsertID)
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("could not save event: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("failed to commit transaction on saving event: %w", err)
	}

	return lastInsertID, nil
}
