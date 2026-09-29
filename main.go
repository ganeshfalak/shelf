package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	var empty Book

	fmt.Printf("empty: %+v\n", empty)

	b := Book{
		Title:  "The Hobbit",
		Author: "J.R.R. Tolkien",
		Pages:  310,
	}

	fmt.Printf("%s by %s (%d pages)\n", b.Title, b.Author, b.Pages)
}
