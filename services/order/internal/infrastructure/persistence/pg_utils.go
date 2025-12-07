package persistence

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, user, passwd, host, port, name string) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		passwd,
		host,
		port,
		name,
	)

	poolCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig: %w", err)
	}

	return pool, nil
}
