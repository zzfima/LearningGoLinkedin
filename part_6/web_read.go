package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main18() {
	resp, errGet := http.Get("https://agilemanifesto.org/principles.html")

	if errGet != nil {
		panic(errGet)
	} else {
		defer resp.Body.Close()
		fmt.Printf("Resp type: %T\n", resp)

		bytes, errReadAll := ioutil.ReadAll(resp.Body)
		if errReadAll != nil {
			panic(errReadAll)
		} else {
			fmt.Println(string(bytes))
		}
	}
}
