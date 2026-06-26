package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func SLICE_REFERENCE_BUG_EXAMPLE() {
	someSlice := []int{}

	for i := range 5 {
		fmt.Println(i)
		someSlice = append(someSlice, i)
	}

	x := &someSlice[3]
	someSlice = append(someSlice, 5)
	someSlice = append(someSlice, 6)
	someSlice = append(someSlice, 7)
	someSlice = append(someSlice, 8)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)
	someSlice = append(someSlice, 9)

	// Trying to trigger reallocation, to see what happens to x?

	fmt.Println("What is x now, after an allocation in the slice?", x, *x, &someSlice[3])
	// Hmmmmm, seems the GC or compiler doesn't clean up or throw an error about x
	//  but looks like the memory address has changed.
	// guess X holds a reference to that memory, but the slice has changed to a new backing array
	*x = 42
	if *x == someSlice[3] {
		fmt.Println("YOU'RE AN IDIOT")
	}
}

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
	SLICE_REFERENCE_BUG_EXAMPLE()
}
