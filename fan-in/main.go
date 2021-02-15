package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(gatherInfo("Top Secret"), gatherInfo("Mission"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("End info gathering")
}

func gatherInfo(msg string) <-chan string {
	out := make(chan string)

	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return out
}

// FAN-IN PATTERN : MULTIPLE CHANNELS INPUT INTO A SINGLE CHANNEL
func fanIn(input1, input2 <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for {
			out <- <-input1
		}
	}()

	go func() {
		for {
			out <- <-input2
		}
	}()

	return out
}
