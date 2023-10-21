package main

import (
	"fmt"
)

type Vehicle interface {
	start()
}

type Car struct {
	brand  string
	model string
}

func (c Car) start() {
	fmt.Printf("Starting the %s %s\n", c.brand, c.model)
}

type Truck struct {
	brand string
}

func (t Truck) start() {
	fmt.Printf("Pedaling the %s bicycle\n", t.brand)
}

func main() {

	myCar := Car{brand: "Toyota", model: "Camry"}
	myTruck := Truck{brand: "Trek"}


	vehicles := []Vehicle{myCar, myTruck}

	for _, v := range vehicles {
		v.start()
	}
}
