package main

import "fmt"

type Status int

const (
	Want Status = iota
	Reading
	Finished
)

func (s Status) String() string {
	switch s {
	case Want:
		return "want"
	case Reading:
		return "reading"
	case Finished:
		return "finished"
	default:
		return "unknown"
	}
}

type Book struct {
	Title  string
	Author string
	Pages  int
	Status Status
}

func main() {
	var empty Book

	fmt.Printf("empty: %+v\n", empty)

	b := Book{
		Title:  "The Hobbit",
		Author: "J.R.R. Tolkien",
		Pages:  310,
		Status: Reading,
	}

	fmt.Printf("%s by %s (%d pages) [%s]\n", b.Title, b.Author, b.Pages, b.Status)
}
