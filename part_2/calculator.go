package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)

	fmt.Println("Please, input value 1")
	input1, _ := reader.ReadString('\n')
	inputInteger1, err := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)

	if err != nil {
		panic(err)
	}

	fmt.Println("Please, input value 2")
	input2, _ := reader.ReadString('\n')
	inputInteger2, err := strconv.ParseInt(strings.TrimSpace(input2), 10, 64)

	if err != nil {
		panic(err)
	}

	fmt.Println("sum is ", inputInteger1+inputInteger2)
}
