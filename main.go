package main

// * .vscode/settings.json unusedfund set to false

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime/debug"
	"sync"

	. "go.scripts/result"
)

func playingWithChannels() {
	type (
		payload struct {
			Meaning int
		}
		rpAlias = Result[payload]
	)
	someChan := make(chan rpAlias, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan rpAlias) {
		defer wg.Done()
		ch <- Ok(payload{42})
	}(someChan)
	go func(ch chan rpAlias) {
		defer wg.Done()
		ch <- Err[payload](errors.New("Something broke, loser..."))
	}(someChan)
	wg.Wait()
	close(someChan)

	for mResult := range someChan {
		if mResult.IsErr() {
			fmt.Fprintf(os.Stderr, "Error from goroutine: %+v\n", mResult)
		}
		if mResult.IsOk() {
			fmt.Printf("From my goroutine: %+v\n", mResult)
		}
	}
}

func referenceToSliceProblems() {
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

func whatsYourName() {
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		fmt.Println("Hello", name)
	} else {
		log.Fatal(err)
	}
}
func gorilla() {
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

func appName() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "badman lol ..."
	}
	return info.Main.Path
}

func main() {
	// v := 42
	// if ok := varIsString(v); !ok {
	// 	fmt.Println("variable 'v' is a string!")
	// } else {
	// 	fmt.Println("variable 'v' is not a string")
	// }

	// HELLO_GO()
	// go GORILLA()
	// WHATS_YOUR_NAME()
	// SLICE_REFERENCE_BUG_EXAMPLE()
	playingWithChannels()
	fmt.Println(appName())
	// fmt.Println("Num args:", len(os.Args), os.Args)
	some := make(map[string]any)
	fmt.Printf("Some map: %v %v", some, some == nil)
	for k, v := range some {
		fmt.Println(k, v)
	}
}
