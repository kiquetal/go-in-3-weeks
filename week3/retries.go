package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

}

func getWithBackoff(url string, count int) (successCount int, throttleCount int, err error) {

	var backoffSchedule = []time.Duration{
		100 * time.Millisecond,
		200 * time.Millisecond,
		400 * time.Millisecond,
		800 * time.Millisecond,
		1000 * time.Millisecond,
	}

	for i := 0; i < count; i++ {

		var res *http.Response
		for _, backoff := range backoffSchedule {
			var req *http.Request
			req, err = http.NewRequest("GET", url, nill)
			if err != nil {
				return
			}
			req.Header.Set("X-Request-ID", strconv.Itoa(i))
			client := &http.Client{}
			res, err = client.Do(req)
			if err != nil {
				return
			}
			if res.StatusCode == http.StatusOK {
				successCount++
				break
			}
			if res.StatusCode == http.StatusTooManyRequests {
				throttleCount++

			}
			log.Printf("Got status %d, retrying in %s", res.StatusCode, backoff)
			time.Sleep(backoff)
		}
		res.Body.Close()
		log.Printf("%s: %d", url, res.Status)

	}
	return
}
