package week2

import "fmt"

func sum(numbers []interface{}) interface{} {

	var s float32
	for _, n := range numbers {
		switch n.(type) {
		case int:
			s += float32(n.(int))
		case float32:
			s += n.(float32)
		default:
			fmt.Print("Invalid type")
			continue
		}
	}

	return s
}
func sum2[T int | float32](numbers []T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}
