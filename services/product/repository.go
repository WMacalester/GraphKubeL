package main

import (
	"fmt"
	"os"
)

func CreateConnString() (string, error){
	user := os.Getenv("PRODUCT_PG_USER")
	password := os.Getenv("PRODUCT_PG_PASSWORD")
	host := os.Getenv("PRODUCT_PG_HOST")
	port := os.Getenv("PRODUCT_PG_PORT")
	dbname := os.Getenv("PRODUCT_PG_DATABASE")
	sslmode := os.Getenv("PRODUCT_PG_SSLMODE")

	if user == "" || password == "" || host == "" || port == "" || dbname == "" {
		return "", fmt.Errorf("one or more required environment variables are missing")
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)

	return connString, nil
}
