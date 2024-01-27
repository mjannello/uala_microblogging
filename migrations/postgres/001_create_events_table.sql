-- migrations/001_create_events_table.sql

CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    event_user_name VARCHAR(255) NOT NULL,
    event_type VARCHAR(255) NOT NULL,
    event_data VARCHAR(255) NOT NULL,
    event_date_created TIMESTAMPTZ DEFAULT current_timestamp
    );
