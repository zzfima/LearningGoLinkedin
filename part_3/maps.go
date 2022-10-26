package main

import (
	"fmt"
	"sort"
)

func main_10() {
	states := make(map[string]string)
	states["start"] = "success"
	states["finish"] = "uncompleted"
	states["current"] = "processing"
	states["panic"] = "exception"
	fmt.Println(states)

	delete(states, "finish")
	fmt.Println(states)

	for k, v := range states {
		fmt.Println(k, ":", v)
	}

	length := len(states)
	process_state := make([]string, length)
	i := 0
	for _, v := range states {
		process_state[i] = v
		i++
	}

	sort.Strings(process_state)
	fmt.Println(process_state)
}
