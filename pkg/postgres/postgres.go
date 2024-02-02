package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

func New(url string) (*pgxpool.Pool, error) {
	connAttempts := _defaultConnAttempts
	connTimeout := _defaultConnTimeout
	var pool *pgxpool.Pool

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("postgres: poolConfig error: %w", err)
	}

	for connAttempts > 0 {
		pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", connAttempts)

		time.Sleep(connTimeout)

		connAttempts--
	}
	if err != nil {
		return nil, fmt.Errorf("postgres: connAttempts == 0: %w", err)
	}
	return pool, nil
}
