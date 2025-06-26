package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Bookworms eating...")
	var bookwormsPath string
	flag.StringVar(&bookwormsPath, "path", "testdata/bookworms.json", "The path to bookworms json file.")
	flag.Parse()

	bookworms, err := loadBookworms(bookwormsPath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr,
			"failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	commonBooks := findCommonBooks(bookworms)

	fmt.Println("Here are books in common:")
	displayBooks(commonBooks)
}
