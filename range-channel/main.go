package main

import "fmt"

func main() {
	c := make(chan int) // UNBUFFERED CHANNEL

	// THIS ANONYMOUS FUNC EXECUTES IN A NEW THREAD. AND MAIN FUNC GOES TO LINE 17.
	// AND WAITS THERE TILL IT GETS A VALUE FROM 'c'
	go func() {
		for i := 0; i < 10; i++ {
			c <- i // PASS A VALUE AND WAITS UNTIL LINE 17 READY TO RECIVE NEXT VALUE.
		}
		close(c) // ONCE DONE, CLOSE THE CHANNEL
	}()

	for n := range c {
		fmt.Println(n)
	}
}
