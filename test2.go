package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var f float32 = 10.4
	i := int(f)
	fmt.Printf("Value of i %d and type %T ", i, i)
	var y string
	fmt.Println(" \n Enter value")
	fmt.Scanf("%s", &y)
	fmt.Printf("Value of Y %s", y)
	reader := bufio.NewReader(os.Stdin)
	name, _, ok := reader.ReadLine()
	if ok == nil {
		fmt.Printf("Name %s", name)
	}

}
