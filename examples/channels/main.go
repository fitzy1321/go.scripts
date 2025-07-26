package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// create a channel. This should be json strings
	c_json_str := make(chan string)

	// goroutine, call websites and parse stuff
	go callWebsites(c_json_str)

	// iterating over a channel, will wait for data then iterate
	// channel must be closed, to prevent infinite deadlock after goroutine ends
	for result := range c_json_str {
		// print json string to screen
		fmt.Println("Returned from call:", result)
	}
}

// An object to turn into a json string
type Response struct {
	SiteName   string `json:"site"`
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func callWebsites(c chan string) {
	// defer will call close once the function ends
	defer close(c)

	// simulating of hitting websites and getting a response
	results := []Response{
		{SiteName: "google.com", StatusCode: 200, Message: ""},
		{SiteName: "amazon.com", StatusCode: 200, Message: "Ready to buy something?"},
		{SiteName: "some_nonsense.com", StatusCode: 404, Message: "Site not found!"},
	}

	// turning objects into json strings
	// errors witll be logged and parsed to error strings
	for _, r := range results {
		j, err := json.Marshal(r)
		if err != nil {
			log.Printf("Error occured parsing objects to json str: %s", err)
			c <- fmt.Sprintf(`{"error": "%s"}`, err)
		}
		c <- string(j)
	}

}
