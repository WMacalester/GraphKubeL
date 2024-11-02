//go:build tools

// Package tools tracks dependencies on tools that are required during development
// but are not imported in the actual code.
package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/redis/go-redis/v9"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)