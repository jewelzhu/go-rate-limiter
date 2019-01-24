package go_rate_limiter


import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestRateLimit(T *testing.T) {
	l := NewRateLimiter(10, 5)
	c := make(chan int, 1)
	go func() {
		for {
			i := <- c
			delay, ok := l.incoming()
			if !ok {
				fmt.Println("-")
			} else if delay > 0 {
				log.Printf("accept %d but need sleep %f seconds %s\n", i, delay, time.Now().Format("2016-01-02 15:04:05.000"))
			} else {
				log.Printf("accept %d at once %s", i, time.Now().Format("2016-01-02 15:04:05.000"))
			}
		}
	}()

	for i:=0;i<20;i++{
		time.Sleep(time.Millisecond * 100)
		c <- i
	}

	for i:=0;i<20;i++{
		time.Sleep(time.Millisecond * 50)
		c <- i
	}

	for i:=0;i<20;i++{
		time.Sleep(time.Millisecond * 20)
		c <- i
	}
}

