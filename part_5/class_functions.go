package main

import "fmt"

func main() {
	dog := Dog{"Mike", 5.6}
	dog.PrintInfo()
}

type Dog struct {
	Name string
	Age  float32
}

func (d Dog) PrintInfo() {
	fmt.Println("My name is:", d.Name, ", my age is:", d.Age)
}
