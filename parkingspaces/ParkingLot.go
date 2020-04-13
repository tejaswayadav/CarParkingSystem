package parkingspaces

import (
	"fmt"
	"github.com/CarParkingSystem/vehicles"
	"strconv"
)

type ParkingLot struct {
	name             string
	capacity         int
	spaceLeft        int
	parkingAllotment map[string]vehicles.Car
}

func GetParkingLot(name string, capacity int) ParkingLot {
	parkinglot := ParkingLot{}
	parkinglot.name = name
	parkinglot.capacity = capacity
	parkinglot.spaceLeft = capacity
	parkinglot.parkingAllotment = parkingAllotmentCreator(capacity)
	return parkinglot
}

func parkingAllotmentCreator(capacity int) map[string]vehicles.Car {
	allotment := make(map[string]vehicles.Car)
	emptyCar := vehicles.Car{}
	for position := 1; position <= capacity; position++ {
		allotment["P"+strconv.Itoa(position)] = emptyCar
	}
	return allotment
}

func (p *ParkingLot) AddCar(c vehicles.Car) {
	if p.spaceLeft <= 0 {
		fmt.Println("Parking Space is currently not available in this Parking Lot.")
	} else {
		for k, v := range p.parkingAllotment {
			if v.GetCarName() == "" {
				p.spaceLeft = p.spaceLeft - 1
				p.parkingAllotment[k] = c
				fmt.Printf("Car %s %s with plates %s is parked at %s in %s\n", c.GetCarManufacturer(), c.GetCarName(), c.GetLicensePlate(), k, p.name)
				fmt.Println("Space left is: ", p.spaceLeft)
				break
			}
		}
	}
}
