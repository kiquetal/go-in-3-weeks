package main

import (
	"fmt"
	"go-in-3-weeks/week2"
	"strings"
)

func main() {

	//fmt.Print("Hello World")

	print(week2.Another())
	lines := strings.Split("aSDASDASD,badada asdsda ,csdasdasda asdasd asdasda ", ",")
	for i, line := range lines {
		fmt.Println(i, line)
		wc := len(strings.Fields(line))
		fmt.Println(wc)
	}
}
