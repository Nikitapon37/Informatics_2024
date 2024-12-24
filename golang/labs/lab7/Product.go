package lab7

type Products interface {
	getName() string
	setPrice(float64)
	getPrice() float64
	getInfo()
	applyDiscount(float64)
}
