package vehicles

type Car struct {
	name string
	manufacturer string
	licensePlate string
}

func GetCar(name string, manufacturer string, licensePlate string) Car{
	c := Car{
		name:         name,
		manufacturer: manufacturer,
		licensePlate: licensePlate,
	}
	return c
}

