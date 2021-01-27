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

	<-done // WAITING TILL A VALUE COMES FROM THE done CHANNEL - BLOCKS THE EXECUTION.
	// SINCE c IS A UNBUFFERED CHANNEL, WHEN A VALUE ADDED TO IT THERE SHOULD BE A CODE
	// TO RETREIVE IT AT THE SAME TIME (LIKE PASSING BATTON IN RELAY RACE)
	// SINCE RECEIVING CODE IS BELOW THE COMPILER WON'T REACH IT BECAUSE MAIN
	// IS WAITING FOR A VALUE FROM THE done CHANNEL.
	<-done
	close(c)

	for n := range c {
		fmt.Println(n)
	}
}
