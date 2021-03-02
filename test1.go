package main

import "fmt"

func main() {
	fmt.Printf("Hello \n")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Hello %d\n", i)

	}
	var arr [4]int = [4]int{10, 20, 30, 40}

	var slc []int

	slc = arr[1:3]

	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] %d\n", i, arr[i])

	}
	for k, v := range arr {
		fmt.Printf("arr[%d] %d\n", k, v)
	}
	fmt.Printf("\n arr %v", arr)
	fmt.Printf("\n slc %v", slc)

	slc[0] = 6

	for k1, v1 := range slc {
		fmt.Printf("\nslc[%d] %d", k1, v1)
	}
	slc2 := []int{10, 20, 30}

	fmt.Printf("\n slc %v", slc)
	fmt.Printf("\n arr %v", arr)
	fmt.Printf("\n arr type %T", arr)

	fmt.Printf("\n Slc Type %T", slc)

	for k1, v1 := range slc2 {
		fmt.Printf("\nslc2[%d] %d", k1, v1)
	}
	slc2 = append(slc2, 8, 9, 10)
	fmt.Printf("\n slc2 %v", slc2)

	slc2 = append(slc2[:3], slc2[4:]...)
	// slc[:i], slc[i+1:]...
	fmt.Printf("\n slc2 %v", slc2)

}
