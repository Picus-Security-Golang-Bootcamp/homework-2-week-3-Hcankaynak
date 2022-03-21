package main

import (
	"fmt"
	"main/BookStorage"
	"os"
	"strings"
)

func main() {
	// init books
	BookStorage.InitBooks()

	// handle arguments
	args := os.Args[1:]
	handleArguments(args)
}

func handleArguments(args []string) {

	if args[0] == "list" {
		BookStorage.ListAllTheBooks(BookStorage.BookList)
	} else if args[0] == "search" {
		bookName := strings.Join(args[1:], " ")
		BookStorage.SearchBook(bookName)
	} else if args[0] == "get" {
		BookStorage.GetBookById(args[1])
	} else if args[0] == "delete" {
		BookStorage.DeleteById(args[1])
	} else if args[0] == "buy" {
		BookStorage.BuyById(args[1], args[2])
	} else {
		// usage
		fmt.Println("Wrong command, please try again as below;")
		fmt.Println("List : go run main.go list")
		fmt.Println("Search: go run main.go search <bookName>")
		fmt.Println("Get: go run main.go get <bookID>")
		fmt.Println("Delete: go run main.go delete <bookID>")
		fmt.Println("Buy: go run main.go buy <bookID> <quantity>")
	}
}
