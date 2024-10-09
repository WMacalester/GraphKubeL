package main_test

import (
	"os"
	"testing"

	"github.com/WMacalester/GraphKubeL/services/product"
	"github.com/stretchr/testify/assert"
)

func TestCreateConnString(t *testing.T){
	setEnv := func() {
		os.Setenv("PRODUCT_PG_USER", "testuser")
		os.Setenv("PRODUCT_PG_PASSWORD", "testpassword")
		os.Setenv("PRODUCT_PG_HOST", "localhost")
		os.Setenv("PRODUCT_PG_PORT", "5432")
		os.Setenv("PRODUCT_PG_DATABASE", "testdb")
		os.Setenv("PRODUCT_PG_SSLMODE", "disable")
	}

	unsetEnv := func() {
		os.Unsetenv("PRODUCT_PG_USER")
		os.Unsetenv("PRODUCT_PG_PASSWORD")
		os.Unsetenv("PRODUCT_PG_HOST")
		os.Unsetenv("PRODUCT_PG_PORT")
		os.Unsetenv("PRODUCT_PG_DATABASE")
		os.Unsetenv("PRODUCT_PG_SSLMODE")
	}

	t.Run("Should create expected conn string", func(t *testing.T) {
        setEnv()
		defer unsetEnv() 

		expected := "postgres://testuser:testpassword@localhost:5432/testdb?sslmode=disable"

		result, err := main.CreateConnString()
        
        assert.NoError(t, err)
        assert.Equal(t, expected, result)
	})

	t.Run("Should return error ", func(t *testing.T) {
		unsetEnv()

		_, err := main.CreateConnString()
        assert.Error(t, err)
	})
}