package task2

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddProduct(t *testing.T) {
	basket := NewBasket()
	product := NewProduct(1, "Fake Product", 99, 5)

	err := basket.AddProduct(product, 20)
	require.NoError(t, err, "нет ошибок при добавлении товара в корзину")

	assert.Equal(t, 1, len(basket.items), "в корзине должен быть 1 товар")
	assert.Equal(t, 100, basket.totalWeight, "общий вес должен быть 100")
}

func TestAddProductZeroPrice(t *testing.T) {
	assert.Panics(t, func() {
		_ = NewProduct(13, "Fake Zero Price Product", 0, 1)
	}, "паника при попытке создать продукт с ценой 0")
}

func TestAddProductHeavyWeight(t *testing.T) {
	basket := NewBasket()
	product := NewProduct(2, "Heavy Fake Product", 25, 15)

	err := basket.AddProduct(product, 7)
	require.Error(t, err, "превышен допустимый вес товара")
}

func TestAddProductMaxCount(t *testing.T) {
	basket := NewBasket()

	for i := 1; i <= 30; i++ {
		product := NewProduct(i, "Product"+strconv.Itoa(i), 20, 1)
		err := basket.AddProduct(product, 1)
		assert.NoError(t, err, "нет ошибок при добавлении товара %d в корзину", i)
	}

	product := NewProduct(31, "Product31", 20, 1)
	err := basket.AddProduct(product, 1)

	assert.Error(t, err, "ошибка при добавлении 31-го товара")
	assert.Equal(t, "максимальное количество товаров в корзине не должно быть больше 30", err.Error())
}

func TestDeleteProduct(t *testing.T) {
	basket := NewBasket()
	product := NewProduct(3, "Fake Product", 77, 10)

	err := basket.AddProduct(product, 2)
	require.NoError(t, err, "нет ошибок при добавлении товара в корзину")

	err = basket.DeleteProduct(3)
	require.NoError(t, err, "нет ошибок при удаление товара")

	assert.Equal(t, 0, len(basket.items), "корзина должна быть пуста после удаления товара")
	assert.Equal(t, 0, basket.totalWeight, "общий вес корзины равен 0 после удаления товара")
}

func TestDeleteNonexistentProduct(t *testing.T) {
	basket := NewBasket()
	product := NewProduct(77, "Fake Product", 50, 5)

	err := basket.DeleteProduct(product.GetID())

	require.Error(t, err, "возвращена ошибка при удалении несуществующего товара")
	assert.Equal(t, "товар не найден в корзине", err.Error(), "должно быть сообщение об ошибке 'товар не найден в корзине'")
}

func TestListProducts(t *testing.T) {
	basket := NewBasket()
	product1 := NewProduct(100, "Product 1", 100, 1)
	product2 := NewProduct(200, "Product 2", 200, 2)
	product3 := NewProduct(300, "Product 3", 300, 3)

	_ = basket.AddProduct(product1, 2)
	_ = basket.AddProduct(product2, 1)
	_ = basket.AddProduct(product3, 3)

	products := basket.ListProducts()

	require.Len(t, products, 3, "количество товаров в корзине должно быть 3")

	assert.Contains(t, products, product1, "товар 1 должен быть в списке")
	assert.Contains(t, products, product2, "товар 2 должен быть в списке")
	assert.Contains(t, products, product3, "товар 3 должен быть в списке")
}

func TestGetPrice(t *testing.T) {
	basket := NewBasket()
	product1 := NewProduct(1000, "Product 1000", 100, 10)
	product2 := NewProduct(2000, "Product 2000", 200, 20)
	_ = basket.AddProduct(product1, 1)
	_ = basket.AddProduct(product2, 2)

	expectedPrice := 500 + basket.GetShippingCost()
	totalPrice := basket.GetPrice()
	assert.Equal(t, expectedPrice, totalPrice, "итоговая стоимость равна 600 с учетом стоимости доставки")
}

func TestGetShippingCost(t *testing.T) {
	basket := NewBasket()
	product := NewProduct(999, "Product 1", 100, 10)
	_ = basket.AddProduct(product, 4)

	shippingCost := basket.GetShippingCost()
	assert.Equal(t, 250, shippingCost, "стоимость доставки должна быть 250 при цене товаров < 500")

	_ = basket.AddProduct(product, 1)
	shippingCost = basket.GetShippingCost()
	assert.Equal(t, 100, shippingCost, "стоимость доставки должна быть 100 при цене товаров 500-999")

	_ = basket.AddProduct(product, 5)
	shippingCost = basket.GetShippingCost()
	assert.Equal(t, 0, shippingCost, "стоимость доставки должна быть 0 при цене товаров >= 1000")
}
