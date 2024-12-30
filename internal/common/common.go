package common

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func FormatPostgresConnString(user, password, host, port, dbname, sslmode string) (string, error) {
	if user == "" || password == "" || host == "" || port == "" || dbname == "" || sslmode == "" {
		return "", fmt.Errorf("one or more required environment variables are missing. user: %v, password: %v, host: %v, port: %v, dbname: %v, sslmode: %v", user, password, host, port, dbname, sslmode)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)

	return connString, nil
}

func ConnectToPostgresDb(ctx context.Context, createConnString func()(string, error)) (*pgxpool.Pool){
	connString, err := createConnString()
	if err != nil {
		log.Fatal(err)
	}

	connPool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	return connPool
}
