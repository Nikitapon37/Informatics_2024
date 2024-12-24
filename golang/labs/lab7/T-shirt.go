package lab7

import "fmt"

type TShirt struct {
	name  string
	color string
	size  float64
	price float64
}

func NewTShirt(name string, color string, size float64, price float64) *TShirt {
	tshirt := &TShirt{name: name, color: color, price: price, size: size}
	return tshirt
}

func (t *TShirt) getName() string {
	return t.name
}

func (t *TShirt) setPrice(newPrice float64) {
	t.price = newPrice
}

func (t *TShirt) getPrice() float64 {
	return t.price
}

func (t *TShirt) getSize() float64 {
	return t.size
}

func (t *TShirt) getColor() string {
	return t.color
}

func (t *TShirt) getInfo() {
	fmt.Printf("name: %s\ncolor: %s\nsize: %.2f\nprice: %.2f\n", t.name, t.color, t.size, t.price)
}

func (t *TShirt) applyDiscount(perDiscount float64) {
	t.price = t.price * (1 - (perDiscount / 100))
}
