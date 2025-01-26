class Movie {
    constructor(title, director, genre, availableCopies) {
        this.title = title;
        this.director = director;
        this.genre = genre;
        this.availableCopies = availableCopies;
    }
}

class Customer {
    constructor(customerName, moviesRented = []) {
        this.customerName = customerName;
        this.moviesRented = moviesRented;
    }
}

class RentalStore {
    constructor(storeName) {
        this.storeName = storeName;
        this.movies = new Map();
        this.customers = new Map();
    }

    addMovies(movie){
        if(!this.movies.has(movie.title)){
            this.movies.set(movie.title, movie);
            console.log(`Movie Added: ${movie.title}`);
        }
    }

    registerCustomer(customer) {
        if(!this.customers.has(customer)){
            this.customers.set(customer.customerName, customer);
            console.log(`Customer Registered: ${customer.customerName}, Movies Rented: ${customer.moviesRented}`);
        }
    }
}

const store = new RentalStore("Movies R Us")

const movie1 = new Movie("Into Pandemonium",
		"Mike Will Made",
		"Fantasy",
		3,)

const movie2 = new Movie("Who's knocking?",
    "Jody Flacko",
    "Thriller",
     2,)

const customer1 = new Customer("Daryl");
const customer2 = new Customer("Rae");

store.addMovies(movie1);
store.addMovies(movie2);
store.registerCustomer(customer1);
store.registerCustomer(customer2);

console.log(`Rental Store: ${store.storeName}`);
console.log(`Available Movies:`);
for (const [key, m] of store.movies){
    console.log(`Title: ${m.title}, Directed By ${m.director}, Genre: ${m.genre}, Copies in store: ${m.availableCopies}`);
}
console.log(`Store Customers:`)
for (const [key, c] of store.customers){
    console.log(`Customer: ${c.customerName}, Movies Rented: ${c.moviesRented}`);
}
