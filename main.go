package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	links := []string{
		"https://fantasy.premierleague.com/",
		"http://google.com/",
		"https://www.facebook.com/",
		"https://golang.org/",
	}
	for _, link := range links {
		checkLink(link)
	}
}
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(link, "is up")
}
