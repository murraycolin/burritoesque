package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

func main() {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	t := time.Now().Hour()
	f := float64(t)

	index := math.Sin(f/math.Pi) + 1

	fmt.Println(index)

	request, err := http.NewRequest("GET", "https://www.murraycolin.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0")

	for index < 2 {

		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}
		index = index + .005
		fmt.Println(index)
		defer response.Body.Close()
	}
}
