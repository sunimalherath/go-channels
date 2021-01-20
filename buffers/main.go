package main

import "fmt"

func main() {
	c := make(chan int, 2) // channel that can fit 2 integers

	c <- 12
	c <- 34

	fmt.Println(<-c) // 12
	fmt.Println(<-c) // 34

	//fmt.Println(<-c) // deadlock. We already took the 2 values from the channel.
}
