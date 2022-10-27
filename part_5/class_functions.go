package main

import "fmt"

func main16() {
	dog := Dog{"Mike", 5.6}
	dog.PrintInfo()

	dog.SpeakThreeTimes()
	dog.SpeakThreeTimes()

	dog.SpeakFourTimes()
	dog.SpeakFourTimes()
}

//Dog describe dog
type Dog struct {
	Name string
	Age  float32
}

//PrintInfo print info about dog
func (d Dog) PrintInfo() {
	fmt.Println("My name is:", d.Name, ", my age is:", d.Age)
}

//SpeakThreeTimes 3 times print name. d object passed  by value
func (d Dog) SpeakThreeTimes() {
	d.Name = fmt.Sprintf("%v %v %v", d.Name, d.Name, d.Name)
	fmt.Println(d.Name)
}

//SpeakFourTimes 4 times print name. d object passed  by reference
func (d *Dog) SpeakFourTimes() {
	d.Name = fmt.Sprintf("%v %v %v %v", d.Name, d.Name, d.Name, d.Name)
	fmt.Println(d.Name)
}
