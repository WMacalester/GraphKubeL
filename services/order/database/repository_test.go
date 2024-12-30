package database

import (
	"context"
	"os"
	"testing"

	"github.com/WMacalester/GraphKubeL/internal/common"
	"github.com/WMacalester/GraphKubeL/services/order/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQueries struct {
	mock.Mock
}

func (m *MockQueries) GetOrders(ctx context.Context) ([]Order, error){
	return common.HandleMockCall[[]Order](m.Called(ctx))
}

func TestGetOrders(t *testing.T) {
	mockQueries := new(MockQueries)
	repo := &OrderRepository{Queries: mockQueries}

	id1 := 1
	id2 := 2
	transactionId1 := 3
	transactionId2 := 4
	productId1 := 5
	productId2 := 6
	numberOfItems1 := 10
	numberOfItems2 := 200

	mockQueries.On("GetOrders", mock.Anything).Return([]Order{
		{ID: int32(id1), TransactionID: int32(transactionId1), ProductID: int32(productId1), NumberOfItems: int32(numberOfItems1)},
		{ID: int32(id2), TransactionID: int32(transactionId2), ProductID: int32(productId2), NumberOfItems: int32(numberOfItems2)},
	}, nil).Once()

	orders, err := repo.GetOrders(context.Background())

	assert.NoError(t, err)
	assert.Len(t, orders, 2)

	assert.Equal(t, []models.Order{
		{Id: id1, TransactionID: transactionId1, ProductId: productId1, NumberOfItems: numberOfItems1},
		{Id: id2, TransactionID: transactionId2, ProductId: productId2, NumberOfItems: numberOfItems2},
	}, orders)
}

func TestCreateConnString(t *testing.T){
	setEnv := func() {
		os.Setenv("PGUSER", "testuser")
		os.Setenv("POSTGRES_PASSWORD", "testpassword")
		os.Setenv("ORDER_PG_HOST", "localhost")
		os.Setenv("ORDER_DB_PORT", "5432")
		os.Setenv("POSTGRES_DB", "testdb")
		os.Setenv("ORDER_PG_SSLMODE", "disable")
	}

	unsetEnv := func() {
		os.Unsetenv("PGUSER")
		os.Unsetenv("POSTGRES_PASSWORD")
		os.Unsetenv("ORDER_PG_HOST")
		os.Unsetenv("ORDER_DB_PORT")
		os.Unsetenv("POSTGRES_DB")
		os.Unsetenv("ORDER_PG_SSLMODE")
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
