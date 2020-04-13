package main

import (
	"fmt"
	"github.com/CarParkingSystem/parkingspaces"
	"github.com/CarParkingSystem/vehicles"
)

func main() {
	p := parkingspaces.GetParkingLot("TCS", 3)
	fmt.Println(p)
	c1 := vehicles.GetCar("Swift", "Maruti Suzuki", "L53H")
	c2 := vehicles.GetCar("Figo", "Ford", "P21A")
	c3 := vehicles.GetCar("DZire", "Maruti Suzuki", "C1A")
	c4 := vehicles.GetCar("SCross", "Maruti Suzuki", "B12P")
	p.AddCar(c1)
	p.AddCar(c2)
	p.AddCar(c3)
	p.AddCar(c4)
}


