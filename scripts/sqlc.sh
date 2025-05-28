#!/bin/bash

export POSTGRES_CONN_STRING="postgres://postgres:password@localhost:5432/postgres?sslmode=disable"

sqlc generate -f ./internal/database/sqlc/sqlc.yaml
echo "generated sqlc for go-tutorial"
