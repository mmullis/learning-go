// See: http://www.sarathlakshman.com/2016/06/15/pitfall-of-golang-scheduler

// Usage: GOMAXPROCS=8 go run main-v2.go

// This version only works if Goshced is called during the for loop
// It's unclear if the article's uthor left in the change
// for  threads := runtime.GOMAXPROCS(0)-1
// which does allow for main to complete without Gosched

package main
import "fmt"
import "time"
import "runtime"

func main() {
    var x int
    threads := runtime.GOMAXPROCS(0)
    for i := 0; i < threads; i++ {
	fmt.Println("creating goroutine")
        go func() {
	    fmt.Println("Entering for loop")
            for { runtime.Gosched(); x++ }
	    fmt.Println("goroutine completed")
        }()
    }
    fmt.Println("Continuing with mainline")
    runtime.Gosched()
    fmt.Println("Finished call to scheduler")
    time.Sleep(time.Second)
    fmt.Println("x =", x)
}

