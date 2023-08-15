package week2

import (
	"fmt"
	"net/http"
)

type greeter struct{}

func (h *greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		return
	}
}
