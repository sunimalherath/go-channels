package main

import "fmt"

func main() {
	c := make(chan int) // UNBUFFERED CHANNEL

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // KEEP ADDING VALUES TO THE CHANNEL
		}
		close(c) // ONCE DONE, CLOSE THE CHANNEL
	}()

	for n := range c {
		fmt.Println(n)
	}
}
