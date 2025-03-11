package task3

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ConveyorTestSuite struct {
	suite.Suite
	conveyor *MyConveyor
}

func (suite *ConveyorTestSuite) SetupTest() {
	suite.conveyor = NewConveyor()
}

func (suite *ConveyorTestSuite) TestAddProduct() {
	testcases := []struct {
		name          string
		products      []*MyProduct
		expectedCount int
	}{
		{
			name: "Добавление одного продукта",
			products: []*MyProduct{
				NewProduct("ProductFake1", 50, 10, 10, 50),
			},
			expectedCount: 1,
		},
		{
			name: "Добавление двух продуктов",
			products: []*MyProduct{
				NewProduct("ProductFake1", 10, 30, 20, 200),
				NewProduct("ProductFake2", 45, 10, 15, 150),
			},
			expectedCount: 2,
		},
	}

	for _, tc := range testcases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.conveyor = NewConveyor()

			for _, product := range tc.products {
				err := suite.conveyor.Add(product)
				require.NoError(suite.T(), err)
			}

			require.Equal(suite.T(), tc.expectedCount, len(suite.conveyor.products))
		})
	}
}

func (suite *ConveyorTestSuite) TestAddProductNegative() {
	product1 := NewProduct("Product1", 1, 70, 10, 50)
	suite.conveyor.Add(product1)
	product2 := NewProduct("Product2", 20, 10, 40, 150)
	suite.conveyor.Add(product2)
	product3 := NewProduct("Product3", 10, 5, 60, 200)
	err := suite.conveyor.Add(product3)

	require.Error(suite.T(), err)
	require.Equal(suite.T(), "на конвейере не может быть более 2 товаров", err.Error())
	require.Equal(suite.T(), 2, len(suite.conveyor.products))
}

func (suite *ConveyorTestSuite) TestCheckProductDimensions() {
	product := NewProduct("TestProduct Positive", 10, 30, 15, 100)
	err := suite.conveyor.Add(product)
	require.NoError(suite.T(), err)

	err = suite.conveyor.CheckNextProductDimensions()
	require.NoError(suite.T(), err)
}

func (suite *ConveyorTestSuite) TestCheckProductDimensionsNegative() {
	testcases := []struct {
		name          string
		dimensions    [3]int
		expectedError string
	}{
		{
			name:          "Сумма габаритов товара меньше 50",
			dimensions:    [3]int{1, 5, 10},
			expectedError: "сумма габаритов товара больше 250 или меньше 50",
		},
		{
			name:          "Сумма габаритов товара больше 250",
			dimensions:    [3]int{200, 30, 30},
			expectedError: "сумма габаритов товара больше 250 или меньше 50",
		},
	}

	for _, tc := range testcases {
		tc := tc
		suite.Run(tc.name, func() {
			product := NewProduct("TestProduct", tc.dimensions[0], tc.dimensions[1], tc.dimensions[2], 100)
			err := suite.conveyor.Add(product)
			require.NoError(suite.T(), err)

			err = suite.conveyor.CheckNextProductDimensions()
			require.Error(suite.T(), err)
			require.Equal(suite.T(), tc.expectedError, err.Error())
		})
	}
}

func (suite *ConveyorTestSuite) TestCheckProductWeight() {
	product := NewProduct("HeavyProduct", 20, 20, 15, 600)

	err := suite.conveyor.Add(product)
	require.NoError(suite.T(), err)

	err = suite.conveyor.CheckNextProductWeight()
	require.Error(suite.T(), err)
	require.Equal(suite.T(), "вес товара превышает 500 или меньше 10", err.Error())
}

func (suite *ConveyorTestSuite) TestCheckProductDefective() {
	suite.conveyor.rng = rand.New(rand.NewSource(1))

	product := NewProduct("Product1", 25, 10, 25, 300)
	suite.conveyor.Add(product)

	err := suite.conveyor.CheckNextProductDefective()
	require.NoError(suite.T(), err)
}

func (suite *ConveyorTestSuite) TestCheckProductDefectiveNegative() {
	suite.conveyor.rng = rand.New(rand.NewSource(42))

	product := NewProduct("DefectiveProduct", 25, 15, 30, 100)
	suite.conveyor.Add(product)

	err := suite.conveyor.CheckNextProductDefective()
	require.Error(suite.T(), err)
	require.Equal(suite.T(), "товар является бракованным", err.Error())
}

func (suite *ConveyorTestSuite) TestProductPackaging() {
	product := NewProduct("Product1", 35, 10, 40, 200)
	suite.conveyor.Add(product)

	err := suite.conveyor.ProductPackaging()
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 1, len(suite.conveyor.Get()))
	require.Equal(suite.T(), 0, len(suite.conveyor.products))

	err = suite.conveyor.ProductPackaging()
	require.Error(suite.T(), err)
	require.Equal(suite.T(), "нет товаров на конвейере для упаковки", err.Error())
}

func TestConveyorTestSuite(t *testing.T) {
	suite.Run(t, new(ConveyorTestSuite))
}
