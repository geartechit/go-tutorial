#!/bin/bash

COMMAND=$1
NUM=$2

if [ -z "$NUM" ]; then
  migrate -path internal/database/sqlc/migrations/ \
          -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" \
          -verbose "$COMMAND"
else
  migrate -path internal/database/sqlc/migrations/ \
          -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" \
          -verbose "$COMMAND" "$NUM"
fi