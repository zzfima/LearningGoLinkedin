package main

import "fmt"

func main_7() {
	var colors [3]string
	colors[0] = "red0"
	colors[1] = "red1"
	colors[2] = "red2"
	fmt.Println(colors[1])
	fmt.Println(colors)

	numbers := [4]int{4, 44, 444, 4444}
	fmt.Println(numbers, len(numbers))
}
