package common_test

import (
	"fmt"
	"testing"

	"github.com/WMacalester/GraphKubeL/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestFormatPostgresConnString(t *testing.T) {
	name := "testuser"
	password := "testpassword"
	host := "hostname"
	port := "5432"
	dbname := "testdb"
	sslmode := "disable"

	t.Run("Should create expected conn string", func(t *testing.T) {
		expected := "postgres://testuser:testpassword@hostname:5432/testdb?sslmode=disable"

		result, err := common.FormatPostgresConnString(name, password, host, port, dbname, sslmode)
        
        assert.NoError(t, err)
        assert.Equal(t, expected, result)
	})

	tests := []struct {
		field string
		name string
		password string
		host string
		port string
		dbname string
		sslmode string
	}{
		{"name", "", password, host, port, dbname, sslmode},
		{"password", name, "", host, port, dbname, sslmode},
		{"host", name, password, "", port, dbname, sslmode},
		{"port", name, password, host, "", dbname, sslmode},
		{"dbname", name, password, host, port, "", sslmode},
		{"sslmode", name, password, host, port, dbname, ""},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(
			fmt.Sprintf("should return error if missing %s", tc.field), func(t *testing.T) {
				result, err := common.FormatPostgresConnString(tc.name, tc.password, tc.host, tc.port, tc.dbname, tc.sslmode)
        		assert.Error(t, err)
        		assert.Equal(t, "", result)
			},
		)
	}
}