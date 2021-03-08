package main

import (
	"errors"
	"fmt"
)

type Drivable interface {
	Drive(acc int) (int, error)
	Brake(amount int) int
}

type Car struct {
	Make  string
	Model string
	Price float32
	Speed int
}

type Bike struct {
	Make  string
	Model string
	Price float32
	Speed int
}

func (bike *Bike) Drive(acc int) (int, error) {
	//err := false
	return bike.Speed, nil
}
func (bike *Bike) Brake(acc int) int {
	//err := false
	return 0
}

func (car *Car) Drive(acc int) (int, error) {
	//err := false
	if car.Speed+acc*2 > 100 {
		return car.Speed, errors.New("Overspeed")
	}
	car.Speed = car.Speed + acc*2

	return car.Speed, nil

}

func (car *Car) Brake(amount int) int {
	car.Speed = car.Speed - amount*2
	return car.Speed
}

func (car *Car) increasePrice(inc float32) {
	car.Price = car.Price + inc
	fmt.Printf("\nchanged Car %v", car)
}

func (car Car) Print() {
	fmt.Printf("\nCar Make is %s, Model - %s with Price %f, Speed: %d", car.Make, car.Model, car.Price, car.Speed)
}

func driveVehicle(vehicle Drivable, i int) (int, error) {
	return vehicle.Drive(i)
}

func main() {
	// 	fmt.Printf("Name %s", name)

	var c Car = Car{Make: "Ford", Model: "Endeavor", Price: 30.5}
	c.increasePrice(2)

	c.Speed = 20
	c.Print()
	fmt.Printf("\nCar %v", &c)
	// var drive Drivable
	// drive = &c
	// drive.Drive(5)
	c.Print()
	c.Brake(5)
	c.Print()
	s, err := driveVehicle(&c, 50)

	if err != nil {
		fmt.Printf("\n Overspeed %d", s)
	} else {
		fmt.Printf("\n Peed change %d", s)
	}
	b := Bike{Make: "Honds", Model: "XC", Price: 10.5}
	b.Speed = 50
	s1, err1 := driveVehicle(&b, 50) //Dynamic Polymorpism
	if err1 != nil {
		fmt.Printf("\n Overspeed %d", s1)
	} else {
		fmt.Printf("\n Peed change %d", s1)
	}

	//int, float, string boo, [4]arr, slice
	// Object

}
