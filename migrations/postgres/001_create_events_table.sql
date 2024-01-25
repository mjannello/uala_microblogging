CREATE TABLE IF NOT EXISTS events (
    event_id SERIAL PRIMARY KEY,
    event_type VARCHAR(255) NOT NULL,
    event_data JSONB,
    event_date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);