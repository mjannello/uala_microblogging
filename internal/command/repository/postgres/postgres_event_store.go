package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"uala/internal/command/event"
	"uala/internal/command/repository"
	"uala/pkg/logger"
)

const (
	InsertEventQuery = "INSERT INTO events (event_user_name, event_type, event_data, event_date_created) VALUES ($1, $2,$3, $4)"
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
	logger.Logger.Print(event.GetUserName())
	logger.Logger.Print(event.GetType())
	logger.Logger.Print(event.GetContent())
	logger.Logger.Print(event.GetDate())

	tx, err := pes.db.Beginx()
	if err != nil {
		return fmt.Errorf("could not begin transaction for saving event: %w", err)
	}
	_, err = tx.Exec(InsertEventQuery, event.GetUserName(), event.GetType(), event.GetContent(), event.GetDate())
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("could not save event: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction on saving event: %w", err)
	}

	return nil
}
