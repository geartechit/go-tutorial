.PHONY: generate
generate:
	@./scripts/sqlc.sh

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: migrate-up
migrate-up:
	@./scripts/migrate.sh up

.PHONY: migrate-down
migrate-down:
	@./scripts/migrate.sh down

.PHONY: migrate-reset
migrate-reset:
	migrate -path internal/database/sqlc/migrations/ -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" force 1

.PHONY: mock-gen
mock-gen:
	mockgen -source=internal/domain/employee/repository.go -destination=internal/mocks/mock_employee_repository.go -package=mocks
	mockgen -source=internal/services/employee_service.go -destination=internal/mocks/mock_employee_service.go -package=mocks
	mockgen -source=pkg/validator/validator.go -destination=internal/mocks/mock_validator.go -package=mocks
