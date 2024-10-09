package main

import (
	"context"
	"os"
	"testing"

	"github.com/WMacalester/GraphKubeL/services/product/database"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQueries struct {
    mock.Mock
}

func (m *MockQueries) GetProducts(ctx context.Context) ([]database.Product, error) {
    args := m.Called(ctx)
    if args.Error(1) != nil {return nil, args.Error(1)}
	return args.Get(0).([]database.Product), nil
}

func TestGetProducts(t *testing.T){
    mockQueries := new(MockQueries)
    repo := &ProductRepository{Queries: mockQueries}

    t.Run("Gets products", func(t *testing.T) {
		mockQueries.On("GetProducts", mock.Anything).Return([]database.Product{
			{ID: 0, Name: "Product1", Description: pgtype.Text{String: "Description1", Valid: true}},
			{ID: 1, Name: "Product2", Description: pgtype.Text{String: "Description2", Valid: true}},
		}, nil).Once()

		products, err := repo.GetProducts(context.Background())

		assert.NoError(t, err)
		assert.Len(t, products, 2)

        product1 := Product{Id: 0, Name: "Product1", Description: "Description1"}
        product2 := Product{Id: 1, Name: "Product2", Description: "Description2"}

        assert.Equal(t, []Product{product1, product2}, products)
		mockQueries.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockQueries.On("GetProducts", mock.Anything).Return(nil, assert.AnError).Once()

		products, err := repo.GetProducts(context.Background())

		assert.Error(t, err)
		assert.Nil(t, products)

		mockQueries.AssertExpectations(t)
	})
}


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

		result, err := CreateConnString()
        
        assert.NoError(t, err)
        assert.Equal(t, expected, result)
	})

	t.Run("Should return error ", func(t *testing.T) {
		unsetEnv()

		_, err := CreateConnString()
        assert.Error(t, err)
        
    })
}