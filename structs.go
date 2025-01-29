package main

import "fmt"


type Books struct {
	Title string
	Author string 
	Pages int 
	IsCheckedOut bool
}

func (b *Books) toggleCheckOut(){
	b.IsCheckedOut = !b.IsCheckedOut
}

func findLongestBook(books []Books) Books {
	mostPages := 0
	var longestBook Books 

	for _, b := range books {
		if b.Pages > mostPages{
			mostPages = b.Pages
			longestBook = b 
		}
	}
	return longestBook
}  

func main(){

books := []Books{
	{
		Title: "Join the War",
		Author: "PP Wario",
		Pages: 269,
		IsCheckedOut: false,
	}, 
	{
		Title: "Winner Winner Chicken Dinner",
		Author: "W.W Chic",
		Pages: 500,
		IsCheckedOut: true,
	}, 
	{
		Title: "Here me, here ye!",
		Author: "Ching Ming",
		Pages: 304,
		IsCheckedOut: true,
	},
}

books[0].toggleCheckOut() 
books[1].toggleCheckOut()
fmt.Println(books)

longestBook := findLongestBook(books)
fmt.Println(longestBook)
}




