#!/bin/bash

# Script para ejecutar las migraciones
# Asegúrate de tener las migraciones configuradas correctamente en ./migrations/postgres

migrate -path=/migrations/postgres -database=postgres://mjannello:uala_db_password@postgres:5432/uala_events_postgres?sslmode=disable up
