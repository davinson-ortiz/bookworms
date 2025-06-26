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

// set es un tipo para manejar conjuntos de libros de forma eficiente
type set map[Book]struct{}

// Contains verifica si un libro está en el conjunto
func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}

// Add añade un libro al conjunto
func (s set) Add(b Book) {
	s[b] = struct{}{}
}

// bookRecommendations mapea cada libro a otros libros "similares" con sus puntuaciones
type bookRecommendations map[Book]map[Book]float64

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

// listOtherBooksOnShelves devuelve todos los libros EXCEPTO el del índice dado
func listOtherBooksOnShelves(excludeIndex int, books []Book) []Book {
	var otherBooks []Book
	for i, book := range books {
		if i != excludeIndex {
			otherBooks = append(otherBooks, book)
		}
	}
	return otherBooks
}

// registerBookRecommendations registra las asociaciones entre un libro y otros libros
// que aparecen en la misma estantería (indicating similarity)
func registerBookRecommendations(sb bookRecommendations, targetBook Book, otherBooks []Book) {
	// Inicializar el mapa si no existe
	if sb[targetBook] == nil {
		sb[targetBook] = make(map[Book]float64)
	}

	// Por cada libro que aparece junto al libro objetivo, incrementamos su puntuación
	for _, otherBook := range otherBooks {
		sb[targetBook][otherBook] += 1.0
	}
}

// recommendBooks genera recomendaciones para un bookworm específico
func recommendBooks(sb bookRecommendations, ownedBooks []Book) []Book {
	// Crear conjunto de libros que ya posee el lector
	owned := make(set)
	for _, book := range ownedBooks {
		owned.Add(book)
	}

	// Mapa para acumular puntuaciones de recomendaciones
	recommendations := make(map[Book]float64)

	// Para cada libro que posee el lector
	for _, ownedBook := range ownedBooks {
		// Obtener libros similares y sus puntuaciones
		if similarBooks, exists := sb[ownedBook]; exists {
			for similarBook, score := range similarBooks {
				// Solo recomendar libros que no posee
				if !owned.Contains(similarBook) {
					recommendations[similarBook] += score
				}
			}
		}
	}

	// Convertir a slice y ordenar por puntuación (descendente) y luego alfabéticamente
	var recommendedBooks []Book
	for book := range recommendations {
		recommendedBooks = append(recommendedBooks, book)
	}

	// Ordenar por puntuación descendente, luego por autor y título
	slices.SortFunc(recommendedBooks, func(a, b Book) int {
		scoreA := recommendations[a]
		scoreB := recommendations[b]

		// Primero por puntuación (descendente)
		if scoreA != scoreB {
			if scoreA > scoreB {
				return -1
			}
			return 1
		}

		// Luego por autor
		if a.Author != b.Author {
			return cmp.Compare(a.Author, b.Author)
		}

		// Finalmente por título
		return cmp.Compare(a.Title, b.Title)
	})

	return recommendedBooks
}

// recommendOtherBooks implementa el algoritmo principal de recomendaciones
func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookRecommendations)

	// Registrar todas las asociaciones de libros en todas las estanterías
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
			registerBookRecommendations(sb, book, otherBooksOnShelves)
		}
	}

	// Generar recomendaciones para cada bookworm
	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		recommendations[i] = Bookworm{
			Name:  bookworm.Name,
			Books: recommendBooks(sb, bookworm.Books),
		}
	}

	return recommendations
}

// displayRecommendations muestra las recomendaciones de forma organizada
func displayRecommendations(recommendations []Bookworm) {
	for _, bookworm := range recommendations {
		fmt.Printf("\n=== Recomendaciones para %s ===\n", bookworm.Name)
		if len(bookworm.Books) == 0 {
			fmt.Println("No hay recomendaciones disponibles.")
		} else {
			displayBooks(bookworm.Books)
		}
	}
}
