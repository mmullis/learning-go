// See: http://www.sarathlakshman.com/2016/06/15/pitfall-of-golang-scheduler
// This is the original version
// Usage: GOMAXPROCS=8 go run main-v1.go

package main
import "fmt"
import "time"
import "runtime"

func main() {
    var x int
    threads := runtime.GOMAXPROCS(0)
    for i := 0; i < threads; i++ {
        go func() {
            for { x++ }
        }()
    }
    time.Sleep(time.Second)
    fmt.Println("x =", x)
}

