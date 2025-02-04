package main

import "fmt"

type Books struct {
	Title string
	Author string
	Pages int 
	isCheckedOut bool
}

type Library struct{
	Books []Books
}

func (b *Books) toggleCheckOut(){
	b.isCheckedOut = !b.isCheckedOut;
}

func (l *Library) addBooks(newBook Books){
	l.Books = append(l.Books, newBook)
}

func (l *Library) removeBooksByTitle(title string) []Books {
	var newLibrary []Books

	for _, book := range l.Books {
		if book.Title != title{
			newLibrary = append(newLibrary, book)
		}
	}
	return newLibrary
}

func (l *Library) getBooksByAuthor(author string) []Books {
	var booksByAuthor []Books

	for _, book := range l.Books {
		if book.Author == author {
			booksByAuthor = append(booksByAuthor, book)
		}
	}
	return booksByAuthor
}

func (l *Library) getTotalPages() int {
totalPages := 0

for _, book := range l.Books{
	totalPages += book.Pages
}
return totalPages
}

func displayLibrary(books []Books){
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Pages: %d, Checked Out: %v\n", book.Title, book.Author, book.Pages, book.isCheckedOut)
	}
}


func main(){

books := []Books{
	{
		Title:        "The Night Ends With Fire",
		Author:       "K. X. Song",
		Pages:        410,
		isCheckedOut: false,
	}, 
	{
		Title:        "Four Ruined Realms",
		Author:       "Mai Corland",
		Pages:        413,
		isCheckedOut: false,
	}, 
	{
		Title:        "The Way of Kings",
		Author:       "Brandon Sanderson",
		Pages:        1007,
		isCheckedOut: false,
	}, {
		Title:        "Words of Radiance",
		Author:       "Brandon Sanderson",
		Pages:        1088,
		isCheckedOut: false,
	}, 
}

library := Library{}

book1 := Books{"Five Broken Blades", "Mai Corland", 410, false}

library.addBooks(books[0])
library.addBooks(books[1])
library.addBooks(books[2])
library.addBooks(books[3])
library.addBooks(book1)

library.Books[0].toggleCheckOut()
library.Books[1].toggleCheckOut()

library.Books = library.removeBooksByTitle("Words of Radiance")

//byAuthor := library.getBooksByAuthor("Mai Corland")

displayLibrary(library.Books)

totalPages := library.getTotalPages()
fmt.Println(totalPages)


}