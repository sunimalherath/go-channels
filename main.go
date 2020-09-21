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

	// cannot use for range to loop through as we are not going to use the link parameter inside the loop
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
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
