package main

import (
	"bufio"
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

// A bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookworms reads the file and returns the list of bookworms,
// and their beloved books, found therein.
func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Set our own buff size to reduce the systems calls to read the file.
	buffedReader := bufio.NewReaderSize(f, 1024*1024)

	// Initialize the type in which the file will be decoded.
	var bookworms []Bookworm

	// Decode the file and store the content in the variable bookworms.
	err = json.NewDecoder(buffedReader).Decode(&bookworms)
	if err != nil {
		return nil, err
	}
	return bookworms, nil
}

// bookCount register all the books and their occurrences
// from the bookworms shelves.
func bookCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)
	seen := make(map[Book]bool)

	for _, bookworm := range bookworms {
		clear(seen) // O(1) - Limpieza optimizada por el runtime

		for _, book := range bookworm.Books {
			if seen[book] {
				continue
			}
			seen[book] = true
			count[book]++
		}
	}
	return count
}

// findCommonBooks returns books that are on more than one bookworm's shelf.
func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelf := bookCount(bookworms)

	var commonBooks []Book

	for book, count := range booksOnShelf {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	return sortBooks(commonBooks)
}

// sortBooks sorts the books by Author and then Title.
func sortBooks(books []Book) []Book {
	slices.SortFunc(books, func(a, b Book) int {
		if a.Author != b.Author {
			return cmp.Compare(a.Author, b.Author)
		}
		return cmp.Compare(a.Title, b.Title)
	})
	return books
}

// displayBooks prints out the titles and authors of a list of books
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
