package main

import "fmt"

func main() {
	fmt.Printf("Hello \n")
	var x int = 10
	y := -20
	var x1 int32
	var x2 int64
	var f1 float32 = 41.2
	var flag bool = true
	var str string = " Hello"
	s1 := "Another string"

	fmt.Printf("X :=> %b Y:=> %d", x, y)
	fmt.Printf("\nX type: %T", x)
	fmt.Printf("\nX1 type: %T", x1)
	fmt.Printf("\nX2 type: %T", x2)
	fmt.Printf("\nX2 type: %f", f1)
	fmt.Printf("\nX2 type: %t", flag)
	fmt.Printf("\nX2 type: %s", str)
	fmt.Printf("\nX2 type: %s", s1)

	if x = -3; x > 0 && y > 0 {
		fmt.Println("\n x and y is +ve")
	} else if x = 0; x < 0 {
		fmt.Println("\n x is -ve")
	} else {
		fmt.Println("\n x is zero")
	}
	a := 40
	switch a {
	default:
		fmt.Println("\n a is not in range")
	case 1:
		fmt.Println("\n a is one")
	case 20:
		fmt.Println("\n a is twenty")
		fallthrough
	case 30:
		fmt.Println("\n a is thirty")

	}

}
