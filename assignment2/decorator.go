package main

import "fmt"

type ICoffee interface {
	getPrice() int
}

type Espresso struct {
}

func (e *Espresso) getPrice() int {
	return 20
}

type Milk struct {
	coffee ICoffee
}

func (c *Milk) getPrice() int {
	coffeePrice := c.coffee.getPrice()
	return coffeePrice + 5
}

type Chocolate struct {
	coffee ICoffee
}

func (c *Chocolate) getPrice() int {
	coffeePrice := c.coffee.getPrice()
	return coffeePrice + 10
}

func main() {
	coffee := &Espresso{}
	coffeeWithMilk := &Milk{
		coffee: coffee,
	}
	coffeeWithMilkAndChocolate := &Chocolate{
		coffee: coffeeWithMilk,
	}
	fmt.Printf("Price of espresso with milk and chocolate is %d\n", coffeeWithMilkAndChocolate.getPrice())
}
