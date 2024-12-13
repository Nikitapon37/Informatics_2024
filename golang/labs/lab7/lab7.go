package lab7

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	Discount(discount float64)
}

type Chips struct {
	Name  string
	Price float64
	Taste string
}

func (c *Chips) GetName() string {
	return c.Name
}

func (c *Chips) GetPrice() float64 {
	return c.Price
}

func (c *Chips) SetPrice(price float64) {
	c.Price = price
}

func (c *Chips) Discount(discount float64) {
	c.Price -= c.Price * discount / 100
}

type Eggs struct {
	Name     string
	Price    float64
	Quantity float64
}

func (e *Eggs) GetName() string {
	return e.Name
}

func (e *Eggs) GetPrice() float64 {
	return e.Price
}

func (e *Eggs) SetPrice(price float64) {
	e.Price = price
}

func (e *Eggs) Discount(discount float64) {
	e.Price -= e.Price * discount / 100
}

func Calculate(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func RunLab7() {
	Chips := &Chips{"Lays", 159.99, "Сметана лук"}
	Eggs := &Eggs{"Деревенские", 100.00, 10}

	products := []Product{Chips, Eggs}
	fmt.Println("Общая стоимость товаров:", Calculate(products))

	Chips.Discount(12)
	Eggs.Discount(20)
	fmt.Println("Общая стоимость товаров после применения скидок:", Calculate(products))
}
