package task2

import (
	"errors"
)

type Product interface {
	GetID() int
	GetName() string
	GetPrice() int
	GetWeight() int
}

type MyProduct struct {
	id     int
	name   string
	price  int
	weight int
}

func validatePrice(price int) {
	if price < 1 {
		panic("Цена на товар не может быть меньше 1")
	}
}

func NewProduct(id int, name string, price int, weight int) *MyProduct {
	validatePrice(price)
	return &MyProduct{id, name, price, weight}
}

func (p *MyProduct) GetID() int {
	return p.id
}

func (p *MyProduct) GetName() string {
	return p.name
}

func (p *MyProduct) GetPrice() int {
	return p.price
}

func (p *MyProduct) GetWeight() int {
	return p.weight
}

type Basket interface {
	AddProduct(product Product, count int) error
	DeleteProduct(id int) error
	ListProducts() []Product
	GetPrice() int
	GetShippingCost() int
}

type BasketItem struct {
	product Product
	count   int
}

type MyBasket struct {
	items       map[int]BasketItem
	totalWeight int
	totalPrice  int
}

func NewBasket() *MyBasket {
	return &MyBasket{
		items:       make(map[int]BasketItem),
		totalWeight: 0,
		totalPrice:  0,
	}
}

const (
	maxWeight = 100
	maxItems  = 30
)

func (b *MyBasket) AddProduct(product Product, count int) error {
	if count <= 0 {
		return errors.New("количество товара должно быть больше 0")
	}
	if len(b.items) >= maxItems {
		return errors.New("максимальное количество товаров в корзине не должно быть больше 30")
	}

	newWeight := b.totalWeight + product.GetWeight()*count
	if newWeight > maxWeight {
		return errors.New("превышен максимальный вес товаров в корзине")
	}

	item, exists := b.items[product.GetID()]
	if exists {
		item.count += count
		b.items[product.GetID()] = item
	} else {
		b.items[product.GetID()] = BasketItem{product: product, count: count}
	}
	b.totalWeight += product.GetWeight() * count
	b.totalPrice += product.GetPrice() * count
	return nil
}

func (b *MyBasket) DeleteProduct(id int) error {
	item, exists := b.items[id]
	if !exists {
		return errors.New("товар не найден в корзине")
	}

	b.totalWeight -= item.product.GetWeight() * item.count
	b.totalPrice -= item.product.GetPrice() * item.count
	delete(b.items, id)
	return nil
}

func (b *MyBasket) ListProducts() []Product {
	var products []Product
	for _, item := range b.items {
		products = append(products, item.product)
	}
	return products
}

func (b *MyBasket) GetPrice() int {
	return b.totalPrice + b.GetShippingCost()
}

func (b *MyBasket) GetShippingCost() int {
	switch {
	case b.totalPrice < 500:
		return 250
	case b.totalPrice < 1000:
		return 100
	default:
		return 0
	}
}
