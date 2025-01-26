package main

import "fmt"

type Movie struct {
	Title string
	Director string
	Genre string 
	AvailableCopies int
}

type Customer struct {
	CustomersName string
	MoviesRented []string
}

type RentalStore struct {
	StoreName string
	Movies map[string]Movie
	Customers map[string]Customer
}

func (rs *RentalStore) addMovie(movie Movie){
	if rs.Movies == nil { // why am i checking if this is nil?
		rs.Movies = make(map[string]Movie) // why am i creating a map here when i already created a map for this inside of the RentalStore Struct?
	}
	rs.Movies[movie.Title] = movie 
	fmt.Println("Movie Added:", movie.Title)
}

func (rs *RentalStore) registerCustomer(name string){
	if rs.Customers == nil {  // why am i checking if this is nil?
		rs.Customers = make(map[string]Customer)  // why am i creating a map here when i already created a map for this inside of the RentalStore Struct?
	}
	rs.Customers[name] = Customer{
		CustomersName: name,
		MoviesRented:  []string{},
	}
	fmt.Println("Customer Registered:", name)
}

func main(){
	store := RentalStore{
		StoreName: "Movies R Us", 
	}

	movie1 := Movie{
		Title: "Into Pandemonium",
		Director: "Mike Will Made",
		Genre: "Fantasy",
		AvailableCopies: 3,
	}

	movie2 := Movie{
		Title: "Who's knocking?",
		Director: "Jody Flacko",
		Genre: "Thriller",
		AvailableCopies: 2,
	}

	store.addMovie(movie1)
	store.addMovie(movie2)
	store.registerCustomer("Daryl")
	store.registerCustomer("Rae")

	fmt.Println("Rental Store:", store.StoreName)
	fmt.Println("Available Movies:")
	for _, m := range store.Movies{
		fmt.Printf("Title: %s, Directed By: %s, Genre: %s, Copies in Store: %d\n", m.Title, m.Director, m.Genre, m.AvailableCopies)
	}
	fmt.Println("Store Customers:")
	for _, c := range store.Customers{
		fmt.Printf("Name: %s, Rented Movies: %v\n", c.CustomersName, c.MoviesRented)
	}
}




