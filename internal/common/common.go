package common

import "fmt"

func FormatPostgresConnString(user, password, host, port, dbname, sslmode string) (string, error) {
	if user == "" || password == "" || host == "" || port == "" || dbname == "" || sslmode == "" {
		return "", fmt.Errorf("one or more required environment variables are missing. user: %v, password: %v, host: %v, port: %v, dbname: %v, sslmode: %v", user, password, host, port, dbname, sslmode)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)

	return connString, nil
}
