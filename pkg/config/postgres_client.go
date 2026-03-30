package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func InitializePostgresConnection() {
	dsn := GetDatabaseURL()

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to PostgreSQL: %v\n", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to ping PostgreSQL: %v\n", err)
	}

	dbPool = pool
	log.Printf("PostgreSQL successfully connected")
}

func GetPool() *pgxpool.Pool {
	return dbPool
}

func DisconnectPostgresClient() {
	if dbPool != nil {
		dbPool.Close()
		dbPool = nil
	}
}
