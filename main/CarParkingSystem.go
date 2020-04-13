package main

import (
	"fmt"
	"github.com/CarParkingSystem/parkingspaces"
	"github.com/CarParkingSystem/vehicles"
)

func main() {
	p := parkingspaces.GetParkingLot("TCS", 3)
	fmt.Println(p)
	// p1 := parkingspaces.GetParkingLot("Tieto", 5)
	c1 := vehicles.GetCar("Swift", "Maruti Suzuki", "MH12BP63")
	c2 := vehicles.GetCar("Figo", "Ford", "MH04BS85")
	c3 := vehicles.GetCar("DZire", "Maruti Suzuki", "MH04BD12")
	c4 := vehicles.GetCar("SCross", "Maruti Suzuki", "MH04GH56")
	c5 := vehicles.GetCar("SX4", "Maruti Suzuki", "MH04BD45")
	p.AddCar(c1)
	p.AddCar(c2)
	p.AddCar(c3)
	p.AddCar(c4)
	p.RemoveCar(c5)
	p.GetParkingSpotDetails("P1")
}


