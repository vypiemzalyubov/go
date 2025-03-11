package task3

import (
	"errors"
	"math/rand"
	"time"
)

type Product interface {
	GetName() string
	GetDimensions() (int, int, int)
	GetWeight() int
}

type MyProduct struct {
	name   string
	length int
	width  int
	height int
	weight int
}

func NewProduct(name string, length, width, height, weight int) *MyProduct {
	return &MyProduct{name, length, width, height, weight}
}

func (p *MyProduct) GetName() string {
	return p.name
}

func (p *MyProduct) GetDimensions() (int, int, int) {
	return p.length, p.width, p.height
}

func (p *MyProduct) GetWeight() int {
	return p.weight
}

type PackagedProduct interface {
	GetProduct() Product
}

type MyPackagedProduct struct {
	product Product
}

func NewPackagedProduct(product Product) *MyPackagedProduct {
	return &MyPackagedProduct{product: product}
}

func (pp *MyPackagedProduct) GetProduct() Product {
	return pp.product
}

type Сonveyor interface {
	Add(product Product) error
	CheckNextProductDimensions() error
	CheckNextProductWeight() error
	CheckNextProductDefective() error
	ProductPackaging() error
	Get() []PackagedProduct
}

type MyConveyor struct {
	products         []Product
	packagedProducts []PackagedProduct
	rng              *rand.Rand
}

func NewConveyor() *MyConveyor {
	return &MyConveyor{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (c *MyConveyor) checkProductCount() error {
	if len(c.products) > 2 {
		return errors.New("на конвейере не может быть более 2 товаров")
	}
	return nil
}

func (c *MyConveyor) Add(product Product) error {
	if len(c.products) >= 2 {
		return errors.New("на конвейере не может быть более 2 товаров")
	}
	c.products = append(c.products, product)
	return nil
}

func (c *MyConveyor) CheckNextProductDimensions() error {
	if len(c.products) == 0 {
		return errors.New("нет товаров на конвейере для проверки габаритов")
	}

	if err := c.checkProductCount(); err != nil {
		return err
	}

	p := c.products[0]
	length, width, height := p.GetDimensions()
	sumDimensions := length + width + height
	if sumDimensions > 250 || sumDimensions < 50 {
		c.products = c.products[1:]
		return errors.New("сумма габаритов товара больше 250 или меньше 50")
	}
	return nil
}

func (c *MyConveyor) CheckNextProductWeight() error {
	if len(c.products) == 0 {
		return errors.New("нет товаров на конвейере для проверки веса")
	}

	if err := c.checkProductCount(); err != nil {
		return err
	}

	p := c.products[0]
	weight := p.GetWeight()
	if weight > 500 || weight < 10 {
		c.products = c.products[1:]
		return errors.New("вес товара превышает 500 или меньше 10")
	}
	return nil
}

func (c *MyConveyor) CheckNextProductDefective() error {
	if len(c.products) == 0 {
		return errors.New("нет товаров на конвейере для проверки на брак")
	}

	if err := c.checkProductCount(); err != nil {
		return err
	}

	if c.rng.Float32() < 0.5 {
		c.products = c.products[1:]
		return errors.New("товар является бракованным")
	}
	return nil
}

func (c *MyConveyor) ProductPackaging() error {
	if len(c.products) == 0 {
		return errors.New("нет товаров на конвейере для упаковки")
	}

	if len(c.packagedProducts) >= 10 {
		return errors.New("на конвейере не может быть более 10 упакованных товаров")
	}

	p := c.products[0]
	c.products = c.products[1:]
	packagedProduct := NewPackagedProduct(p)
	c.packagedProducts = append(c.packagedProducts, packagedProduct)
	return nil
}

func (c *MyConveyor) Get() []PackagedProduct {
	return c.packagedProducts
}
