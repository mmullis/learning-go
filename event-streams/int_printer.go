package main

import (
	"fmt"
	"os"
	"sync"
	"log"
)

func main() {
	asIsLog := log.New(os.Stderr, "", 0)
	detailedLog := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.LUTC|log.Lmicroseconds|log.Llongfile)

	bufferedCh := make(chan int, 3)
	go func() {
		for i := 0; i < 100; i++ {
			bufferedCh <- i
		}
		close(bufferedCh)
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range bufferedCh {
			fmt.Printf("%d\n", v)
		}
	}()

	wg.Wait()

	// https://stackoverflow.com/questions/29721449/how-can-i-print-to-stderr-in-go-without-using-log
	//fmt.Fprintln(os.Stderr, "Done")
	asIsLog.Println("Done")
	detailedLog.Println("Done")
}
