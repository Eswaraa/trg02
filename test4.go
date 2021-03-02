package main

import "fmt"

type Direction int

const (
	North int = iota
	South
	East
	West
)

func main() {

	// Declaring a variable myDirection with type Direction
	var myDirection int
	myDirection = West

	if myDirection == West {
		fmt.Println("myDirection is West:", myDirection)
	}

}
