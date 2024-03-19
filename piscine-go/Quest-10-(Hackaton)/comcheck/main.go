package main

import (
	"fmt"
	"os"
)

func main() {
	var result bool
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			result = true
		}
	}
	if result {
		fmt.Println("Alert!!!")
	}
}
