package main

import (
	"fmt"
	"time"

	"github.com/K1la/taskL2.14/or"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or.Or(
		sig(5*time.Minute),
		sig(3*time.Second),
		sig(4*time.Second),
	)
	fmt.Printf("done after %v\n", time.Since(start))

	start = time.Now()
	<-or.Or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))

	start = time.Now()
	<-or.Or()
	fmt.Printf("done after %v\n", time.Since(start))
}
