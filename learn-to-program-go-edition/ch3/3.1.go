package main

import (
    "fmt"
    "strings"
)

func main() { 
	fmt.Println("I like" + "apple pie.")
	fmt.Println("I like " + "apple pie.")
	fmt.Println("I like" + " apple pie.")

	// FAIL: fmt.Println("blink " * 4)
	// Use alternative
	fmt.Println(strings.Repeat("blink ", 4))
}

