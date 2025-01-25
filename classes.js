class Author {
    constructor(authorName, nationality) {
        this.authorName = authorName;
        this.nationality = nationality;
    }
}

class Genre {
    constructor(genreName) {
        this.genreName = genreName;
    }
}

class Book {
    constructor(title, pages, author, genre) {
        this.title = title;
        this.pages = pages;
        this.author = author;
        this.genre = genre;
    }
}

class Library {
    constructor(libraryName, location) {
        this.libraryName = libraryName;
        this.location = location;
        this.books = [];
        this.genres = [];
    }


addBook(book){
this.books.push(book);
}

addGenre(genre){
    this.genres.push(genre);
}
}

function createAuthor(authorName, nationality){
    const author = new Author(authorName, nationality);
    return author;
}

function createGenre(genreName){
    const genre = new Genre(genreName);
    return genre;
}

function createBook(title, pages, author, genre) {
    const book = new Book(title, pages, author, genre);
    return book;
}

function createLibrary(libraryName, location){
    const library = new Library(libraryName, location);
    return library;
}


const author1 = createAuthor("Daryl", "American");
const author2 = createAuthor("Peterson", "American");

const genre1 = createGenre("Comedy");
const genre2 = createGenre("fantasy");

const book1 = createBook("Daryl's World of Excellence", 600, author1, genre1);
const book2 = createBook("Peterson's Highway Expedition", 400, author2, genre2);

const library = createLibrary("New Orlean's Public Library", "New Orleans");

library.addBook(book1);
library.addBook(book2);
library.addGenre(genre1);
library.addGenre(genre2);

console.log(`${library.libraryName}, ${library.location}`);
console.log(`Books in the Library:`);
for (const b of library.books){
console.log(`- ${b.title} by ${b.author.authorName}`);
}
for (const g of library.genres){
    console.log(`- ${g.genreName}`);
}
