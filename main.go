package main

import (
	"fmt"
	"os"
	"strings"
)

var counter int = 0

func main() {
	argsWithoutProg := os.Args[1:]
	s := strings.Fields(string(argsWithoutProg[0]))
	fmt.Println(len(s))
}
