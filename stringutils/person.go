package stringutils

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) setAge(age int) {
	p.age = age
}

func Show() {
	person := Person{name: "Alice", age: 30}
	fmt.Println(person)
	fmt.Println((person).age) // Output: 30

	person.setAge(40)
	fmt.Println((person).age) // Output: 40
}
