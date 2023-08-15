package exercises

import "fmt"

type Person struct {
	Name    string
	address string
	city    string
	state   string
	zip     string
}

func (p Person) String() string {
	return fmt.Sprintf("%s\n%s, %s %s", p.Name, p.address, p.city, p.state, p.zip)
}
