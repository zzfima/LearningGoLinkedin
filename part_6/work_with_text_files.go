package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main17() {
	defer ReadFile()
	WriteToFile()
}

// ReadFile read and show file contain
func ReadFile() {
	bytes, err := ioutil.ReadFile("someText.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("File length is:", len(bytes))
		fmt.Println("File text:", string(bytes))
	}
}

// WriteToFile fill file with some data
func WriteToFile() {
	fileHandler, err := os.Create("someText.txt")
	defer fileHandler.Close()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		io.WriteString(fileHandler, time.Now().String())
	}
}
