package main

import (
	"fmt"
	"sort"
)

func main_8() {
	//usual array
	colors := [3]string{"red", "green", "blue"}
	fmt.Println(colors)
	fmt.Printf("%T\n", colors)

	//slice
	numbers := []int{4444, 44, 444}
	fmt.Println(numbers)
	fmt.Printf("%T\n", numbers)
	//append numbers
	numbers = append(numbers, 5)
	fmt.Println(numbers)

	//sorting
	sort.Ints(numbers)
	fmt.Println(numbers)
}
