package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main_13() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(3)
	fmt.Println(dow)

	switch dow {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("panic!")
	}

}
