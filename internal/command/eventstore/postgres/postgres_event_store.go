package postgres

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"uala/internal/command/eventstore"
	"uala/internal/model"
)

const (
	InsertEventQuery = "INSERT INTO events (event_data, event_date_created) VALUES ($1, $2) RETURNING id"
)

type postgresEventStore struct {
	db *sqlx.DB
}

func NewPostgresEventStore(db *sqlx.DB) eventstore.EventStore {
	return &postgresEventStore{
		db: db,
	}
}

func (pes *postgresEventStore) SaveEvent(event model.Event) (int64, error) {
	tx, err := pes.db.Beginx()
	if err != nil {
		return 0, fmt.Errorf("could not begin transaction for saving model: %w", err)
	}
	eventData, err := json.Marshal(event.GetData())
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("could not marshal event data: %w", err)
	}
	var lastInsertID int64
	err = tx.QueryRow(InsertEventQuery, eventData, event.GetDate()).Scan(&lastInsertID)
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("could not save model: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("failed to commit transaction on saving model: %w", err)
	}

	return lastInsertID, nil
}
