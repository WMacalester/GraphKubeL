package database

import (
	"context"
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
