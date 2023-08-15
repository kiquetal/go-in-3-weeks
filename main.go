package main

import (
	"go-in-3-weeks/stringutils"
	"go-in-3-weeks/week2"
)

func main() {

	//fmt.Print("Hello World")

	print(week2.Another())
	/*	lines := strings.Split("aSDASDASD,badada asdsda ,csdasdasda asdasd asdasda ", ",")
		for i, line := range lines {
			fmt.Println(i, line)
			wc := len(strings.Fields(line))
			fmt.Println(wc)
		}

	*/
	stringutils.Show()
}
