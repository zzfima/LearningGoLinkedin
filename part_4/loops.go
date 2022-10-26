package main

import "fmt"

func main_14() {
	colors := [3]string{"r1", "c1", "d1"}

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for i, j := range colors {
		fmt.Println(i, j)
	}

	if 6 > 3 {
		goto theEnd
	}

	fmt.Println("not here")

theEnd:
	fmt.Println("finish")
}
