package main

import "fmt"

func main15() {
	res := addTwoNumbers(3, 4)
	fmt.Println(res)

	res, l := addNumbers(4, 5, 6, 7)
	fmt.Println(res, l)
}

func addTwoNumbers(v1 int, v2 int) int {
	return v1 + v2
}

func addNumbers(values ...int) (int, int) {
	sum := 0

	for i := range values {
		sum += values[i]
	}

	return sum, len(values)
}
