package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Info writes logs in the color blue with "INFO: " as prefix
var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)

// Warning writes logs in the color yellow with "WARNING: " as prefix
var Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)

// Error writes logs in the color red with "ERROR: " as prefix
var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)

// Debug writes logs in the color cyan with "DEBUG: " as prefix
var Debug = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)

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
			Error.Printf("Error occured parsing objects to json str: %s", err)
			c <- fmt.Sprintf(`{"error": "%s"}`, err)
		}
		c <- string(j)
	}

}
