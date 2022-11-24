package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var words string

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	s := strings.Fields(text)
	fmt.Println(len(s))

}
