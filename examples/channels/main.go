package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	c_json_str := make(chan string)
	// resp, err := http.Get("https://google.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println()
	go something(c_json_str)
	for result := range c_json_str {
		fmt.Println("Returned from call:", result)
		fmt.Printf("Type of 'result' is %T\n", result)
	}
	// fmt.Println()
}

// const

type Response struct {
	SiteName   string `json:"site"`
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func something(c chan string) {
	defer close(c)
	results := []Response{
		{SiteName: "google.com", StatusCode: 200, Message: ""},
		{SiteName: "amazon.com", StatusCode: 200, Message: "Ready to buy something?"},
		{SiteName: "some_nonsense.com", StatusCode: 404, Message: "Site not found!"},
	}
	for _, r := range results {
		j, _ := json.Marshal(r)
		c <- string(j)
	}

}
