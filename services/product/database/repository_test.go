package database

import (
	"context"
	"os"
	"testing"

	"github.com/WMacalester/GraphKubeL/services/product/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQueries struct {
    mock.Mock
}

func handleMockCall[T any](args mock.Arguments) (T, error){
    var zero T
	if args.Error(1) != nil {return zero, args.Error(1)}
	return args.Get(0).(T), nil
}

func (m *MockQueries) GetProducts(ctx context.Context) ([]Product, error) {
	return handleMockCall[[]Product](m.Called(ctx))
}   

func (m *MockQueries) GetProductCategoryById(ctx context.Context, id int32) (ProductCategory, error) {
    return handleMockCall[ProductCategory](m.Called(ctx))
}   

func (m *MockQueries) GetProductCategories(ctx context.Context) ([]ProductCategory, error) {
    return handleMockCall[[]ProductCategory](m.Called(ctx))
}   

func (m *MockQueries) InsertProductCategory(ctx context.Context, name string) (ProductCategory, error) {
    return handleMockCall[ProductCategory](m.Called(ctx))
}

func TestGetProducts(t *testing.T){
    mockQueries := new(MockQueries)
    repo := &ProductRepository{Queries: mockQueries}

    t.Run("Gets products", func(t *testing.T) {
		mockQueries.On("GetProducts", mock.Anything).Return([]Product{
			{ID: 0, Name: "Product1", Description: pgtype.Text{String: "Description1", Valid: true}},
			{ID: 1, Name: "Product2", Description: pgtype.Text{String: "Description2", Valid: true}},
		}, nil).Once()

		products, err := repo.GetProducts(context.Background())

		assert.NoError(t, err)
		assert.Len(t, products, 2)

        product1 := models.Product{Id: 0, Name: "Product1", Description: "Description1"}
        product2 := models.Product{Id: 1, Name: "Product2", Description: "Description2"}

        assert.Equal(t, []models.Product{product1, product2}, products)
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

func TestGetProductCategoryById(t *testing.T){
    mockQueries := new(MockQueries)
    repo := &ProductRepository{Queries: mockQueries}

    t.Run("Gets product category", func(t *testing.T) {
		expected := models.ProductCategory{Id: 1, Name: "Category 1"}

		mockQueries.On("GetProductCategoryById", mock.Anything).Return(
			ProductCategory{ID: 1, Name: "Category 1"},
			nil,
		).Once()

		result, err := repo.GetProductCategoryById(context.Background(), 1)

		assert.NoError(t, err)
        assert.Equal(t, expected, result)
		mockQueries.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockQueries.On("GetProductCategoryById", mock.Anything).Return(nil, assert.AnError).Once()

		result, err := repo.GetProductCategoryById(context.Background(), -1)

		assert.Error(t, err)
		assert.Equal(t, models.ProductCategory{}, result)

		mockQueries.AssertExpectations(t)
	})
}

func TestGetProductCategories(t *testing.T){
    mockQueries := new(MockQueries)
    repo := &ProductRepository{Queries: mockQueries}

    t.Run("Gets product categories", func(t *testing.T) {
		categoryDaos := []ProductCategory{
			{ID: 0, Name: "Category 1"},
			{ID: 1, Name: "Category 2"},
		}

		mockQueries.On("GetProductCategories", mock.Anything).Return(categoryDaos, nil).Once()

		products, err := repo.GetProductCategories(context.Background())

		assert.NoError(t, err)
		assert.Len(t, products, 2)

        category1 := models.ProductCategory{Id: 0, Name: "Category 1"}
        category2 := models.ProductCategory{Id: 1, Name: "Category 2"}

        assert.Equal(t, []models.ProductCategory{category1,category2}, products)
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

func TestInsertProductCategory(t *testing.T){
    mockQueries := new(MockQueries)
    repo := &ProductRepository{Queries: mockQueries}

    t.Run("Inserted product category returns id", func(t *testing.T) {
		name :=  "some-returned-name"
        mockQueries.On("InsertProductCategory", mock.Anything).Return(ProductCategory{ID: 12, Name:name}, nil).Once()
        expected := models.ProductCategory{Id: 12, Name: name}

        id, err := repo.InsertProductCategory(context.Background(), models.ProductCategory{Name: name})

        assert.NoError(t, err)
        assert.Equal(t, expected, id)
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