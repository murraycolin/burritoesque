package main

import (
	"log"
	"math"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	request, err := http.NewRequest("GET", "https://www.murraycolin.com/", nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0")
	t := time.Now().Hour()
	tf := float64(t)
	index := math.Sin(tf/math.Pi) + 1
	for index < 2 {
		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()
		index = index + .005
	}
}
