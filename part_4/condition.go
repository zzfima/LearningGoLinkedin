package main

import "fmt"

func main_12() {
	if 5 > 6 {
		fmt.Println("ok1")
	} else if 6 > 5 {
		fmt.Println("ok2")
	} else {
		fmt.Println("ok3")
	}

	if x := 4; x > 1 {
		fmt.Println("good")
	}
}
