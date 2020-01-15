package CarParkingSystem

type car struct {
	name string
	manufacturer string
	licensePlate string
}

func Car(name string, manufacturer string, licensePlate string) car{
	c := car{name, manufacturer, licensePlate,}
	return c
}

