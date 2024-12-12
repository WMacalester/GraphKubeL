package database

import (
	"context"
	"os"
	"testing"

	"github.com/WMacalester/GraphKubeL/internal/common"
	"github.com/WMacalester/GraphKubeL/services/product/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQueries struct {
    mock.Mock
}

func (m *MockQueries) GetProductById(ctx context.Context, id int32) (GetProductByIdRow, error) {
	return common.HandleMockCall[GetProductByIdRow](m.Called(ctx))
}

func (m *MockQueries) GetProducts(ctx context.Context) ([]GetProductsRow, error) {
	return common.HandleMockCall[[]GetProductsRow](m.Called(ctx))
}   

func (m *MockQueries) GetProductCategoryById(ctx context.Context, id int32) (ProductCategory, error) {
    return common.HandleMockCall[ProductCategory](m.Called(ctx))
}   

func (m *MockQueries) GetProductCategories(ctx context.Context) ([]ProductCategory, error) {
    return common.HandleMockCall[[]ProductCategory](m.Called(ctx))
}   

func (m *MockQueries) InsertProduct(ctx context.Context, productCreateDao InsertProductParams) (Product, error) {
    return common.HandleMockCall[Product](m.Called(ctx))
}

func (m *MockQueries) InsertProductCategory(ctx context.Context, name string) (ProductCategory, error) {
    return common.HandleMockCall[ProductCategory](m.Called(ctx))
}

func TestGetProducts(t *testing.T){
    mockQueries := new(MockQueries)
    repo := &ProductRepository{Queries: mockQueries}

    t.Run("Gets products", func(t *testing.T) {
		productId1 := 0
		productId2 := 1
		productName1 := "Product 1"
		productName2 := "Product 2"
		categoryId1 := 10
		categoryId2 := 11
		categoryName1 := "category 1"
		categoryName2 := "category 2"
		description1 := "description 1"
		description2 := "description 2"


		mockQueries.On("GetProducts", mock.Anything).Return([]GetProductsRow{
			{ID: int32(productId1), Name: productName1, CategoryID: int32(categoryId1), CategoryName: categoryName1, Description: pgtype.Text{String: description1, Valid: true}},
			{ID: int32(productId2), Name: productName2, CategoryID: int32(categoryId2), CategoryName: categoryName2, Description: pgtype.Text{String: description2, Valid: true}},
		}, nil).Once()

		products, err := repo.GetProducts(context.Background())

		assert.NoError(t, err)
		assert.Len(t, products, 2)

        product1 := models.Product{Id: productId1, Name: productName1, Category: models.ProductCategory{Id: categoryId1, Name: categoryName1}, Description: description1}
        product2 := models.Product{Id: productId2, Name: productName2, Category: models.ProductCategory{Id: categoryId2, Name: categoryName2}, Description: description2}

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

func TestGetProductById(t *testing.T){
    mockQueries := new(MockQueries)
    repo := &ProductRepository{Queries: mockQueries}

    t.Run("Gets product", func(t *testing.T) {
		expected := models.Product{Id: 1, Name: "product 1"}

		mockQueries.On("GetProductById", mock.Anything).Return(
			GetProductByIdRow{ID: 1, Name: "product 1"},
			nil,
		).Once()

		result, err := repo.GetProductById(context.Background(), 1)

		assert.NoError(t, err)
        assert.Equal(t, expected, result)
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

func TestInsertProduct(t *testing.T){
    mockQueries := new(MockQueries)
    repo := &ProductRepository{Queries: mockQueries}

    t.Run("Inserted product returns product", func(t *testing.T) {
		name :=  "some-product-name"
		description :=  "some-description"
        mockQueries.On("InsertProduct", mock.Anything).Return(Product{ID: 12, Name:name, Description: pgtype.Text{String: description, Valid: true}, CategoryID: pgtype.Int4{Int32: 1, Valid: true}}, nil).Once()
        expected := models.Product{Id: 12, Name: name, Description: description}

        id, err := repo.InsertProduct(context.Background(), models.Product{Name: name, Description: description})

        assert.NoError(t, err)
        assert.Equal(t, expected, id)
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
		os.Setenv("PRODUCT_DB_PORT", "5432")
		os.Setenv("PRODUCT_PG_DATABASE", "testdb")
		os.Setenv("PRODUCT_PG_SSLMODE", "disable")
	}

	unsetEnv := func() {
		os.Unsetenv("PRODUCT_PG_USER")
		os.Unsetenv("PRODUCT_PG_PASSWORD")
		os.Unsetenv("PRODUCT_PG_HOST")
		os.Unsetenv("PRODUCT_DB_PORT")
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