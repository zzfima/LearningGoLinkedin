package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main_3() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please, input integer")
	input, _ := reader.ReadString('\n')
	inputInteger, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Type is %T\n", inputInteger)
	}

	var i int8 = 120
	i += 10
	fmt.Println(i)
}
