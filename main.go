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

type Book struct {
	Title  string
	Author string
	Pages  int
	Status Status
}

type Shelf struct {
	books []Book
}

func (s *Shelf) add(b Book) error {
	if err := validate(b); err != nil {
		return err
	}

	s.books = append(s.books, b)
	return nil
}

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
	var shelf Shelf

	if err := shelf.add(Book{}); err != nil {
		fmt.Println("empty: ", err)
	}

	dune, err := newBook("Dune", "Frank Herbert", 412, Want)
	if err != nil {
		fmt.Println("dune: ", err)
		return
	}

	if err := shelf.add(dune); err != nil {
		fmt.Println("add dune: ", err)
		return
	}

	hobbit, err := newBook("The Hobbit", "J.R.R. Tolkien", 310, Reading)
	if err != nil {
		fmt.Println("hobbit: ", err)
	}

	if err := shelf.add(hobbit); err != nil {
		fmt.Println("add hobbit: ", err)
		return
	}

	fmt.Printf("shelf has %d books\n", len(shelf.books))
	for i, b := range shelf.books {
		fmt.Printf("%d %s by %s (%d pages) [%s]\n", i+1, b.Title, b.Author, b.Pages, b.Status)
	}
}
