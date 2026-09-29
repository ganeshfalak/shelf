package main

import (
	"errors"
	"fmt"
)

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

func validate(b Book) error {
	if b.Title == "" {
		return errors.New("title is required")
	}
	if b.Author == "" {
		return errors.New("author is required")
	}
	if b.Pages <= 0 {
		return fmt.Errorf("pages must be positive, got %d", b.Pages)
	}

	switch b.Status {
	case Want, Reading, Finished:
		return nil
	default:
		return fmt.Errorf("Invalid status: %s", b.Status)
	}
}

func newBook(title, author string, pages int, status Status) (Book, error) {
	b := Book{
		Title:  title,
		Author: author,
		Pages:  pages,
		Status: status,
	}

	if err := validate(b); err != nil {
		return Book{}, err
	}
	return b, nil
}

func main() {
	var empty Book

	if err := validate(empty); err != nil {
		fmt.Println("empty: ", err)
	}

	if _, err := newBook("Dune", "Frank Herbert", 0, Want); err != nil {
		fmt.Println("dune: ", err)
	}

	b, err := newBook("The Hobbit", "J.R.R. Tolkien", 310, Reading)
	if err != nil {
		fmt.Println("hobbit: ", err)
	}

	fmt.Printf("%s by %s (%d pages) [%s]\n", b.Title, b.Author, b.Pages, b.Status)
}
