package main

import (
	"fmt"
)


type Author struct{
	Name string
	Nationality string
}

type Book struct{
	Title string
	Pages int
	Author Author 
}

func createBook (title string, pages int, name string, nationality string) Book {
	book := Book{
		Title:  title,
		Pages:  pages,
		Author: Author{
			Name:        name,
			Nationality: nationality,
		},
	}
	return book
}

func printBook(book Book){
	fmt.Printf("Book(Title: %s, Pages: %d, Author: [Name: %s, Nationality: %s])\n", 
book.Title, 
book.Pages, 
book.Author.Name,
book.Author.Nationality)
}

func main(){
b := createBook("Daryl's World", 369, "Daryl Peterson", "American")
printBook(b)
}