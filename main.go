package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://apple.com",
		"http://google.com",
		"http://microsoft.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "seems to be down")
		c <- "Might be down I guess"
		return
	}

	fmt.Println(link, "is up and running!")
	c <- "Yep, it's working"
}
