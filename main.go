package main

import (
	"fmt"
	"os"
	"strings"
)

var counter int = 0

func main() {
	argsWithoutProg := os.Args[1:]
	s := strings.Split(argsWithoutProg[0], " ")
	for range s {
		counter += 1
	}
	fmt.Println(counter)
}
