package main

import "fmt"

func main() {
	c := incrementor()
	sum := puller(c)
	for n := range sum {
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

/* The optional <- operator specifies the channel direction (i.e: send or receive).
if no direction is given, then the channel is bidirectional.
more info: https://golang.org/ref/spec#Channel_types
*/
