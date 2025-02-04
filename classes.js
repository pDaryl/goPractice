class Books{
    constructor(title, author, pages, isCheckedOut) {
        this.title = title;
        this.author = author; 
        this.pages = pages;
        this.isCheckedOut = isCheckedOut;
    }
    toggleFav(){
        this.isCheckedOut = !this.isCheckedOut;
    }
}

class Library{
    constructor(){
        this.books = [];
    }

    addBook(newBook){
        this.books.push(newBook);
    }

    removeBookByTitle(title){
     return this.books = this.books.filter(book => book.title !== title);
    }

    getBooksByAuthor(author){
      return this.books.filter(book => book.author === author);
    }

    getTotalPages(){
        return this.books.reduce((total, book) => total + book.pages, 0);
    }
}

function displayLibrary(library){
    for (const book of library.books){
        console.log(`Title: ${book.title}, Author: ${book.author}, Pages ${book.pages}, isCheckedOut: ${book.isCheckedOut}`);
    }
}

function displayBooks(books){
    for(const book of books){
        console.log(`Title: ${book.title}, Author: ${book.author}, Pages ${book.pages}, isCheckedOut: ${book.isCheckedOut}`)
    }
}



const books = [
    new Books("The Night Ends With Fire", "K. X. Song", 410, false),
    new Books("Four Ruined Realms", "Mai Corland", 413, false), 
    new Books("The Way of Kings", "Brandon Sanderson", 1007, false), 
    new Books("Words of Radiance", "Brandon Sanderson", 1088, false),
]

const library = new Library();

for(const book of books){
    library.addBook(book);
}

const book1 = new Books("Five Broken Blades", "Mai Corland", 410, false);
library.addBook(book1)

library.books[0].toggleFav()

library.removeBookByTitle("Words of Radiance");

displayLibrary(library);

console.log();
displayBooks(library.getBooksByAuthor("Mai Corland"));

console.log();
console.log(library.getTotalPages());







