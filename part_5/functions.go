package main

import "fmt"

func main_15() {
	res := add_two_numbers(3, 4)
	fmt.Println(res)

	res, l := add_numbers(4, 5, 6, 7)
	fmt.Println(res, l)
}

func add_two_numbers(v1 int, v2 int) int {
	return v1 + v2
}

func add_numbers(values ...int) (int, int) {
	sum := 0

	for i := range values {
		sum += values[i]
	}

	return sum, len(values)
}
