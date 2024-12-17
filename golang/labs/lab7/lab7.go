package lab7

import (
	"fmt"
)

type Products interface {
	getName() string
	setPrice(float64)
	getPrice() float64
	getInfo()
	applyDiscount(float64)
}

func calculateTotalPrice(products []Products) float64 {
	var sum float64 = 0
	for _, product := range products {
		sum += product.getPrice()
	}

	return sum
}

func RunLab7() {
	product1 := NewTshirt("Nike", "White", 70, 1000)
	product2 := NewChips("Lays", "Crab", 100, 150)

	listOfProduct := []Products{product1, product2}

	fmt.Println("Товар-----------Цена")
	for _, product := range listOfProduct {
		fmt.Printf("%s-------%.2f ₽\n", product.getName(), product.getPrice())
	}
	fmt.Printf("Цена корзины до скидки: %.2f ₽\n\n", calculateTotalPrice(listOfProduct))

	product1.applyDiscount(10)
	product2.applyDiscount(27)

	fmt.Println("Товар-----------Цена")
	for _, product := range listOfProduct {
		fmt.Printf("%s-------%.2f ₽\n", product.getName(), product.getPrice())
	}
	fmt.Printf("Цена корзины после скидки: %.2f ₽\n", calculateTotalPrice(listOfProduct))

	fmt.Println("Информация про футболку")
	fmt.Println("Цвет:", product1.getColor())
	fmt.Println("Размер:", product1.getSize())
	fmt.Println("Бренд:", product1.getName())

	fmt.Println("Информация про xbgcs")
	fmt.Println("Вкус:", product2.getTaste())
	fmt.Println("Масса:", product2.getWeight())
	fmt.Println("Название:", product2.getName())
}
