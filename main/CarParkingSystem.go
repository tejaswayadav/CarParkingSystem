package main

import (
	"fmt"
	"github.com/CarParkingSystem/parking"
	"github.com/CarParkingSystem/vehicles"
)

func main() {
	p := parking.GetParkingLot("TCS", 10)
	fmt.Println(p)
	c1 := vehicles.GetCar("ns", "maruti", "L53H")
	fmt.Println(c1)
}


