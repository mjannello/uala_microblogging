-- migrations/001_create_events_table.sql

CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    event_data JSONB NOT NULL,
    event_date_created TIMESTAMPTZ DEFAULT current_timestamp
    );
