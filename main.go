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

func varIsString(v any) bool {
	// .(type) is the conversion operator, returns the value and a bool
	// check if the bool 'ok' is true
	// if x, ok := v.(string); ok {
	// 	fmt.Println(x)
	// }
	_, ok := v.(string)
	return ok
}

func main() {
	v := 42
	if ok := varIsString(v); !ok {
		fmt.Println("variable 'v' is a string!")
	} else {
		fmt.Println("variable 'v' is not a string")
	}

	HELLO_GO()
	go GORILLA()
	WHATS_YOUR_NAME()
}
