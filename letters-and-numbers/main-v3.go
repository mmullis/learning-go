package main

import (
	"fmt"
	"runtime"
	"sync"
//	"time"
)

// Parallel example using GOMAXPROCS > 1
func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting")
	go func() {
		defer wg.Done()
		//time.Sleep(1 * time.Microsecond)
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

	go func() {
		defer wg.Done()
		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

