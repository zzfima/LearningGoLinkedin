package main

import "fmt"

func main_5() {
	fmt.Println("Memory allocation ways")

	//incorrect - cannot compile
	/*
		var mem1 = new(map[string]int)
		mem1["key1"] = 11
		fmt.Println(mem1["key1"])
	*/

	//correct
	var mem2 = make(map[string]int)
	mem2["key2"] = 22
	fmt.Println(mem2["key2"])
}
