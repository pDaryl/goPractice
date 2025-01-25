package main

import (
	"fmt"
)


type Author struct{
	Name string 
	Nationality string
}

type Genre struct{
	Name string
}

type Book struct{
	Title string
	Pages int
	Author Author 
	Genre Genre
}

type Library struct{
	Name string
	Location string
	Book []Book
	Genre []Genre
}

func createAuthor(name string, nationality string) Author {
	author := Author{
		Name:        name,
		Nationality: nationality,
	}
	return author
}

func createGenre(name string) Genre {
	genre := Genre{
		Name: name,
	}
	return genre
}

func createBook(title string, pages int, author Author, genre Genre) Book {
	book := Book{
		Title:  title,
		Pages:  pages,
		Author: author,
		Genre:  genre,
	}
	return book
}

func createLibrary(name string, location string) Library{
	library := Library{
		Name:     name,
		Location: location,
		Book:     []Book{},
		Genre:    []Genre{},
	}
	return library
}

func (l *Library) addBook(book Book){ // what is going on in here? i know i am passing a pointer to library but why?
	l.Book = append(l.Book, book) // why am i appending the book here?
}

func (l *Library) addGenre(genre Genre){ // what is going on here? i know i am passing a pointer to library but why? 
	l.Genre = append(l.Genre, genre) // why am i appending the genre here?
}

func main(){
author1 := createAuthor("daryl", "USA")
author2 := createAuthor("peterson", "USA")

genre1 := createGenre("Fantasy")
genre2 := createGenre("Action")

book1 := createBook("D's World", 306, author1, genre1)
book2 := createBook("P's World", 406, author2, genre2)

library := createLibrary("New Orlean's Public Library", "New Orleans")

library.addBook(book1)
library.addBook(book2)
library.addGenre(genre1)
library.addGenre(genre2)

fmt.Printf("Library(Name: %s, Location: %s)\n", library.Name, library.Location)
fmt.Println("Books in Library:")
for _, books := range library.Book{
	fmt.Printf("- %s by %s\n", books.Title, books.Author.Name)
}
fmt.Println("Genre's in Library:")
for _, genre := range library.Genre {
	fmt.Printf("- %s\n", genre.Name)
}

}