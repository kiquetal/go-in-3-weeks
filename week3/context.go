package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

}

func greetHandlerCtx(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling greeting request")
	defer log.Println("Handled greeting request")

	completeAfter := time.After(5 * time.Second)
	ctx := r.Context()

	for {
		select {
		case <-completeAfter:

		}
	}

}
