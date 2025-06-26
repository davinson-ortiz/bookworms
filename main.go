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
	fmt.Println("\n=== LIBROS COMUNES ===")
	commonBooks := findCommonBooks(bookworms)
	displayBooks(commonBooks)

	fmt.Println("\n=== RECOMENDACIONES ===")
	recommendations := recommendOtherBooks(bookworms)
	displayRecommendations(recommendations)
}
