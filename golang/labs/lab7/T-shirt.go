package lab7

import "fmt"

type Tshirt struct {
	name  string
	color string
	size  float64
	price float64
}

func NewTshirt(name string, color string, size float64, price float64) *Tshirt {
	tshirt := &Tshirt{name: name, color: color, price: price, size: size}
	return tshirt
}

func (t *Tshirt) getName() string {
	return t.name
}

func (t *Tshirt) setPrice(newPrice float64) {
	t.price = newPrice
}

func (t *Tshirt) getPrice() float64 {
	return t.price
}

func (t *Tshirt) getSize() float64 {
	return t.size
}

func (t *Tshirt) getColor() string {
	return t.color
}

func (t *Tshirt) getInfo() {
	fmt.Printf("name: %s\ncolor: %s\nsize: %.2f\nprice: %.2f\n", t.name, t.color, t.size, t.price)
}

func (t *Tshirt) applyDiscount(perDiscount float64) {
	t.price = t.price * (1 - (perDiscount / 100))
}
