package main

import (
	"bufio"
	"fmt"
	"os"
)

func main_2() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please, input name")
	line, _ := reader.ReadString('\n')
	println("Hello, ", line)
}
