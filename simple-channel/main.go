package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // PASSING THE BATTON IN A RELAY RACE. PASSES AND STOPS UNTIL IT IS RECIEVED.
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second) // WAIT A SECOND FOR THE GOROUTINES TO FINISH.
}
