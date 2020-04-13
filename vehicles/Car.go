package vehicles

type Car struct {
	name string
	manufacturer string
	licensePlate string
}

var CarList []Car

func GetCar(name string, manufacturer string, licensePlate string) Car{
	c := Car{}
	c.name = name
	c.manufacturer = manufacturer
	c.licensePlate= licensePlate
	CarList = append(CarList, c)
	return c
}

func (c Car) GetCarName() string{
	return c.name
}

func (c Car) GetCarManufacturer() string{
	return c.manufacturer
}

func (c Car) GetLicensePlate() string{
	return c.licensePlate
}
