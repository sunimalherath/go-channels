package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done   // WAITING TILL A VALUE COMES FROM THE done CHANNEL
		<-done   // WAITING TILL THE SECODE VALUE COMES FROM THE CHANNEL
		close(c) // CLOSE THE CHANNEL
	}()

	for n := range c {
		fmt.Println(n)
	}
}
