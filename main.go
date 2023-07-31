package main

import (
	"fmt"
	"strings"
)

func main() {

	//fmt.Print("Hello World")

	lines := strings.Split("aSDASDASD,badada asdsda ,csdasdasda asdasd asdasda ", ",")
	for i, line := range lines {
		fmt.Println(i, line)
		wc := len(strings.Fields(line))
		fmt.Println(wc)
	}
}
