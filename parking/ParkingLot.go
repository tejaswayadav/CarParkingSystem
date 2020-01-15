package parking

import (
	"github.com/CarParkingSystem/vehicles"
	"strconv"
)

type ParkingLot struct {
	name string
	capacity int
	spaceLeft int
	parkingAllotment map[string] vehicles.Car
}

func GetParkingLot(name string, capacity int) ParkingLot{
	parkinglot := ParkingLot{
		name:             name,
		capacity:         capacity,
		spaceLeft:        capacity,
		parkingAllotment: parkingAllotmentCreator(capacity),
	}
	return parkinglot
}

func parkingAllotmentCreator(capacity int) map[string] vehicles.Car{
	allotment := make(map[string]vehicles.Car)
	emptyCar := vehicles.GetCar("nil", "nil", "nil")
	for position := 1; position <= capacity+1; position++ {
		allotment["P"+strconv.Itoa(position)] = emptyCar
	}
	return allotment
}

func (p ParkingLot) AddCar(c vehicles.Car){
	if p.spaceLeft > 0{

	}
}
