package BookStorage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
)

// BookList storing book list which is created from json
var BookList []Book

// GetBookById finding book by id and print it
func GetBookById(idArgument string) {
	id, _ := strconv.Atoi(idArgument)
	_, book, err := findById(id)
	if err != nil {
		fmt.Println(err)
	} else {
		printBook(book)
	}
}

// DeleteById delete book by id, before delete print it, then update book list.
func DeleteById(idArgument string) {
	id, _ := strconv.Atoi(idArgument)
	index, book, err := findById(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted Book")
		printBook(book)
		BookList = append(BookList[:index], BookList[index+1:]...)
	}
}

// BuyById Buying book by id if there is a book that id and quantity satisfied.
func BuyById(idArgument string, quantityArgument string) {
	id, _ := strconv.Atoi(idArgument)
	quantity, _ := strconv.Atoi(quantityArgument)
	index, book, err := findById(id)
	if err != nil {
		fmt.Println(err)
	} else {
		if book.Stock > quantity {
			book.Stock = book.Stock - quantity
		}
		printBook(book)
		BookList[index] = book
		fmt.Println(BookList)
	}
}

// findById private function for finding by id, it used multiple times. return error if no book found.
func findById(id int) (int, Book, error) {
	for index, book := range BookList {
		if book.Id == id {
			return index, book, nil
		}
	}
	errorMessage := "Couldn't find a book by id:" + strconv.Itoa(id)
	return -1, *new(Book), errors.New(errorMessage)
}

// printBook simple printing function for book.
func printBook(book Book) {
	fmt.Println("id:", book.Id)
	fmt.Println("name:", book.Name)
	fmt.Println("page number:", book.PageNumber)
	fmt.Println("stock:", book.Stock)
	fmt.Println("price:", book.Price)
	fmt.Println("stock code:", book.StockCode)
	fmt.Println("isbn:", book.ISBN)
	fmt.Println("author id:", book.Author.Id)
	fmt.Println("author name:", book.Author.Name)

}

// InitBooks initialize books from json.
func InitBooks() {
	var initialBookStorage []Book
	contents, err := ioutil.ReadFile("BooksDB.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(contents, &initialBookStorage); err != nil {
		panic(err)
	}
	BookList = initialBookStorage
}
