package main

import "fmt"

type Pizza struct {
	Price int
}

func (p Pizza) GetPrice() int {
	return p.Price
}

type PizzaWithAnana struct {
	Pizza // embedding
	AnanaPrice int
}

func (p PizzaWithAnana) GetPrice() int {
	return p.Pizza.GetPrice() + p.AnanaPrice
}

func main() {
	pizza := Pizza{Price: 2500}

	fmt.Println(pizza.GetPrice())

	pizzaWithAnana := PizzaWithAnana{
		Pizza: pizza,
		AnanaPrice: 500,
	}

	fmt.Println(pizzaWithAnana.GetPrice())
}