package main

import (
	"fmt"
)

func main() {
	var vehicle iVehicle

	car := Car{
		name:     "Lexus",
		maxSpeed: 250,
	}

	vehicle = &car

	vehicle.move()

	plane := Plane{
		name:      "Boing",
		maxSpeed:  900,
		flyMethod: "Jet",
	}

	vehicle = &plane

	vehicle.move()
}

type iVehicle interface {
	move()
}

type Car struct {
	name     string
	maxSpeed int
}

func (car *Car) move() {
	fmt.Printf("%s is moving on the ground", car.name)
}

type Plane struct {
	name      string
	maxSpeed  int
	flyMethod string
}

func (plane *Plane) move() {
	fmt.Printf("%s is flying in the sky", plane.name)
}
