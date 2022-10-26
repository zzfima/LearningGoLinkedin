package main

import "fmt"

func main_11() {
	dog1 := Dog{11, "Wenti", false}
	fmt.Println(dog1)
	cat1 := cat{12, "Diona", false}
	fmt.Println(cat1)
}

type Dog struct {
	Age         int
	Name        string
	Initialized bool
}

type cat struct {
	age         int
	name        string
	initialized bool
}
