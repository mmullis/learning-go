package main

import (
	"fmt"
	"runtime"
	"sync"
	//	"time"
)

// Use a goroutine for every letter/character
// Parallel example using GOMAXPROCS > 1
func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	fmt.Println("Starting")
	for char := 'a'; char < 'a'+26; char++ {
		wg.Add(1)
		go func(c rune) {
			defer wg.Done()
			fmt.Printf("char: %c  V: %V T: %T\n", c, c, c)
		}(char)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	//	runtime.Gosched()
	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
