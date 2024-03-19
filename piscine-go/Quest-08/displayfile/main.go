package main

import (
	"fmt"
	"os"
)

func main() {
	var fileName []string = os.Args[1:]
	if len(fileName) > 1 {
		fmt.Println("Too many arguments")
	} else if len(fileName) < 1 {
		fmt.Println("File name missing")
	} else {
		file, err := os.Open(fileName[0])
		if err != nil {
			fmt.Println(err)
		}
		arr := make([]byte, 14)
		file.Read(arr)
		fmt.Println(string(arr))
		file.Close()
	}
}
