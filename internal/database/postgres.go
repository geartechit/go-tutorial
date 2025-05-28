package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-tutorial/internal/config"
)

type PostgresDB struct {
	Pool          *pgxpool.Pool
	migrationPath string
}

func Connect(ctx context.Context, cfg config.PostgresConfig) (*PostgresDB, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Config()
		return nil, fmt.Errorf("failed to ping database pool: %w", err)
	}

	return &PostgresDB{
		Pool:          pool,
		migrationPath: cfg.MigrationPaths,
	}, nil
}
