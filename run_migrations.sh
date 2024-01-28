#!/bin/bash

migrate -path=/migrations/postgres -database=postgres://mjannello:uala_db_password@postgres:5432/uala_events_postgres?sslmode=disable up
