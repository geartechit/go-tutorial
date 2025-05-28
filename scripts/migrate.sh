#!/bin/bash

COMMAND=$1

migrate -path internal/database/sqlc/migrations/ -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose "$COMMAND"