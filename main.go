package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func HELLO_GO() {
	fmt.Println("Hello Go!")
}
func WHATS_YOUR_NAME() {
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		fmt.Println("Hello", name)
	} else {
		log.Fatal(err)
	}
}
func GORILLA() {
	gorilla := '🦍'
	fmt.Printf("gorilla: %v %v", gorilla, reflect.TypeOf(gorilla))
}

func main() {
	HELLO_GO()
	go GORILLA()
	WHATS_YOUR_NAME()
}
