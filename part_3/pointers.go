package main

import "fmt"

func main() {
	int1 := 55
	ptr1 := &int1
	fmt.Println("value:", *ptr1, "address:", ptr1)
	add(ptr1)
	fmt.Println("value:", *ptr1, "address:", ptr1)
}

func add(ptr *int) {
	(*ptr)++
}
