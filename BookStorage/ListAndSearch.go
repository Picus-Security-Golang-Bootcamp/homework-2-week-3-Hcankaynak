package BookStorage

import (
	"fmt"
	"strings"
)

// Below functions reused from previous assignment.

// ListAllTheBooks printing book list
func ListAllTheBooks(books []Book) {
	for _, book := range books {
		fmt.Println(book.Name)
	}
}

// SearchBook Search given prefix book name in book storage.
func SearchBook(searchBookName string) {
	searchedBooks := make([]Book, 0)
	for _, book := range BookList {
		if strings.Contains(strings.ToLower(book.Name), strings.ToLower(searchBookName)) {
			searchedBooks = append(searchedBooks, book)
		}
	}
	ListAllTheBooks(searchedBooks)
}
