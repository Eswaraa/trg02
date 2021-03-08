package main

import (
	//"errors"
	"fmt"
)

//variadic function
func sum(i ...int) int {
	var sum int = 0
	for _, n := range i {
		sum = sum + n
	}
	return sum
}

type Phone struct {
	Camera
	resolution int
	screen     int
}
type Camera struct {echo 
	resolution int
}

func main() {
	a := []int{4, 6, 6, 7, 8, 6, 2}
	weightage := make(map[string]int)
	weightage["High"] = 10
	weightage["medium"] = 5
	weightage["low"] = 1

	fmt.Printf("Test %d", sum(a...))
	p := Phone{Camera{3}, 5, 4}
	fmt.Printf("Test %v %T", p, p)
	fmt.Printf("\n Camera Resolution %d", p.Camera.resolution)
	fmt.Printf("\n Screen Resolution %d", p.resolution)

	for k, n := range weightage {
		fmt.Printf("\n Weightage : %s , value: %d", k, n)
	}

}
