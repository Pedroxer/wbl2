package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	or := make(chan interface{})
	done := make(chan interface{}, len(channels)-1)
	for _, c := range channels {
		go func(c <-chan interface{}) {
			done <- <-c
		}(c)
	}
	go func() {
		<-done
		close(or)
	}()
	return or
}

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
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))

}
