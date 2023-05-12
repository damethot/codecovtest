package main

import "fmt"

func main() {
	hello(true)
}

func hello(b bool) {
	if b {
		fmt.Println("Hello, world!")
	} else {
		fmt.Println("Goodbye, world!")
	}
}
