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
	parkedVehiclesLicensePlateList []string
}

func parkingAllotmentCreator(capacity int) map[string]vehicles.Car {
	allotment := make(map[string]vehicles.Car)
	emptyCar := vehicles.Car{}
	for position := 1; position <= capacity; position++ {
		allotment["P"+strconv.Itoa(position)] = emptyCar
	}
	return allotment
}

func GetParkingLot(name string, capacity int) ParkingLot {
	parkinglot := ParkingLot{}
	parkinglot.name = name
	parkinglot.capacity = capacity
	parkinglot.spaceLeft = capacity
	parkinglot.parkingAllotment = parkingAllotmentCreator(capacity)
	parkinglot.parkedVehiclesLicensePlateList = []string{}
	return parkinglot
}

func (p *ParkingLot) GetParkingLotDetails() {
	fmt.Println("Name                      :", p.name)
	fmt.Println("Capacity                  :", p.capacity)
	fmt.Println("Current Space Left        :", p.spaceLeft)
	fmt.Println("ParkingAllotment          :", p.parkingAllotment)
	fmt.Println("Vehicle LicensePlate List :", p.parkedVehiclesLicensePlateList)
}

func (p *ParkingLot) GetParkingSpotDetails(parkingSpot string) {
	if p.parkingAllotment[parkingSpot].GetCarName() == "" {
		fmt.Println("This Parking Spot is empty!")
	} else {
		fmt.Printf("Car %s %s with plates %s is parked at %s.", p.parkingAllotment[parkingSpot].GetCarManufacturer(), p.parkingAllotment[parkingSpot].GetCarName(), p.parkingAllotment[parkingSpot].GetLicensePlate(), parkingSpot)
	}
}

func (p *ParkingLot) AddCar(c vehicles.Car) {
	if p.spaceLeft <= 0 {
		fmt.Println("Parking Space is currently not available in this Parking Lot.")
		return
	} else {
		for k, v := range p.parkingAllotment {
			for _, val := range p.parkedVehiclesLicensePlateList {
				if val == c.GetLicensePlate(){
					fmt.Println("Vehicle already present in the parking lot.")
					return
				} else {
					continue
				}
			}
			if v.GetCarName() == "" {
				p.spaceLeft = p.spaceLeft - 1
				p.parkingAllotment[k] = c
				p.parkedVehiclesLicensePlateList = append(p.parkedVehiclesLicensePlateList, c.GetLicensePlate())
				fmt.Printf("Car %s %s with plates %s is parked at %s in %s\n", c.GetCarManufacturer(), c.GetCarName(), c.GetLicensePlate(), k, p.name)
				fmt.Println("Space left is: ", p.spaceLeft)
				return
			}
		}
	}
}

func (p *ParkingLot) RemoveCar(c vehicles.Car) {
	if len(p.parkedVehiclesLicensePlateList) == 0 {
		fmt.Println("This Parking Lot is empty.")
		return
	} else {
		for _, val := range p.parkedVehiclesLicensePlateList {
			if val == c.GetLicensePlate() {
				for k, _ := range p.parkingAllotment {
					delete(p.parkingAllotment, k)
					p.spaceLeft = p.spaceLeft + 1
					fmt.Printf("Car %s %s with plates %s has left %s\n", c.GetCarManufacturer(), c.GetCarName(), c.GetLicensePlate(), p.name)
					return
				}
			}
		}
		fmt.Printf("Car %s %s with plates %s is not parked at %s\n", c.GetCarManufacturer(), c.GetCarName(), c.GetLicensePlate(), p.name)
	}
}

