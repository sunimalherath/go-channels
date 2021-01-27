package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		// IF wg.Add(1) IS HERE, THEN A RACE CONDITION.
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		// IF wg.Add(1) IS HERE, THEN A RACE CONDITION. COZ BOTH GO ROUTINES TRYING
		// WRITE TO THE SAME VARIABLE - A BIG NO NO IN GO.
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
