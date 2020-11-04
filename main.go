package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://fantasy.premierleague.com/",
		"http://google.com/",
		"https://www.facebook.com/",
		"https://golang.org/",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error: ", err)
		c <- "might be down"
		return
	}
	fmt.Println(link, "is up")
	c <- "yup its up"
}
