package main

import "fmt"

func main() {
	c := make(chan int) // no buffer size given

	go func() {
		c <- 23
		c <- 45
		c <- 34
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//fmt.Println(<-c) // deadlock for this one as we already got values from the channel.
}
